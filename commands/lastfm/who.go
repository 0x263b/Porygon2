package lastfm

import (
	"fmt"
	"github.com/0x263b/Porygon2"
	"github.com/0x263b/Porygon2/web"
	"net/url"
	"strings"
	"time"
)

func whosPlaying(command *bot.Cmd, matches []string) (msg string, err error) {
	users := bot.GetNames(strings.ToLower(command.Channel))

	var playing []string

	for _, user := range users {
		if bot.GetUserKey(user, "lastfm") != "" {
			playing = append(playing, user)
		}
	}

	for _, user := range playing {
		username := checkLastfm(user, user)

		data := &NowPlaying{}
		err = web.GetJSON(fmt.Sprintf(NowPlayingURL, username, bot.Config.API.Lastfm), data)
		if err != nil || data.Error > 0 {
			continue
		}
		if data.Recenttracks.Attr.Total == "0" {
			continue
		}

		if data.Recenttracks.Track[0].Attr.Nowplaying != "true" {
			continue
		}

		var fmttags string
		tags := &ArtistTags{}
		err = web.GetJSON(fmt.Sprintf(ArtistTagsURL, url.QueryEscape(data.Recenttracks.Track[0].Artist.Text), bot.Config.API.Lastfm), tags)
		if err != nil {
			continue
		}

		var trunc int = 4
		if len(tags.Toptags.Tag) < trunc {
			trunc = len(tags.Toptags.Tag)
		}

		for i := range tags.Toptags.Tag[:trunc] {
			fmttags += fmt.Sprintf("%s, ", tags.Toptags.Tag[i].Name)
		}

		fmttags = strings.TrimSuffix(fmttags, ", ")

		bot.Conn.Privmsg(command.Channel, fmt.Sprintf("%s (%s) | “%s” by %s | %s",
			user,
			username,
			data.Recenttracks.Track[0].Name,
			data.Recenttracks.Track[0].Artist.Text,
			fmttags))

		time.Sleep(500 * time.Millisecond)
	}

	return "", nil
}

func init() {
	bot.RegisterCommand(
		"^wp$",
		whosPlaying)
}
