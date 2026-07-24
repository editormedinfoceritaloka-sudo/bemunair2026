import express, {
  type NextFunction,
  type Request,
  type Response,
} from "express"

import { apiKeyMiddleware } from "./middlewares/api-key.middleware.js"
import { messageRoutes } from "./routes/message.routes.js"

export function createApp() {
  const app = express()

  app.disable("x-powered-by")
  app.use(express.json({ limit: "1mb" }))

  app.get("/health", (_req, res) => {
    res.json({
      success: true,
      message: "ok",
      timestamp: new Date().toISOString(),
    })
  })

  app.use("/api", apiKeyMiddleware, messageRoutes)

  app.use(
    (
      err: Error,
      _req: Request,
      res: Response,
      _next: NextFunction,
    ) => {
      console.error("[WA-HTTP] request failed:", err)

      res.status(500).json({
        success: false,
        message:
          process.env.NODE_ENV === "production"
            ? "Internal server error"
            : err.message,
      })
    },
  )

  return app
}