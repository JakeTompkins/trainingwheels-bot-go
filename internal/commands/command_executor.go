package commands

import "errors"

var DIFFICULTY_MAPPING = map[string]int{
	"Easy":   1,
	"Medium": 2,
	"Hard":   3,
}

var VALID_COMMANDS = [...]string{"!user", "!help", "!claim", "!challenge", "!rank", "!status", "!new-challenge", "!group-status"}

type ICommandController interface {
	validateCommand() string
	parseCommand(message string) *BotCommand
	commandClaim(discordId string, leetcodeId string) string
	commandChallenge() string
	commandRank(leetcodeId string) string
	commandStatus(leetcodeId string) string
	commandNewChallenge() string
	commandUser(leetcodeId string) string
	commandGroupStatus() string
}

type CommandController struct{}

func (c *CommandController) parseCommand(message string) *BotCommand {
	return NewBotCommand(message)
}

func (c *CommandController) validateCommand(message string, expectedCommand string) *BotCommand {
	parsedCommand := c.parseCommand(message)
	if parsedCommand.action == expectedCommand {
		parsedCommand.errors = append(parsedCommand.errors, errors.New("Bot command does not match expected command"))
	}

	return parsedCommand
}

func (c *CommandController) commandClaim(discordId string, leetcodeId string) string {
	// TODO: Implement
	return "1"
}

func (c *CommandController) commandChallenge() string {
	// TODO: Implement
	return "1"
}

func (c *CommandController) commandRank(leetcodeId string) string {
	// TODO: Implement
	return "1"
}

func (c *CommandController) commandStatus(leetcodeId string) string {
	// TODO: Implement
	return "1"
}

func (c *CommandController) commandNewChallenge() string {
	// TODO: Implement
	return "1"
}

func (c *CommandController) commandUser(leetcodeId string) string {
	// TODO: Implement
	return "1"
}

func (c *CommandController) commandGroupStatus() string {
	// TODO: Implement
	return "1"
}
