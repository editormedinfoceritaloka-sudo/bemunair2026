package middlewares

import (
	"net/http"
	"strings"
	"time"

	"bemunair2026/server/pkg/constants"
	response "bemunair2026/server/pkg/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   uint64  `json:"user_id"`
	Role     string  `json:"role"`
	Ministry *string `json:"ministry,omitempty"`
	jwt.RegisteredClaims
}

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost", "http://localhost:8081", "http://localhost:3000", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "X-Request-Id"},
		ExposeHeaders:    []string{"X-Request-Id"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-Id")
		if requestID == "" {
			requestID = time.Now().Format("20060102150405.000000000")
		}
		c.Header("X-Request-Id", requestID)
		c.Set("request_id", requestID)
		start := time.Now()
		c.Next()
		gin.DefaultWriter.Write([]byte(c.Request.Method + " " + c.Request.URL.Path + " " + time.Since(start).String() + "\n"))
	}
}

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered any) {
		response.Error(c, http.StatusInternalServerError, response.InternalError, "Terjadi kesalahan internal")
	})
}

func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			response.Error(c, http.StatusUnauthorized, response.Unauthenticated, "Token tidak ditemukan")
			c.Abort()
			return
		}
		tokenText := strings.TrimPrefix(header, "Bearer ")
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenText, claims, func(token *jwt.Token) (any, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			response.Error(c, http.StatusUnauthorized, response.Unauthenticated, "Token tidak valid")
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func Role(roles ...string) gin.HandlerFunc {
	allowed := map[string]bool{}
	for _, role := range roles {
		allowed[role] = true
	}
	return func(c *gin.Context) {
		claims, ok := c.MustGet("claims").(*Claims)
		if !ok || !allowed[claims.Role] {
			response.Error(c, http.StatusForbidden, response.Forbidden, "Akses ditolak")
			c.Abort()
			return
		}
		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return Role(constants.RoleAdmin)
}

func MentriOnly() gin.HandlerFunc {
	return Role(constants.RoleMentri)
}

func CurrentClaims(c *gin.Context) *Claims {
	claims, _ := c.Get("claims")
	if typed, ok := claims.(*Claims); ok {
		return typed
	}
	return nil
}
