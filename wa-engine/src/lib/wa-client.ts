import makeWASocket, { DisconnectReason, WASocket, fetchLatestBaileysVersion, useMultiFileAuthState } from "@whiskeysockets/baileys"
import { Boom } from "@hapi/boom"
import fs from "node:fs/promises"
import path from "node:path"
import qrcode from "qrcode"
import { env } from "../config/env.js"
import { normalizePhone } from "./phone.js"

type WAState = "disconnected" | "connecting" | "open" | "reconnecting" | "logged_out"

export type WAGroupInfo = {
  jid: string
  subject: string
  size: number
}

export class WAClient {
  private static instance: WAClient
  private socket: WASocket | null = null
  private state: WAState = "disconnected"
  private qrCode: string | null = null

  private constructor() {}

  static getInstance(): WAClient {
    if (!WAClient.instance) WAClient.instance = new WAClient()
    return WAClient.instance
  }

  async connect(): Promise<void> {
    if (this.state === "connecting" || this.state === "open") return
    this.state = "connecting"
    console.log("[WA] connecting...")
    const sessionPath = path.resolve(env.sessionDir)
    const { state, saveCreds } = await useMultiFileAuthState(sessionPath)
    const { version } = await fetchLatestBaileysVersion()
    this.socket = makeWASocket({ version, auth: state })
    this.socket.ev.on("creds.update", saveCreds)
    this.socket.ev.on("connection.update", async (update) => {
      const { connection, lastDisconnect, qr } = update
      if (qr) {
        this.qrCode = qr
        console.log("[WA] QR code received. Scan this QR from Docker logs:")
        console.log(await qrcode.toString(qr, { type: "terminal", small: true }))
        console.log(`[WA] QR raw: ${qr}`)
      }
      if (connection === "open") {
        this.state = "open"
        this.qrCode = null
        const user = this.socket?.user?.id ?? "unknown"
        console.log(`[WA] open - connected as ${user}`)
      }
      if (connection === "close") {
        const code = new Boom(lastDisconnect?.error).output.statusCode
        if (code === DisconnectReason.loggedOut) {
          this.state = "logged_out"
          this.socket = null
          await fs.rm(sessionPath, { recursive: true, force: true })
          console.log("[WA] logged_out - session cleared")
          return
        }
        this.state = "reconnecting"
        console.log("[WA] close - reconnecting in 3s...")
        setTimeout(() => void this.connect(), 3000)
      }
    })
  }

  async sendText(to: string, message: string): Promise<void> {
    const jid = normalizePhone(to)
    console.log(`[WA] sendText requested state=${this.state} from=${this.senderLabel()} to=${jid} message="${preview(message)}"`)
    if (this.state !== "open" || !this.socket) {
      console.error("[WA] sendText failed: socket is not open. Scan QR in docker logs first.")
      throw new Error("WA socket is not open")
    }
    await this.socket.sendMessage(jid, { text: message })
    console.log(`[WA] sendText sent to=${jid}`)
  }

  async sendGroup(groupJid: string, message: string): Promise<void> {
    console.log(`[WA] sendGroup requested state=${this.state} from=${this.senderLabel()} to=${groupJid} message="${preview(message)}"`)
    if (this.state !== "open" || !this.socket) {
      console.error("[WA] sendGroup failed: socket is not open. Scan QR in docker logs first.")
      throw new Error("WA socket is not open")
    }
    await this.socket.sendMessage(groupJid, { text: message })
    console.log(`[WA] sendGroup sent to=${groupJid}`)
  }

  async listGroups(): Promise<WAGroupInfo[]> {
    console.log(`[WA] listGroups requested state=${this.state} from=${this.senderLabel()}`)
    if (this.state !== "open" || !this.socket) {
      console.error("[WA] listGroups failed: socket is not open. Scan QR in docker logs first.")
      throw new Error("WA socket is not open")
    }
    const groups = await this.socket.groupFetchAllParticipating()
    const result = Object.values(groups)
      .map((group) => ({
        jid: group.id,
        subject: group.subject ?? "(no subject)",
        size: group.participants?.length ?? 0,
      }))
      .sort((a, b) => a.subject.localeCompare(b.subject))
    console.log(`[WA] groups found count=${result.length}`)
    for (const group of result) {
      console.log(`[WA-GROUP] jid=${group.jid} subject="${group.subject}" size=${group.size}`)
    }
    return result
  }

  getStatus(): { state: string; qr?: string } {
    return this.qrCode ? { state: this.state, qr: this.qrCode } : { state: this.state }
  }

  private senderLabel(): string {
    return this.socket?.user?.id ?? (env.senderPhone || "unknown")
  }
}

function preview(message: string): string {
  return message.replace(/\s+/g, " ").slice(0, 120)
}
