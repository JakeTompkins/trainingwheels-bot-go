package cli_client

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	commands "github.com/jaketompkins/trainingwheels-bot-go/internal/commands"
)

type CLIClient struct {
	LeetcodeId string
}

func NewCliClient() (*CLIClient, error) {
	leetcodeId := os.Getenv("LEETCODE_ID")

	if leetcodeId == "" {
		return nil, errors.New("LEETCODE_ID not found in env variables")
	}

	return &CLIClient{leetcodeId}, nil
}

func (c *CLIClient) Run() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("->")
		input, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		botCommand, err := commands.ParseCommand(input)
		if err != nil {
			return err
		}

		res, err := commands.ExecuteCommand(botCommand)
		if err != nil {
			return err
		}

		fmt.Println(res)
	}
}
