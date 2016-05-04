package google

import (
	"github.com/0x263b/Porygon2"
)

func google(command *bot.Cmd, matches []string) (msg string, err error) {
	return "This command is deprecated. https://ajax.googleapis.com/ajax/services/search/web?v=1.0&q=deprecated", nil
}

func init() {
	bot.RegisterCommand(
		"^g(?:oogle)? (.+)$",
		google)
}
