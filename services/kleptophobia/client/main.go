package main

import (
	"fmt"
	"kleptophobia/utils"
	"log"
)

type Command = func(cliClient *CliClient) error

func main() {
	var cliClient CliClient
	cliClient.init("localhost:50051")

	var commands = map[int64]func() error{
		1: cliClient.Register,
		2: cliClient.GetPublicInfo,
	}

	fmt.Println(commands)
	for {
		fmt.Println("1. Registration")
		fmt.Println("2. Get public info")
		fmt.Println("0. Exit")

		choice := utils.ReadIntValue("Input option number: ")
		if choice == 0 {
			break
		}
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
