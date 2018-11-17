package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		err := SendString(context.ClientIP())
		if err != nil {
			context.JSON(200, err.Error())
			return
		}
		context.JSON(200, gin.H{
			"status": "ok",
		})

	})
	r.Run(":8888")
}
