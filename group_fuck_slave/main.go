package main

import (
	"github.com/fvbock/endless"
	"group_fuck_slave/globals"
	"group_fuck_slave/server"
	"group_fuck_slave/service"
	"log"
)

func init() {
	globals.InitConfig()
	globals.InitTaskList()
	server.InitServer()
	service.InitTaskService()
}

func main() {
	endlessServer := endless.NewServer(globals.GlobalConfig.Server.Host, server.Server)
	err := endlessServer.ListenAndServe()
	if err != nil {
		log.Println("something wrong with starting.")
	}
}
