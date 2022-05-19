package main

import (
	"flag"
	"fmt"
	"kleptophobia/models"
	"kleptophobia/utils"
	"log"
	"os"
)

func main() {
	configFilename := flag.String("config", "dev_config.json", "client config")
	flag.Parse()

	var clientConfig models.ClientConfig
	utils.InitConfig[*models.ClientConfig](*configFilename, &clientConfig)

	var cliClient CliClient
	cliClient.init(&clientConfig)

	var commands = map[int64]func() error{
		1: cliClient.Register,
		2: cliClient.GetPublicInfo,
		3: cliClient.GetFullInfo,
		0: func() error {
			fmt.Println("Exit!")
			os.Exit(0)
			return nil
		},
	}

	fmt.Println(commands)
	for {
		fmt.Println("1. Registration")
		fmt.Println("2. Get public info")
		fmt.Println("3. Get full info")
		fmt.Println("0. Exit")

		choice := utils.ReadIntValue("Input option number: ")
		cmd, ok := commands[choice]

		if !ok {
			log.Println("Wrong option number")
			continue
		}

		if err := cmd(); err != nil {
			log.Println("Can not perform command: " + err.Error())
			continue
		}
	}
}
