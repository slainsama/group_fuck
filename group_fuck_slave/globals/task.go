package globals

import "group_fuck_slave/models"

var TaskList map[string]*models.Task

var TaskChannel chan string

func InitTaskList() {
	TaskList = make(map[string]*models.Task)
	TaskChannel = make(chan string)
}
