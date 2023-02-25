package commands

import "strings"

type BotCommand struct {
	action string
	args   []string
	result string
	errors []error
}

func NewBotCommand(message string) *BotCommand {
	b := new(BotCommand)

	messageParts := strings.Split(message, " ")
	b.action = messageParts[0]
	b.args = messageParts[1:]

	return b
}
