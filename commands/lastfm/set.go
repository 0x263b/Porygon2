package lastfm

import (
	"fmt"
	"github.com/0x263b/porygon2"
	"strings"
)

func checkLastfm(senderNick string, queryNick string) string {
	lastfm := ""
	if queryNick == "" {
		lastfm = bot.GetUserKey(senderNick, "lastfm")
	} else {
		lastfm = bot.GetUserKey(queryNick, "lastfm")
		if lastfm == "" {
			lastfm = queryNick
		}
	}
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
