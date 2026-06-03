package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"bemunair2026/server/pkg/constants"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func TestAuthAndRoleMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	secret := "secret"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{UserID: 1, Role: constants.RoleAdmin, RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour))}})
	signed, _ := token.SignedString([]byte(secret))
	r := gin.New()
	r.GET("/admin", Auth(secret), AdminOnly(), func(c *gin.Context) { c.Status(http.StatusNoContent) })
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/admin", nil)
	req.Header.Set("Authorization", "Bearer "+signed)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusNoContent {
		t.Fatalf("admin status=%d", w.Code)
	}

	w = httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/admin", nil))
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("missing token status=%d", w.Code)
	}
}
