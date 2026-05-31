package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type WASender interface {
	SendTextMessage(phone, message string) error
	SendGroupMessage(groupJID, message string) error
}

type WAClient struct {
	URL    string
	APIKey string
	Client *http.Client
}

func NewWAClient(url, apiKey string) *WAClient {
	return &WAClient{URL: strings.TrimRight(url, "/"), APIKey: apiKey, Client: &http.Client{Timeout: 10 * time.Second}}
}

func (c *WAClient) SendTextMessage(phone, message string) error {
	log.Printf("[WA-BACKEND] queue personal notification target=%s message=%q", phone, previewWAMessage(message))
	return c.post("/api/send-message", map[string]string{"to": phone, "message": message})
}

func (c *WAClient) SendGroupMessage(groupJID, message string) error {
	if groupJID == "" {
		log.Printf("[WA-BACKEND] skip group notification: empty groupJID message=%q", previewWAMessage(message))
		return nil
	}
	log.Printf("[WA-BACKEND] queue group notification groupJID=%s message=%q", groupJID, previewWAMessage(message))
	return c.post("/api/send-group-message", map[string]string{"groupJid": groupJID, "message": message})
}

func (c *WAClient) post(path string, body any) error {
	if c.URL == "" || c.APIKey == "" {
		log.Printf("[WA-BACKEND] skip request path=%s reason=missing WA_ENGINE_URL or WA_ENGINE_API_KEY", path)
		return nil
	}
	payload, _ := json.Marshal(body)
	log.Printf("[WA-BACKEND] POST %s%s payload=%s", c.URL, path, string(payload))
	req, err := http.NewRequest(http.MethodPost, c.URL+path, bytes.NewReader(payload))
	if err != nil {
		log.Printf("[WA-BACKEND] request build failed path=%s error=%v", path, err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	resp, err := c.Client.Do(req)
	if err != nil {
		log.Printf("[WA-BACKEND] request failed path=%s error=%v", path, err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		log.Printf("[WA-BACKEND] request returned error path=%s status=%d", path, resp.StatusCode)
		return fmt.Errorf("wa engine status %d", resp.StatusCode)
	}
	log.Printf("[WA-BACKEND] request success path=%s status=%d", path, resp.StatusCode)
	return nil
}

func previewWAMessage(message string) string {
	message = strings.Join(strings.Fields(message), " ")
	if len(message) > 120 {
		return message[:120]
	}
	return message
}
