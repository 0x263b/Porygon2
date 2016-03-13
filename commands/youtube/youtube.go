package youtube

import (
	"fmt"
	"github.com/0x263b/Porygon2"
	"github.com/0x263b/Porygon2/web"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

const (
	youtubeSearchURL = "https://www.googleapis.com/youtube/v3/search?part=snippet&type=video&q=%s&key=%s"
	youtubeVideoURL  = "https://www.googleapis.com/youtube/v3/videos?part=snippet,contentDetails,statistics&id=%s&key=%s"
)

type youtubeSearch struct {
	Pageinfo struct {
		Totalresults int `json:"totalResults"`
	} `json:"pageInfo"`
	Items []struct {
		ID struct {
			Videoid string `json:"videoId"`
		} `json:"id"`
	} `json:"items"`
}

type youtubeVideo struct {
	Items []struct {
		Snippet struct {
			Title string `json:"title"`
		} `json:"snippet"`
		Contentdetails struct {
			Duration string `json:"duration"`
		} `json:"contentDetails"`
	} `json:"items"`
}

func ParseDuration(str string) time.Duration {
	durationRegex := regexp.MustCompile(`P(?P<years>\d+Y)?(?P<months>\d+M)?(?P<days>\d+D)?T?(?P<hours>\d+H)?(?P<minutes>\d+M)?(?P<seconds>\d+S)?`)
	matches := durationRegex.FindStringSubmatch(str)

	years := ParseInt64(matches[1])
	months := ParseInt64(matches[2])
	days := ParseInt64(matches[3])
	hours := ParseInt64(matches[4])
	minutes := ParseInt64(matches[5])
	seconds := ParseInt64(matches[6])

	hour := int64(time.Hour)
	minute := int64(time.Minute)
	second := int64(time.Second)
	return time.Duration(years*24*365*hour + months*30*24*hour + days*24*hour + hours*hour + minutes*minute + seconds*second)
}

func ParseInt64(value string) int64 {
	if len(value) == 0 {
		return 0
	}
	parsed, err := strconv.Atoi(value[:len(value)-1])
	if err != nil {
		return 0
	}
	return int64(parsed)
}

func youtube(command *bot.Cmd, matches []string) (msg string, err error) {
	search := &youtubeSearch{}
	err = web.GetJSON(fmt.Sprintf(youtubeSearchURL, url.QueryEscape(matches[1]), bot.Config.API.Youtube), search)
	if err != nil {
		return fmt.Sprintf("YouTube | Could not find video for: %s", matches[1]), nil
	}

	if search.Pageinfo.Totalresults == 0 {
		return fmt.Sprintf("YouTube | Could not find video for: %s", matches[1]), nil
	}

	id := search.Items[0].ID.Videoid

	video := &youtubeVideo{}
	err = web.GetJSON(fmt.Sprintf(youtubeVideoURL, id, bot.Config.API.Youtube), video)
	if err != nil {
		return fmt.Sprintf("YouTube | Could not find video for: %s", matches[1]), nil
	}

	reg := regexp.MustCompile("\\s+")
	title := video.Items[0].Snippet.Title
	title = reg.ReplaceAllString(title, " ") // Strip excessive spaces

	duration := ParseDuration(video.Items[0].Contentdetails.Duration)

	output := fmt.Sprintf("YouTube | %s | %s | https://youtu.be/%s", title, duration, id)

	return output, nil
}

func init() {
	bot.RegisterCommand(
		"^yt (.+)$",
		youtube)
}
