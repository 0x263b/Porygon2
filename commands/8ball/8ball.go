package eightBall

import (
	"fmt"
	"github.com/0x263b/Porygon2"
	"math/rand"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func eightBall(command *bot.Cmd, matches []string) (msg string, err error) {
	choices := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes - definitely",
		"You may rely on it",
		"As I see it, yes",
		"Most likely",
		"Outlook good",
		"Signs point to yes",
		"Yes",
		"Better not tell you now",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
		"Don't care, go away",
		"Seeking what is true is not seeking what is desirable",
		"Man is nothing else but what he makes of himself",
		"There is no reality except in action",
		"It's always better when you discover answers on your own",
	}
	chosen := random(0, len(choices))
	msg = fmt.Sprintf("%s: %s", command.Nick, choices[chosen])
	return
}

func init() {
	bot.RegisterCommand(
		"^8ball",
		eightBall)
}
