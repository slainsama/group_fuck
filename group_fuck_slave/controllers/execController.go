package controllers

import (
	"github.com/gin-gonic/gin"
	"group_fuck_slave/utils"
	"net/http"
)

func ExecController(context *gin.Context) {
	data := context.PostForm("data")
	if data == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "need data!"})
		return
	}
	decryptData, err := utils.DecryptData(data)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	newTaskId := utils.CreateTask(decryptData)
	context.JSON(http.StatusOK, gin.H{"Id": newTaskId})
}
