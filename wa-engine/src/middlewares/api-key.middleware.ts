import type { NextFunction, Request, Response } from "express"
import { env } from "../config/env.js"

export function apiKeyMiddleware(req: Request, res: Response, next: NextFunction) {
  const token = req.header("authorization")
  if (!isValidApiKey(token)) {
    res.status(401).json({ success: false, message: "Unauthorized" })
    return
  }
  next()
}

export function isValidApiKey(header: string | undefined): boolean {
  const token = header?.replace(/^Bearer\s+/i, "")
  return Boolean(token && token === env.apiKey)
}
