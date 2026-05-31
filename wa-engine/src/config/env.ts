import "dotenv/config"

export const env = {
  port: Number(process.env.WA_ENGINE_PORT ?? process.env.PORT ?? 3001),
  apiKey: process.env.WA_ENGINE_API_KEY ?? "dev_internal_key",
  sessionDir: process.env.WA_SESSION_DIR ?? "./sessions",
  autoConnect: (process.env.WA_AUTO_CONNECT ?? "true") === "true",
  senderPhone: process.env.WA_SENDER_PHONE ?? "",
}
