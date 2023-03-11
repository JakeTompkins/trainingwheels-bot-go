package commands

import (
	"errors"
	"strings"
)

type BotCommand struct {
	Action string
	Args   []string
	Result string
}

func NewBotCommand(message string) (*BotCommand, error) {
	b := new(BotCommand)

	messageParts := strings.Split(message, " ")
	b.Action = messageParts[0]
	b.Args = messageParts[1:]

	for _, validCommand := range VALID_COMMANDS {
		if strings.Compare(b.Action, validCommand) == 0 {
			return b, nil
		}
	}

	return nil, errors.New("Invalid command, type !help for valid commands")
}
