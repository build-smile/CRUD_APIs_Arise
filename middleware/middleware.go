package middleware

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	secretKey string
}

func NewMiddleware(secretKey string) *Middleware {
	return &Middleware{
		secretKey: secretKey,
	}
}

func (m Middleware) ValidateOptionMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Status(http.StatusNoContent)
			return
		}
		c.Next()
	}

}
func (m Middleware) Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log the error and stack trace
				log.Printf("Recovered from panic: %v\n%s", err, debug.Stack())

				// Send a generic error response
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error",
				})
				c.Abort()
			}
		}()

		// Continue with the next handler
		c.Next()
	}
}
