package routes

import (
	"github.com/gin-gonic/gin"
	"learnGo/examples/go-micro-api/controller"
)

func PayRouter(r *gin.Engine) *gin.Engine {
	pc := &controller.PayController{}
	r.GET("/api/pay/:uid", pc.GetPayInfo)
	r.POST("/api/pay", AC.JWTAuth(), pc.PostPayInfo)
	r.PUT("/api/pay", AC.JWTAuth(), pc.UpdatePayInfo)
	return r
}
