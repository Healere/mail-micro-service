package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Accsee-Control-Max-Age", "86400")
		c.Writer.Header().Set("Accsee-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Accsee-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("content-type", "application/json")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
