package gif

import (
	"fmt"
	"github.com/0x263b/porygon2"
	"github.com/0x263b/porygon2/web"
	"math/rand"
	"net/url"
	"time"
)

const (
	giphyURL = "https://api.giphy.com/v1/gifs/search?q=%s&api_key=%s&limit=50"
)

type giphy struct {
	Data []struct {
		Images struct {
			FixedHeight struct {
				Url string `json:"url"`
			} `json:"fixed_height"`
			Original struct {
				Url string `json:"url"`
			} `json:"original"`
		} `json:"images"`
	} `json:"data"`
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
