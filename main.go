package main

import (
	"DynamicConfigDemo/src/dynamic_config"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		fmt.Println("Current redis host is: ", dynamic_config.GlobalConfig.GetString("service.redis.host"))
		context.JSON(
			200, gin.H{
				"message": "You are welcome!",
			})
	})
	r.Run(":9292")
}
