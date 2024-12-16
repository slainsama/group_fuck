package globals

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"group_fuck_slave/models"
	"log"
	"os"
)

var GlobalConfig models.Config

func InitConfig() {
	configFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal(configFile, &GlobalConfig)
	if err != nil {
		log.Println(err)
	}
	log.Println(GlobalConfig)
	whitelist := GlobalConfig.Secure.Whitelist
	fmt.Println("Whitelist:", whitelist)
}
