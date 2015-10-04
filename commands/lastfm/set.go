package lastfm

import (
	"fmt"
	"github.com/killwhitey/Porygon2"
	"strings"
)

func checkLastfm(senderNick string) string {
	lastfm := bot.GetUserKey(senderNick, "lastfm")
	return string(lastfm)
}

func setLastfm(command *bot.Cmd, matches []string) (msg string, err error) {
	bot.SetUserKey(command.Nick, "lastfm", strings.TrimSpace(matches[1]))
	return fmt.Sprintf("%s: lastfm updated to: %s", command.Nick, matches[1]), nil
}

func init() {
	bot.RegisterCommand(
		"^set lastfm (.+)$",
		setLastfm)
}
