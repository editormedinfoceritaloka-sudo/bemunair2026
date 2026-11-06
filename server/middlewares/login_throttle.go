package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	response "bemunair2026/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

// LoginLimiter throttles repeated failed login attempts independently of any
// edge/reverse-proxy rate limiter. It is keyed on the real TCP peer address
// (never a client-supplied header, which cannot be trusted at this layer)
// combined with the attempted email, so one abusive client cannot exhaust a
// shared account-only or IP-only budget for everyone else.
//
// Only failed attempts (401) count against the budget; a successful login
// clears it. Entries are swept on a timer so an attacker cycling through
// many email addresses cannot grow the map without bound.
type LoginLimiter struct {
	mu       sync.Mutex
	attempts map[string]*loginAttempt
	max      int
	window   time.Duration
}

type loginAttempt struct {
	count   int
	resetAt time.Time
}

// defaultLoginLimiter backs the auth routes. A package-level singleton is
// intentional: the login endpoint is registered once per process and the
// budget must be shared across every request to it.
var defaultLoginLimiter = NewLoginLimiter(5, 10*time.Minute)

// LoginThrottle returns the middleware guarding the login route. Independent
// of any edge-level (nginx) rate limiting, so a brute-force attempt is still
// bounded even if the edge limiter is bypassed, misconfigured, or absent
// (e.g. hitting the server port directly in a dev environment).
func LoginThrottle() gin.HandlerFunc {
	return defaultLoginLimiter.Middleware()
}

func NewLoginLimiter(max int, window time.Duration) *LoginLimiter {
	l := &LoginLimiter{
		attempts: make(map[string]*loginAttempt),
		max:      max,
		window:   window,
	}
	go l.sweep()
	return l
}

func (l *LoginLimiter) sweep() {
	ticker := time.NewTicker(l.window)
	defer ticker.Stop()
	for now := range ticker.C {
		l.mu.Lock()
		for key, a := range l.attempts {
			if now.After(a.resetAt) {
				delete(l.attempts, key)
			}
		}
		l.mu.Unlock()
	}
}

func (l *LoginLimiter) key(c *gin.Context, email string) string {
	return clientRemoteIP(c.Request) + "|" + strings.ToLower(strings.TrimSpace(email))
}

// Middleware blocks requests whose ip+email key is already over budget, lets
// the handler run otherwise, then records the outcome: 401 counts as a
// failed attempt, 200/201 clears any accumulated count for that key.
func (l *LoginLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := peekLoginEmail(c)
		key := l.key(c, email)
		now := time.Now()

		l.mu.Lock()
		a, exists := l.attempts[key]
		if exists && now.Before(a.resetAt) && a.count >= l.max {
			retryAfter := int(a.resetAt.Sub(now).Seconds())
			l.mu.Unlock()
			c.Header("Retry-After", fmt.Sprintf("%d", retryAfter))
			response.Error(c, http.StatusTooManyRequests, response.RateLimited,
				fmt.Sprintf("Terlalu banyak percobaan login. Coba lagi dalam %ds.", retryAfter))
			c.Abort()
			return
		}
		l.mu.Unlock()

		c.Next()

		status := c.Writer.Status()
		l.mu.Lock()
		switch {
		case status == http.StatusUnauthorized:
			a, exists := l.attempts[key]
			if !exists || now.After(a.resetAt) {
				a = &loginAttempt{resetAt: now.Add(l.window)}
				l.attempts[key] = a
			}
			a.count++
		case status == http.StatusOK || status == http.StatusCreated:
			delete(l.attempts, key)
		}
		l.mu.Unlock()
	}
}

// clientRemoteIP returns the actual TCP peer address nginx/Go net/http saw,
// stripped of its port. Deliberately ignores X-Forwarded-For and friends:
// at this layer those are advisory at best, and in a dev setup the server
// port can be reached directly, bypassing any upstream proxy entirely.
func clientRemoteIP(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

type loginEmailPeek struct {
	Email string `json:"email"`
}

// peekLoginEmail reads the request body to extract the attempted email for
// per-account keying, then restores the body so the real handler's bind
// still sees the full payload. A body that isn't valid/readable JSON simply
// degrades the key to IP-only, which is still safe.
func peekLoginEmail(c *gin.Context) string {
	if c.Request.Body == nil {
		return ""
	}
	raw, err := io.ReadAll(io.LimitReader(c.Request.Body, 1<<20))
	c.Request.Body = io.NopCloser(bytes.NewReader(raw))
	if err != nil {
		return ""
	}
	var peek loginEmailPeek
	_ = json.Unmarshal(raw, &peek)
	return peek.Email
}
