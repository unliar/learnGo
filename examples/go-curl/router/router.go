//用于处理路由对象关系
package router

import (
	"github.com/gin-gonic/gin"
	"learnGo/examples/go-curl/controller"
)

func InitRoute() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.GET("/", controller.HomeController)
	//router.GET("/api/users/:id", controller.UserController)
	//router.GET("/api/users", controller.UsersQueryByAge)
	router.POST("/api/users/login", controller.UserSignInByLoginName)
	return router
}
