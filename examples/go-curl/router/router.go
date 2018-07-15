//用于处理路由对象关系
package router

import (
	"github.com/gin-gonic/gin"
	"gp-curl/controller"
)

func InitRoute() *gin.Engine {
	router := gin.Default()
	router.GET("/", controller.HomeController)
	router.GET("/users/:id", controller.UserController)
	router.GET("/users", controller.UsersQueryByAge)
	return router
}
