package routes

import (
	"mailer-service/cmd/api/middleware"

	"github.com/gin-gonic/gin"
)

func Collection(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware)
	//r.POST("/api/email", controller.SendMail)
	return r
}
