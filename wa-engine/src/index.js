import makeWASocket, {
  useMultiFileAuthState,
  DisconnectReason,
  fetchLatestBaileysVersion,
} from "@whiskeysockets/baileys";
import { Boom } from "@hapi/boom";
import path from "path";
import "dotenv/config";

const AUTH_DIR = process.env.AUTH_DIR ?? "./auth_info_baileys";
const PORT = process.env.PORT ?? 3001;

async function connectToWhatsApp() {
  const { state, saveCreds } = await useMultiFileAuthState(
    path.resolve(AUTH_DIR)
  );

  const { version } = await fetchLatestBaileysVersion();
  console.log(`[WA] Using Baileys v${version.join(".")}`);

  const sock = makeWASocket({
    version,
    auth: state,
    printQRInTerminal: process.env.WA_PRINT_QR_IN_TERMINAL === "true",
  });

  sock.ev.on("creds.update", saveCreds);

  sock.ev.on("connection.update", (update) => {
    const { connection, lastDisconnect, qr } = update;

    if (qr) {
      console.log("[WA] Scan QR code above to login");
    }

    if (connection === "close") {
      const reason = new Boom(lastDisconnect?.error)?.output?.statusCode;
      const shouldReconnect = reason !== DisconnectReason.loggedOut;

      console.log(`[WA] Connection closed. Reason: ${reason}. Reconnect: ${shouldReconnect}`);

      if (shouldReconnect) {
        setTimeout(connectToWhatsApp, Number(process.env.WA_RECONNECT_INTERVAL) || 5000);
      } else {
        console.log("[WA] Logged out. Delete auth folder and restart.");
      }
    }

    if (connection === "open") {
      console.log("[WA] Connected to WhatsApp!");
    }
  });

  sock.ev.on("messages.upsert", async ({ messages, type }) => {
    if (type !== "notify") return;

    for (const msg of messages) {
      if (!msg.key.fromMe) {
        console.log("[WA] Incoming:", JSON.stringify(msg, null, 2));
        // TODO: forward ke SERVER_WEBHOOK_URL
      }
    }
  });

  return sock;
}

connectToWhatsApp();
console.log(`[WA Engine] Running on port ${PORT}`);