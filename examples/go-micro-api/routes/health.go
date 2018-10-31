package routes

import (
	"github.com/gin-gonic/gin"
	"learnGo/examples/go-micro-api/controller"
)

func AddHealth(r *gin.Engine) *gin.Engine {
	ac := controller.AccountController{}
	r.GET("/api/health", ac.GetHealthStatus)
	return r
}
