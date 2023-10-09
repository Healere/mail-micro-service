package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
	c.Writer.Header().Set("Access-Control-Allow-headers", "*")
	c.Writer.Header().Set("Access-Control-Max-Age", "86400")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Content-Type", "application/json")

	if c.Request.Method == http.MethodOptions {
		c.AbortWithStatus(200)
	} else {
		c.Next()
	}
}
