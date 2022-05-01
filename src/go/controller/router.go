package controller

import (
	"../mqtt"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
author Lei Wang
*/

func Run() {
	//初始化Router和MqttClient
	clientPointer := mqtt.GetMqttClient()
	//获取一个Mqtt客户端
	client := *clientPointer
	client.Connect()
	router := gin.Default()
	//加载指定目录下的html文件
	router.LoadHTMLGlob("templates/**/*")
	//本示例为router基本使用。项目运行时用浏览器访问'http://localhost:8080/info'即可看到效果
	router.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Programming language": "Go",
			"Author":               "Lei Wang",
			"From":                 "BUCT",
		})
	})
	//为无上下文(context)的URL路径配置404NotFound界面。由于本项目不提供html代码，所以代码被注释，使用JSON代替
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "404 Not Found",
		})
		//c.HTML(http.StatusNotFound, "404.html", nil)
	})

	//Router启动
	err := router.Run()
	//处理错误
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
}
