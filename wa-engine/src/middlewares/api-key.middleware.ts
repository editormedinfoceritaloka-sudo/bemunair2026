import type {
  NextFunction,
  Request,
  Response,
} from "express"

import { env } from "../config/env.js"

export function isValidApiKey(authorization?: string): boolean {
  const providedKey = authorization
    ?.replace(/^Bearer\s+/i, "")
    .trim()

  return Boolean(providedKey && providedKey === env.apiKey.trim())
}

export function apiKeyMiddleware(
  req: Request,
  res: Response,
  next: NextFunction,
): void {
  const authorization = req.header("authorization")

  const bearerToken = authorization
    ?.replace(/^Bearer\s+/i, "")
    .trim()

  const xApiKey = req
    .header("x-api-key")
    ?.trim()

  const providedKey = bearerToken || xApiKey

  if (!isValidApiKey(providedKey)) {
    console.warn("[WA-HTTP] Unauthorized", {
      method: req.method,
      path: req.originalUrl,
      hasBearerToken: Boolean(bearerToken),
      hasXApiKey: Boolean(xApiKey),
    })

    res.status(401).json({
      success: false,
      message: "Unauthorized",
    })

    return
  }

  next()
}