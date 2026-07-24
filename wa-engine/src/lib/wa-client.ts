import { Boom } from "@hapi/boom"
import makeWASocket, {
  DisconnectReason,
  fetchLatestBaileysVersion,
  useMultiFileAuthState,
  type BaileysEventMap,
  type WASocket,
} from "@whiskeysockets/baileys"
import fs from "node:fs/promises"
import path from "node:path"
import qrcode from "qrcode"
import { env } from "../config/env.js"
import { normalizePhone } from "./phone.js"

type WAState =
  | "disconnected"
  | "connecting"
  | "open"
  | "reconnecting"
  | "logged_out"

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

  private connectPromise: Promise<void> | null = null
  private reconnectTimer: ReturnType<typeof setTimeout> | null = null
  private reconnectAttempt = 0

  /**
   * Dipakai untuk mengabaikan event dari socket lama.
   */
  private socketGeneration = 0

  private constructor() {}

  static getInstance(): WAClient {
    if (!WAClient.instance) {
      WAClient.instance = new WAClient()
    }

    return WAClient.instance
  }

  async connect(): Promise<void> {
    if (this.state === "open" && this.socket) {
      return
    }

    // Mencegah dua proses connect berjalan bersamaan.
    if (this.connectPromise) {
      return this.connectPromise
    }

    this.clearReconnectTimer()

    this.connectPromise = this.createSocket().finally(() => {
      this.connectPromise = null
    })

    return this.connectPromise
  }

  private async createSocket(): Promise<void> {
    const sessionPath = path.resolve(env.sessionDir)

    await fs.mkdir(sessionPath, { recursive: true })

    this.state =
      this.reconnectAttempt > 0 ? "reconnecting" : "connecting"

    console.log(
      `[WA] ${this.state} attempt=${this.reconnectAttempt}`,
    )

    try {
      const { state, saveCreds } =
        await useMultiFileAuthState(sessionPath)

      const { version, isLatest } =
        await fetchLatestBaileysVersion()

      console.log(
        `[WA] Baileys version=${version.join(".")} latest=${isLatest}`,
      )

      const socket = makeWASocket({
        version,
        auth: state,
      })

      const generation = ++this.socketGeneration
      this.socket = socket

      socket.ev.on("creds.update", () => {
        // Jangan biarkan socket lama menulis ulang sesi.
        if (!this.isCurrentSocket(socket, generation)) {
          return
        }

        void saveCreds().catch((error: unknown) => {
          console.error("[WA] failed to save credentials:", error)
        })
      })

      socket.ev.on("connection.update", (update) => {
        void this.handleConnectionUpdate(
          socket,
          generation,
          sessionPath,
          update,
        ).catch((error: unknown) => {
          console.error(
            "[WA] connection update handler failed:",
            error,
          )

          if (this.isCurrentSocket(socket, generation)) {
            this.socket = null
            this.socketGeneration += 1
            this.state = "reconnecting"
            this.scheduleReconnect()
          }
        })
      })
    } catch (error) {
      this.socket = null
      this.state = "reconnecting"

      console.error("[WA] socket creation failed:", error)

      this.scheduleReconnect()
      throw error
    }
  }

  private async handleConnectionUpdate(
    socket: WASocket,
    generation: number,
    sessionPath: string,
    update: BaileysEventMap["connection.update"],
  ): Promise<void> {
    if (!this.isCurrentSocket(socket, generation)) {
      return
    }

    const { connection, lastDisconnect, qr } = update

    if (qr) {
      this.qrCode = qr
      this.state = "connecting"

      console.log("[WA] QR code received")
      console.log(
        await qrcode.toString(qr, {
          type: "terminal",
          small: true,
        }),
      )
    }

    if (connection === "connecting") {
      console.log("[WA] connecting...")
      return
    }

    if (connection === "open") {
      this.state = "open"
      this.qrCode = null
      this.reconnectAttempt = 0
      this.clearReconnectTimer()

      const user = socket.user?.id ?? "unknown"
      console.log(`[WA] open - connected as ${user}`)
      return
    }

    if (connection !== "close") {
      return
    }

    const statusCode = (
      lastDisconnect?.error as Boom | undefined
    )?.output?.statusCode

    const message =
      lastDisconnect?.error instanceof Error
        ? lastDisconnect.error.message
        : String(lastDisconnect?.error ?? "unknown error")

    console.error("[WA] connection closed:", {
      statusCode,
      message,
    })

    // Jadikan event berikutnya dari socket ini stale.
    this.socket = null
    this.socketGeneration += 1
    this.qrCode = null

    switch (statusCode) {
      case DisconnectReason.loggedOut:
      case DisconnectReason.badSession: {
        this.clearReconnectTimer()
        this.state = "logged_out"

        // Hapus isi session, bukan folder mount /app/sessions.
        await this.clearSessionContents(sessionPath)

        console.error(
          "[WA] session invalid/logged out. Session contents cleared.",
        )
        console.error(
          "[WA] call POST /api/connect to generate a new QR.",
        )
        return
      }

      case DisconnectReason.connectionReplaced:
        this.clearReconnectTimer()
        this.state = "disconnected"

        console.error(
          "[WA] connection replaced by another WhatsApp session.",
        )
        return

      case DisconnectReason.forbidden:
      case DisconnectReason.multideviceMismatch:
        this.clearReconnectTimer()
        this.state = "disconnected"

        console.error(
          `[WA] terminal disconnect status=${statusCode}; automatic reconnect disabled.`,
        )
        return

      case DisconnectReason.restartRequired:
        /*
         * Normal setelah pairing: socket lama harus diganti
         * dengan socket baru.
         */
        this.state = "reconnecting"
        this.scheduleReconnect(1_000)
        return

      default:
        this.state = "reconnecting"
        this.scheduleReconnect()
    }
  }

  private scheduleReconnect(fixedDelayMs?: number): void {
    if (this.reconnectTimer) {
      return
    }

    this.reconnectAttempt += 1

    const exponentialDelay = Math.min(
      30_000,
      1_000 * 2 ** Math.min(this.reconnectAttempt - 1, 5),
    )

    const jitter = Math.floor(Math.random() * 500)
    const delayMs =
      fixedDelayMs ?? exponentialDelay + jitter

    console.log(
      `[WA] reconnect scheduled in ${delayMs}ms attempt=${this.reconnectAttempt}`,
    )

    this.reconnectTimer = setTimeout(() => {
      this.reconnectTimer = null

      void this.connect().catch((error: unknown) => {
        console.error("[WA] reconnect attempt failed:", error)

        this.state = "reconnecting"
        this.scheduleReconnect()
      })
    }, delayMs)
  }

  private clearReconnectTimer(): void {
    if (!this.reconnectTimer) {
      return
    }

    clearTimeout(this.reconnectTimer)
    this.reconnectTimer = null
  }

  /**
   * Docker volume root tidak dapat dihapus.
   * Hanya file dan subfolder di dalamnya yang dihapus.
   */
  private async clearSessionContents(
    sessionPath: string,
  ): Promise<void> {
    await fs.mkdir(sessionPath, { recursive: true })

    const entries = await fs.readdir(sessionPath, {
      withFileTypes: true,
    })

    await Promise.all(
      entries.map((entry) =>
        fs.rm(path.join(sessionPath, entry.name), {
          recursive: true,
          force: true,
        }),
      ),
    )
  }

  private isCurrentSocket(
    socket: WASocket,
    generation: number,
  ): boolean {
    return (
      this.socket === socket &&
      this.socketGeneration === generation
    )
  }

  async sendText(
    to: string,
    message: string,
  ): Promise<void> {
    const socket = this.requireOpenSocket()
    const jid = normalizePhone(to)

    if (!message.trim()) {
      throw new Error("Message cannot be empty")
    }

    console.log(
      `[WA] sendText from=${this.senderLabel()} to=${jid} message="${preview(message)}"`,
    )

    await socket.sendMessage(jid, { text: message })

    console.log(`[WA] sendText sent to=${jid}`)
  }

  async sendGroup(
    groupJid: string,
    message: string,
  ): Promise<void> {
    const socket = this.requireOpenSocket()

    if (!groupJid.endsWith("@g.us")) {
      throw new Error("Invalid WhatsApp group JID")
    }

    if (!message.trim()) {
      throw new Error("Message cannot be empty")
    }

    console.log(
      `[WA] sendGroup from=${this.senderLabel()} to=${groupJid} message="${preview(message)}"`,
    )

    await socket.sendMessage(groupJid, { text: message })

    console.log(`[WA] sendGroup sent to=${groupJid}`)
  }

  async listGroups(): Promise<WAGroupInfo[]> {
    const socket = this.requireOpenSocket()

    console.log(
      `[WA] listGroups from=${this.senderLabel()}`,
    )

    const groups =
      await socket.groupFetchAllParticipating()

    return Object.values(groups)
      .map((group) => ({
        jid: group.id,
        subject: group.subject ?? "(no subject)",
        size: group.participants?.length ?? 0,
      }))
      .sort((a, b) =>
        a.subject.localeCompare(b.subject),
      )
  }

  getStatus(): {
    state: WAState
    qr?: string
    user?: string
  } {
    return {
      state: this.state,
      ...(this.qrCode ? { qr: this.qrCode } : {}),
      ...(this.socket?.user?.id
        ? { user: this.socket.user.id }
        : {}),
    }
  }

  async shutdown(): Promise<void> {
    this.clearReconnectTimer()

    const socket = this.socket

    this.socket = null
    this.socketGeneration += 1
    this.state = "disconnected"
    this.qrCode = null

    if (socket) {
      await socket.end(undefined).catch((error: unknown) => {
        console.error("[WA] socket shutdown failed:", error)
      })
    }
  }

  private requireOpenSocket(): WASocket {
    if (this.state !== "open" || !this.socket) {
      throw new Error(
        `WA socket is not open; current state=${this.state}`,
      )
    }

    return this.socket
  }

  private senderLabel(): string {
    return (
      this.socket?.user?.id ??
      env.senderPhone ??
      "unknown"
    )
  }
}

function preview(message: string): string {
  return message.replace(/\s+/g, " ").slice(0, 120)
}