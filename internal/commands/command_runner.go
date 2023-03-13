package commands

import (
	"errors"

	"github.com/jaketompkins/trainingwheels-bot-go/internal/database"
	leetcodeClient "github.com/jaketompkins/trainingwheels-bot-go/internal/leetcode"
)

var DIFFICULTY_MAPPING = map[string]int{
	"Easy":   1,
	"Medium": 2,
	"Hard":   3,
}

var VALID_COMMANDS = [...]string{
	"!user",
	"!help",
	"!claim",
	"!challenge",
	"!rank",
	"!status",
	"!new-challenge",
	"!group-status",
}

func ParseCommand(message string) (*BotCommand, error) {
	return NewBotCommand(message)
}

func ExecuteCommand(botCommand *BotCommand) (string, error) {
	switch botCommand.Action {
	case "!help":
		return commandHelp(), nil
	}

	return "", errors.New("Invalid command")
}

func PopulateDb() (string, error) {
	questions, err := leetcodeClient.LoadAllQuestions()
	if err != nil {
		return "", err
	}

	err = database.InsertQuestions(questions)
	if err != nil {
		return "", err
	}

	return "Leetcode questions successfully downloaded", nil
}

func NewChallenge() {
}

func commandHelp() string {
	return `
	"!user"				Get user's leetcode stats
	"!help"				Show valid commands
	"!claim"			Associate discord user with leetcode user
	"!challenge"		Show the current leetcode challenge
	"!rank"				Show the user's leetcode rank
	"!status"			Show the user's progress in the current leetcode challenge
	"!new-challenge"	Generate a new leetcode challenge
	"!group-status"		Show the progress of all users in the current leetcode challenge
	`
}

// NOTE: This probably should not be run in local mode
func commandClaim(discordId string, leetcodeId string) string {
	// TODO: Implement
	return "1"
}

func commandChallenge() string {
	// TODO: Implement
	return "1"
}

func commandRank(leetcodeId string) string {
	// TODO: Implement
	return "1"
}

func commandStatus(leetcodeId string) string {
	// TODO: Implement
	return "1"
}

func commandNewChallenge() string {
	// TODO: Implement
	return "1"
}

func commandUser(leetcodeId string) string {
	// TODO: Implement
	return "1"
}

func commandGroupStatus() string {
	// TODO: Implement
	return "1"
}
