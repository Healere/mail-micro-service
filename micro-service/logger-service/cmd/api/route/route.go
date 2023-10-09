package route

import (
	"log-service/cmd/api/controller"
	"log-service/cmd/api/middleware"

	"github.com/gin-gonic/gin"
)

func Collections(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleWare())
	r.POST("api/log", controller.WriteLog)
	return r
}
