import { describe, expect, it } from "vitest"
import { env } from "../src/config/env.js"
import { isValidApiKey } from "../src/middlewares/api-key.middleware.js"

describe("api key middleware", () => {
  it("rejects missing api key", async () => {
    expect(isValidApiKey(undefined)).toBe(false)
    expect(isValidApiKey("Bearer wrong")).toBe(false)
    expect(isValidApiKey(`Bearer ${env.apiKey}`)).toBe(true)
  })
})
