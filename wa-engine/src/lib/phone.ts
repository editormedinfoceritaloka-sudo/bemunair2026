export function normalizePhone(phone: string): string {
  let p = phone.replace(/[\s\-\(\)]/g, "")
  if (p.startsWith("+")) p = p.slice(1)
  if (p.startsWith("08")) p = "62" + p.slice(1)
  if (!p.startsWith("62")) p = "62" + p
  return p + "@s.whatsapp.net"
}
