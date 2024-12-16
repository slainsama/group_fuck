package controllers

import (
	"github.com/gin-gonic/gin"
	"group_fuck_slave/globals"
	"group_fuck_slave/models"
	"net/http"
	"sort"
	"strconv"
)

func GetTasksController(c *gin.Context) {
	limitParam := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		limit = 10
	}
	var tasks []*models.Task
	for _, task := range globals.TaskList {
		tasks = append(tasks, task)
	}
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Time.After(tasks[j].Time)
	})
	if len(tasks) > limit {
		tasks = tasks[:limit]
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}
