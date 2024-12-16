package utils

import (
	"github.com/google/uuid"
	"group_fuck_slave/globals"
	"group_fuck_slave/models"
	"time"
)

func CreateTask(data string) (taskId string) {
	var newTask models.Task
	newId := uuid.New().String()
	newTask.Id = newId
	newTask.IsFinish = false
	newTask.CmdString = data
	newTask.Time = time.Now()
	globals.TaskList[newId] = &newTask
	globals.TaskChannel <- newTask.Id
	return newId
}
