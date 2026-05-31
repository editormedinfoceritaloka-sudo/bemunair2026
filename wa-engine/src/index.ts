import { createApp } from "./app.js"
import { env } from "./config/env.js"
import { WAClient } from "./lib/wa-client.js"

createApp().listen(env.port, () => {
  console.log(`[WA Engine] Running on port ${env.port}`)
  if (env.senderPhone) console.log(`[WA] sender phone label: ${env.senderPhone}`)
  if (env.autoConnect) {
    void WAClient.getInstance().connect().catch((error) => {
      console.error(`[WA] auto-connect failed: ${error instanceof Error ? error.message : String(error)}`)
    })
  }
})
