package gif

import (
	"fmt"
	"github.com/0x263b/Porygon2"
	"github.com/0x263b/Porygon2/web"
	"math/rand"
	"net/url"
	"time"
)

const (
	giphyURL = "https://api.giphy.com/v1/gifs/search?q=%s&api_key=%s&limit=50"
)

type giphy struct {
	Data []struct {
		BitlyUrl string `json:"bitly_url"`
		Images   struct {
			FixedHeight struct {
				Height string `json:"height"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_height"`
			FixedHeightDownsampled struct {
				Height string `json:"height"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_height_downsampled"`
			FixedHeightStill struct {
				Height string `json:"height"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_height_still"`
			FixedWidth struct {
				Height string `json:"height"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_width"`
			FixedWidthDownsampled struct {
				Height string `json:"height"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_width_downsampled"`
			FixedWidthStill struct {
				Height string `json:"height"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"fixed_width_still"`
			Original struct {
				Frames string `json:"frames"`
				Height string `json:"height"`
				Size   string `json:"size"`
				Url    string `json:"url"`
				Width  string `json:"width"`
			} `json:"original"`
		} `json:"images"`
		Type        string `json:"type"`
		Username    string `json:"username"`
		BitlyGifUrl string `json:"bitly_gif_url"`
		EmbedUrl    string `json:"embed_url"`
		Id          string `json:"id"`
		Rating      string `json:"rating"`
		Source      string `json:"source"`
		Url         string `json:"url"`
	} `json:"data"`
	Meta struct {
		Msg    string `json:"msg"`
		Status int64  `json:"status"`
	} `json:"meta"`
	Pagination struct {
		Count      int64 `json:"count"`
		Offset     int64 `json:"offset"`
		TotalCount int64 `json:"total_count"`
	} `json:"pagination"`
}

func gif(command *bot.Cmd, matches []string) (msg string, err error) {
	data := &giphy{}
	err = web.GetJSON(fmt.Sprintf(giphyURL, url.QueryEscape(matches[1]), bot.Config.API.Giphy), data)
	if err != nil {
		return "", err
	}

	if len(data.Data) == 0 {
		return "No gifs found.", nil
	}

	index := rand.Intn(len(data.Data))
	return fmt.Sprintf(data.Data[index].Images.FixedHeight.Url), nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
	bot.RegisterCommand(
		"^gif (.+)$",
		gif)
}
