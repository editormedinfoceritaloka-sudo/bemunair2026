import { describe, expect, it } from "vitest"
import { WAClient } from "../src/lib/wa-client.js"

describe("WAClient", () => {
  it("is a singleton", () => {
    expect(WAClient.getInstance()).toBe(WAClient.getInstance())
  })

  it("throws when sending while not open", async () => {
    await expect(WAClient.getInstance().sendText("081", "hi")).rejects.toThrow("WA socket is not open")
  })
})
