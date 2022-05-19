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
	closable := cliClient.init(&clientConfig)
	defer closable.Close()

	if err := cliClient.Ping(); err != nil {
		panic("can not start client, ping request is not successful: " + err.Error())
	}

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

	for {
		fmt.Println("1. Registration")
		fmt.Println("2. Get public info")
		fmt.Println("3. Get full info")
		fmt.Println("0. Exit")
		fmt.Println()

		choice := utils.ReadIntValue("Input option number: ")
		cmd, ok := commands[choice]

		if !ok {
			log.Println("Wrong option number")
			continue
		}

		fmt.Println()

		if err := cmd(); err != nil {
			log.Println("Can not perform command: " + err.Error())
			continue
		}
	}
}
