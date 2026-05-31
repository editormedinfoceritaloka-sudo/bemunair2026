import { Router } from "express"
import qrcode from "qrcode"
import { WAClient } from "../lib/wa-client.js"

export const messageRoutes = Router()

messageRoutes.get("/status", (req, res) => {
  const status = WAClient.getInstance().getStatus()
  console.log(`[WA-HTTP] GET /api/status state=${status.state}`)
  res.json({ success: true, data: status })
})

messageRoutes.get("/qr", async (req, res) => {
  const status = WAClient.getInstance().getStatus()
  console.log(`[WA-HTTP] GET /api/qr state=${status.state} hasQr=${Boolean(status.qr)}`)
  if (status.state === "open") {
    res.json({ success: true, data: { state: "open", qr: null } })
    return
  }
  if (!status.qr) {
    res.status(404).json({ success: false, message: "QR belum tersedia" })
    return
  }
  res.json({ success: true, data: { image: await qrcode.toDataURL(status.qr), qr: status.qr } })
})

messageRoutes.get("/groups", async (req, res, next) => {
  try {
    console.log("[WA-HTTP] GET /api/groups")
    const groups = await WAClient.getInstance().listGroups()
    res.json({ success: true, message: "groups fetched", data: groups })
  } catch (error) { next(error) }
})

messageRoutes.post("/connect", async (req, res, next) => {
  try {
    console.log("[WA-HTTP] POST /api/connect")
    await WAClient.getInstance().connect()
    res.json({ success: true, message: "connect triggered", data: WAClient.getInstance().getStatus() })
  } catch (error) { next(error) }
})

messageRoutes.post("/send-message", async (req, res, next) => {
  try {
    console.log(`[WA-HTTP] POST /api/send-message to=${req.body.to} message="${preview(req.body.message)}"`)
    await WAClient.getInstance().sendText(req.body.to, req.body.message)
    res.json({ success: true, message: "message sent" })
  } catch (error) { next(error) }
})

messageRoutes.post("/send-group-message", async (req, res, next) => {
  try {
    console.log(`[WA-HTTP] POST /api/send-group-message groupJid=${req.body.groupJid} message="${preview(req.body.message)}"`)
    await WAClient.getInstance().sendGroup(req.body.groupJid, req.body.message)
    res.json({ success: true, message: "group message sent" })
  } catch (error) { next(error) }
})

function preview(message: unknown): string {
  return String(message ?? "").replace(/\s+/g, " ").slice(0, 120)
}
