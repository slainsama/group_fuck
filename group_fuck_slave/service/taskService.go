package service

import (
	"fmt"
	"group_fuck_slave/globals"
	"group_fuck_slave/utils"
	"log"
)

func InitTaskService() {
	ListenToTaskChannel()
	log.Println("Task service start...")
}

func ListenToTaskChannel() {
	go func() {
		fmt.Println("ListenToTaskChannel...")
		for {
			select {
			case newTaskId := <-globals.TaskChannel:
				go handleTask(newTaskId)
			}
		}
	}()
}

func handleTask(newTaskId string) {
	result, err := utils.RunShellCommand(globals.TaskList[newTaskId].CmdString)
	if err != nil {
		globals.TaskList[newTaskId].Out = err.Error()
	} else {
		globals.TaskList[newTaskId].Out = result
	}
	globals.TaskList[newTaskId].IsFinish = true
}
