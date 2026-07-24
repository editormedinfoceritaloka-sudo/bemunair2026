import { createApp } from "./app.js"
import { env } from "./config/env.js"
import { WAClient } from "./lib/wa-client.js"

const app = createApp()
const waClient = WAClient.getInstance()

const server = app.listen(env.port, () => {
  console.log(`[WA Engine] Running on port ${env.port}`)

  if (env.senderPhone) {
    console.log(`[WA] sender phone label: ${env.senderPhone}`)
  }

  if (env.autoConnect) {
    void waClient.connect().catch((error: unknown) => {
      console.error(
        "[WA] auto-connect failed:",
        error instanceof Error ? error.message : String(error),
      )
    })
  }
})

let isShuttingDown = false

async function shutdown(signal: string): Promise<void> {
  if (isShuttingDown) {
    return
  }

  isShuttingDown = true

  console.log(`[WA Engine] ${signal} received, shutting down...`)

  // Hentikan penerimaan request baru.
  server.close((error) => {
    if (error) {
      console.error("[WA Engine] HTTP shutdown failed:", error)
      process.exitCode = 1
    } else {
      console.log("[WA Engine] HTTP server closed")
    }
  })

  try {
    // Menutup WebSocket Baileys tanpa logout dari WhatsApp.
    await waClient.shutdown()
    console.log("[WA] connection closed safely")
  } catch (error) {
    console.error("[WA] shutdown failed:", error)
    process.exitCode = 1
  }

  process.exit()
}

process.once("SIGTERM", () => {
  void shutdown("SIGTERM")
})

process.once("SIGINT", () => {
  void shutdown("SIGINT")
})

process.on("unhandledRejection", (reason: unknown) => {
  console.error("[WA Engine] unhandled rejection:", reason)
})

process.on("uncaughtException", (error: Error) => {
  console.error("[WA Engine] uncaught exception:", error)

  void shutdown("uncaughtException")
})