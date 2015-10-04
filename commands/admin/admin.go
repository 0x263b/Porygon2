package admin

import (
	"fmt"
	"github.com/killwhitey/Porygon2"
	"strings"
)

const (
	helpURL = "https://github.com/killwhitey/Porygon2#functions"
)

func help(command *bot.Cmd, matches []string) (msg string, err error) {
	return fmt.Sprintf("%s: %s", command.Nick, helpURL), nil
}

func setIgnore(command *bot.Cmd, matches []string) (msg string, err error) {
	if !bot.IsAdmin(command.Nick) || !bot.IsPrivateMsg(command.Channel, command.Nick) {
		return
	}
	bot.SetUserKey(strings.TrimSpace(matches[1]), "ignore", "true")
	return "I never liked him anyway", nil
}

func setUnignore(command *bot.Cmd, matches []string) (msg string, err error) {
	if !bot.IsAdmin(command.Nick) || !bot.IsPrivateMsg(command.Channel, command.Nick) {
		return
	}
	bot.DeleteUserKey(strings.TrimSpace(matches[1]), "ignore")
	return "Sorry about that", nil
}

func init() {
	bot.RegisterCommand("^help",
		help)

	bot.RegisterCommand(
		"^set ignore (\\S+)$",
		setIgnore)

	bot.RegisterCommand(
		"^set unignore (\\S+)$",
		setUnignore)
}
