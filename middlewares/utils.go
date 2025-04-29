package middlewares

import (
	"os"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UtilsMiddleware struct{}

// CORSMiddleware ...
// CORS (Cross-Origin Resource Sharing)
func (mid UtilsMiddleware) CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil origin frontend dari environment variable
		origin := os.Getenv("FE_URL")
		posibleOrigin := []string{origin}
		reqOrigin := c.Request.Header.Get("Origin")

		// Jika ada lebih dari satu origin yang diizinkan, ambil dari environment variable
		if os.Getenv("ALLOWED_ORIGIN") != "" {
			posibleOrigin = strings.Split(os.Getenv("ALLOWED_ORIGIN"), ",")
		}

		// Periksa apakah origin permintaan diizinkan
		checkOrigin := slices.Contains(posibleOrigin, reqOrigin)
		if checkOrigin {
			origin = reqOrigin
		}

		// Set header CORS untuk mengizinkan origin
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Jika permintaan adalah OPTIONS (preflight), kirim respons OK
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		// Lanjutkan ke permintaan berikutnya
		c.Next()
	}
}

// RequestIDMiddleware ...
// Generate a unique ID and attach it to each request for future reference or use
func (mid UtilsMiddleware) RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.New()
		c.Writer.Header().Set("X-Request-Id", uuid.String())
		c.Next()
	}
}
