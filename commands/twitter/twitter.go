package twitter

import (
	"encoding/json"
	"fmt"
	"github.com/0x263b/Porygon2"
	"github.com/kurrik/oauth1a"
	"github.com/kurrik/twittergo"
	"net/http"
	"net/url"
)

func LoadCredentials() (client *twittergo.Client, err error) {
	config := &oauth1a.ClientConfig{
		ConsumerKey:    bot.Config.API.TwitterConsumerKey,
		ConsumerSecret: bot.Config.API.TwitterConsumerSecret,
	}
	client = twittergo.NewClient(config, nil)
	return
}

func twitter(command *bot.Cmd, matches []string) (msg string, err error) {
	var (
		client  *twittergo.Client
		query   url.Values
		req     *http.Request
		resp    *twittergo.APIResponse
		results *twittergo.Timeline
		output  string
	)

	if client, err = LoadCredentials(); err != nil {
		return "Twitter | Could not get tweet", nil
	}

	query = url.Values{}
	query.Set("count", "1")
	query.Set("exclude_replies", "true")
	query.Set("screen_name", matches[1])

	endpoint := fmt.Sprintf("/1.1/statuses/user_timeline.json?%v", query.Encode())

	if req, err = http.NewRequest("GET", endpoint, nil); err != nil {
		return "Twitter | Could not get tweet", nil
	}
	if resp, err = client.SendRequest(req); err != nil {
		return "Twitter | Could not get tweet", nil
	}
	results = &twittergo.Timeline{}
	resp.Parse(results)
	for _, tweet := range *results {
		if _, err = json.Marshal(tweet); err != nil {
			return "Twitter | Could not get tweet", nil
		}
		output = fmt.Sprintf("Twitter | %s (@%s) | %s | %s",
			tweet.User().Name(),
			tweet.User().ScreenName(),
			tweet.Text(),
			Time(tweet.CreatedAt()))
	}

	return output, nil
}

func init() {
	bot.RegisterCommand(
		"^tw(?:itter)? (\\S+)$",
		twitter)
}
