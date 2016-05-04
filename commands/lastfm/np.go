package lastfm

import (
	"fmt"
	"github.com/0x263b/porygon2"
	"github.com/0x263b/porygon2/web"
	"net/url"
	"strings"
)

const (
	NowPlayingURL = "http://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=%s&limit=1&api_key=%s&format=json"
	ArtistTagsURL = "http://ws.audioscrobbler.com/2.0/?method=artist.gettoptags&artist=%s&api_key=%s&format=json"
)

func nowPlaying(command *bot.Cmd, matches []string) (msg string, err error) {
	username := checkLastfm(command.Nick, matches[1])

	if username == "" {
		return "Lastfm not provided, nor on file. Use `-set lastfm <lastfm>` to save", nil
	}

	data := &NowPlaying{}
	err = web.GetJSON(fmt.Sprintf(NowPlayingURL, username, bot.Config.API.Lastfm), data)
	if err != nil || data.Error > 0 {
		return fmt.Sprintf("Could not get scrobbles for %s", username), nil
	}
	if data.Recenttracks.Attr.Total == "0" {
		return fmt.Sprintf("Could not get scrobbles for %s", username), nil
	}

	album := ""
	if data.Recenttracks.Track[0].Album.Text != "" {
		album = fmt.Sprintf(" from %s", data.Recenttracks.Track[0].Album.Text)
	}

	var fmttags string
	tags := &ArtistTags{}
	err = web.GetJSON(fmt.Sprintf(ArtistTagsURL, url.QueryEscape(data.Recenttracks.Track[0].Artist.Text), bot.Config.API.Lastfm), tags)
	if err != nil {
		return fmt.Sprintf("Could not get scrobbles for %s", username), nil
	}

	var trunc int = 4
	if len(tags.Toptags.Tag) < trunc {
		trunc = len(tags.Toptags.Tag)
	}

	for i := range tags.Toptags.Tag[:trunc] {
		fmttags += fmt.Sprintf("%s, ", tags.Toptags.Tag[i].Name)
	}

	fmttags = strings.TrimSuffix(fmttags, ", ")

	state := "last played"
	if data.Recenttracks.Track[0].Attr.Nowplaying == "true" {
		state = "is playing"
	}

	output := fmt.Sprintf("Last.fm | %s %s: “%s” by %s%s | %s",
		username,
		state,
		data.Recenttracks.Track[0].Name,
		data.Recenttracks.Track[0].Artist.Text,
		album,
		fmttags)

	return output, nil
}

func init() {
	bot.RegisterCommand(
		"^np(?: (\\S+))?$",
		nowPlaying)
}
