package controller

import (
	"github.com/gin-gonic/gin"

	"learnGo/examples/go-curl/model"
	"learnGo/examples/go-curl/service"
	"net/http"
)

func HomeController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "this is home route",
	})
}

////使用id查询用户信息
//func UserController(c *gin.Context) {
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"status":  400,
//			"message": "must be int",
//			"result":  nil,
//		})
//		return
//	}
//	u := service.User{Id: id}
//
//	res, err := u.GetUserById()
//
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"status":  400,
//			"message": err.Error(),
//			"result":  nil,
//		})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"status":  200,
//		"message": "ok",
//		"result":  res,
//	})
//}
//
////查询所有同龄人
//func UsersQueryByAge(c *gin.Context) {
//	age, err := strconv.Atoi(c.Query("age"))
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"status":  http.StatusBadRequest,
//			"message": "age must be int",
//		})
//		return
//	}
//	pageIndex, err := strconv.Atoi(c.DefaultQuery("pageIndex", "1"))
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"status":  http.StatusBadRequest,
//			"message": "pageIndex must be int",
//		})
//		return
//	}
//	pageCount, err := strconv.Atoi(c.DefaultQuery("pageCount", "15"))
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"status":  http.StatusBadRequest,
//			"message": "pageCount must be int",
//		})
//		return
//	}
//	u := &service.User{Age: age}
//	res, err := u.GetUsersByAge(pageIndex, pageCount)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"status":  http.StatusBadRequest,
//			"message": err.Error(),
//			"result":  nil,
//		})
//		return
//	}
//	if len(res) == 0 {
//		c.JSON(http.StatusNotFound, gin.H{
//			"status":  404,
//			"message": "no result",
//			"result":  nil,
//		})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"status":  200,
//		"message": "ok",
//		"result":  res,
//	})
//}

// 用户名登录

func UserSignInByLoginName(c *gin.Context) {
	// 判断参数是否合法
	var json model.UserLoginRequest
	err := c.ShouldBindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"result":  nil,
		})
		return
	}
	// 参数合法 开始调用逻辑
	serviceResponse, err := service.SignInByLoginName(json.User, json.Password)

	// 调用非法
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"result":  nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "",
		"result":  serviceResponse,
	})

}
