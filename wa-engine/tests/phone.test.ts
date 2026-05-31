import { describe, expect, it } from "vitest"
import { normalizePhone } from "../src/lib/phone.js"

describe("normalizePhone", () => {
  it.each([
    ["08123456789", "628123456789@s.whatsapp.net"],
    ["+628123456789", "628123456789@s.whatsapp.net"],
    ["628123456789", "628123456789@s.whatsapp.net"],
    ["08 123-456 (789)", "628123456789@s.whatsapp.net"],
  ])("normalizes %s", (input, output) => {
    expect(normalizePhone(input)).toBe(output)
  })
})
