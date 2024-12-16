package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"group_fuck_slave/controllers"
	"group_fuck_slave/globals"
	"group_fuck_slave/middles"
)

var Server *gin.Engine

func InitServer() {
	Server = gin.Default()
	Server.Use(gin.Logger())
	Server.Use(gin.Recovery())
	loadRouter(Server)
	fmt.Printf("Server running on port: %s\n", globals.GlobalConfig.Server.Host)
}

func loadRouter(router *gin.Engine) {
	router.POST("/exec", middles.AuthWithKey(), middles.AuthWithWhiteList(), controllers.ExecController)
	router.GET("/getTasks", middles.AuthWithKey(), middles.AuthWithWhiteList(), controllers.GetTasksController)
}
