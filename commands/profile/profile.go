package profile

import (
	"fmt"
	"github.com/0x263b/Porygon2"
	"strings"
)

func whoIs(command *bot.Cmd, matches []string) (msg string, err error) {
	nick := matches[1]

	lastfm := bot.GetUserKey(nick, "lastfm")
	if lastfm != "" {
		lastfm = fmt.Sprintf(" | Last.fm: %s", lastfm)
	}

	twitter := bot.GetUserKey(nick, "twitter")
	if twitter != "" {
		twitter = fmt.Sprintf(" | Twitter: @%s", twitter)
	}

	url := bot.GetUserKey(nick, "url")
	if url != "" {
		url = fmt.Sprintf(" | URL: %s", url)
	}

	return fmt.Sprintf("%s%s%s", nick, lastfm, twitter, url), nil
}

func setUrl(command *bot.Cmd, matches []string) (msg string, err error) {
	bot.SetUserKey(command.Nick, "url", strings.TrimSpace(matches[1]))
	return fmt.Sprintf("%s: url updated to: %s", command.Nick, matches[1]), nil
}

func setTwitter(command *bot.Cmd, matches []string) (msg string, err error) {
	bot.SetUserKey(command.Nick, "twitter", strings.TrimSpace(matches[1]))
	return fmt.Sprintf("%s: twitter updated to: %s", command.Nick, matches[1]), nil
}

func init() {
	bot.RegisterCommand("^set url (.+)$", setUrl)
	bot.RegisterCommand("^set twitter (\\S+)$", setTwitter)
	bot.RegisterCommand("^whois (\\S+)$", whoIs)
}
