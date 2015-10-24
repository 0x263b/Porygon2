package lastfm

import (
	"fmt"
	"github.com/0x263b/Porygon2"
	"github.com/0x263b/Porygon2/web"
	"strings"
)

func whosPlaying(command *bot.Cmd, matches []string) (msg string, err error) {
	users := bot.GetNames(strings.ToLower(command.Channel))

	var playing []string

	for index, user := range users {
		if bot.GetUserKey(user, "lastfm") != "" {
			playing = append(playing, user)
		}
	}

	for _, user := range playing {
		username := checkLastfm(user)

		data := &NowPlaying{}
		err = web.GetJSON(fmt.Sprintf(NowPlayingURL, username, bot.API.Lastfm), data)
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
		if len(data.Recenttracks.Track[0].Artist.Mbid) > 10 {
			tags := &ArtistTags{}
			err = web.GetJSON(fmt.Sprintf(ArtistTagsURL, data.Recenttracks.Track[0].Artist.Mbid, bot.API.Lastfm), tags)
			if err != nil {
				continue
			}

			for i := range tags.Toptags.Tag[:4] {
				fmttags += fmt.Sprintf("%s, ", tags.Toptags.Tag[i].Name)
			}

			fmttags = strings.TrimSuffix(fmttags, ", ")
		}

		bot.Conn.Privmsg(command.Channel, fmt.Sprintf("%s (%s): “%s” by %s | %s",
			user,
			username,
			data.Recenttracks.Track[0].Name,
			data.Recenttracks.Track[0].Artist.Text,
			fmttags))
	}

	return "", nil
}

func init() {
	bot.RegisterCommand(
		"^wp$",
		whosPlaying)
}
