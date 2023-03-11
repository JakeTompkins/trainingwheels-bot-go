package main

import (
	"fmt"
	"os"

	cli_client "github.com/jaketompkins/trainingwheels-bot-go/internal/clients/cli_client"
	"github.com/jaketompkins/trainingwheels-bot-go/internal/commands"
	database "github.com/jaketompkins/trainingwheels-bot-go/internal/database"
)

func main() {
	database.InitUserCollection()
	database.InitQuestionCollection()
	database.InitChallengeCollection()
	database.InitChallengeQuestionCollection()

	switch runMode := os.Args[1]; runMode {
	case "--local":
		{
			client, err := cli_client.NewCliClient()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println("Running in local mode")
			err = client.Run()

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	case "--discord":
		{
			os.Exit(0)
		}
	case "--init":
		{
			res, err := commands.PopulateDb()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println(res)
			os.Exit(0)
		}
	default:
		{
			fmt.Println("Invalid argument, valid arguments are --local, --discord, and --init")
		}

	}
}
