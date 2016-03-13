package google

import (
	"fmt"
	"github.com/0x263b/Porygon2"
	"github.com/0x263b/Porygon2/web"
	"github.com/kennygrant/sanitize"
	"net/url"
)

const (
	googleURL = "https://ajax.googleapis.com/ajax/services/search/web?v=1.0&q=%s"
)

type SearchResults struct {
	Responsedata struct {
		Results []struct {
			URL               string `json:"url"`
			Title             string `json:"title"`
			Titlenoformatting string `json:"titleNoFormatting"`
		} `json:"results"`
	} `json:"responseData"`
}

func google(command *bot.Cmd, matches []string) (msg string, err error) {
	results := &SearchResults{}
	err = web.GetJSON(fmt.Sprintf(googleURL, url.QueryEscape(matches[1])), results)
	if err != nil {
		return fmt.Sprintf("No results for %s", matches[1]), nil
	}

	if len(results.Responsedata.Results) == 0 {
		return fmt.Sprintf("No results for %s", matches[1]), nil
	}

	title := sanitize.HTML(results.Responsedata.Results[0].Title)
	link, _ := url.QueryUnescape(results.Responsedata.Results[0].URL)

	output := fmt.Sprintf("Google | %s | %s", title, link)
	return output, nil
}

func init() {
	bot.RegisterCommand(
		"^g(?:oogle)? (.+)$",
		google)
}
