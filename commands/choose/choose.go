package choose

import (
	"fmt"
	"github.com/killwhitey/Porygon2"
	"math/rand"
	"strings"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func choose(command *bot.Cmd, matches []string) (msg string, err error) {
	splits := strings.Split(matches[1], "|")
	chosen := random(0, len(splits))
	msg = fmt.Sprintf("%s: %s", command.Nick, splits[chosen])
	return
}

func init() {
	bot.RegisterCommand(
		"^r(?:and)? (.+)$",
		choose)
}
