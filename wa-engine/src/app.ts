import express from "express"
import { apiKeyMiddleware } from "./middlewares/api-key.middleware.js"
import { messageRoutes } from "./routes/message.routes.js"

export function createApp() {
  const app = express()
  app.use(express.json())
  app.get("/health", (req, res) => res.json({ success: true, message: "ok" }))
  app.use("/api", apiKeyMiddleware, messageRoutes)
  app.use((err: Error, req: express.Request, res: express.Response, next: express.NextFunction) => {
    res.status(500).json({ success: false, message: err.message })
  })
  return app
}
