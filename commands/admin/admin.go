package admin

import (
	"fmt"
	"github.com/0x263b/porygon2"
	"github.com/steveyen/gkvlite"
	"strings"
)

const (
	helpURL = "https://github.com/0x263b/Porygon2/blob/master/USAGE.md"
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

func listChannels(command *bot.Cmd, matches []string) (msg string, err error) {
	if !bot.IsAdmin(command.Nick) || !bot.IsPrivateMsg(command.Channel, command.Nick) {
		return
	}
	output := "I'm in:"

	bot.Channels.VisitItemsAscend([]byte(""), true, func(i *gkvlite.Item) bool {
		if bot.GetChannelKey(string(i.Key), "auto_join") == true {
			output = fmt.Sprintf("%s %s", output, string(i.Key))
		}
		return true
	})

	return output, nil
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

	bot.RegisterCommand(
		"^list channels$",
		listChannels)
}
