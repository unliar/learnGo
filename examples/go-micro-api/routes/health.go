package routes

import "github.com/gin-gonic/gin"

func AddHealth(r *gin.Engine) *gin.Engine {
	r.GET("/api/health", AC.GetHealthStatus)
	return r
}
