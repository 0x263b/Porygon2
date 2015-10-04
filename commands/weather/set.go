package weather

import (
	"fmt"
	"github.com/0x263b/Porygon2"
	"strings"
)

func checkLocation(senderNick string) string {
	location := bot.GetUserKey(senderNick, "location")
	return string(location)
}

func setLocation(command *bot.Cmd, matches []string) (msg string, err error) {
	bot.SetUserKey(command.Nick, "location", strings.TrimSpace(matches[1]))
	return fmt.Sprintf("%s: Location updated to: %s", command.Nick, matches[1]), nil
}

func init() {
	bot.RegisterCommand(
		"^set location (.+)$",
		setLocation)
}
