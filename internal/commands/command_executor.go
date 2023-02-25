package commands

import "errors"

var DIFFICULTY_MAPPING = map[string]int{
	"Easy":   1,
	"Medium": 2,
	"Hard":   3,
}

var VALID_COMMANDS = [...]string{"!user", "!help", "!claim", "!challenge", "!rank", "!status", "!new-challenge", "!group-status"}

type ICommandExecutor interface {
	validateCommand() string
	parseCommand(message string) *BotCommand
	commandClaim(discordId string, leetcodeId string) string
	commandChallenge() string
	commandRank(leetcodeId string) string
	commandStatus(leetcodeId string) string
	commandNewChallenge() string
	commandUser(leetcodeId string) string
	commandGroupStatus() string
	run() string
}

type CommandExecutor struct{}

func (c *CommandExecutor) parseCommand(message string) *BotCommand {
	return NewBotCommand(message)
}

func (c *CommandExecutor) validateCommand(message string, expectedCommand string) *BotCommand {
	parsedCommand := c.parseCommand(message)
	if parsedCommand.action == expectedCommand {
		parsedCommand.errors = append(parsedCommand.errors, errors.New("Bot command does not match expected command"))
	}

	return parsedCommand
}

func (c *CommandExecutor) commandClaim(discordId string, leetcodeId string) (string, error) {
	// TODO: Implement
	return "1", nil
}

func (c *CommandExecutor) commandChallenge() string {
	// TODO: Implement
	return "1"
}

func (c *CommandExecutor) commandRank(leetcodeId string) string {
	// TODO: Implement
	return "1"
}

func (c *CommandExecutor) commandStatus(leetcodeId string) string {
	// TODO: Implement
	return "1"
}

func (c *CommandExecutor) commandNewChallenge() string {
	// TODO: Implement
	return "1"
}

func (c *CommandExecutor) commandUser(leetcodeId string) string {
	// TODO: Implement
	return "1"
}

func (c *CommandExecutor) commandGroupStatus() string {
	// TODO: Implement
	return "1"
}

func run(c *CommandExecutor) string {
	// TODO: Implement
	return "1"
}
