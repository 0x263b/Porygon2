package bing

import (
	"encoding/json"
	"fmt"
	"github.com/0x263b/porygon2"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	searchURL = "https://api.cognitive.microsoft.com/bing/v5.0/search?q=%s&count=1&mkt=en-us&responseFilter=Webpages"
)

type SearchResults struct {
	WebPages *struct {
		Value []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"value"`
	} `json:"webPages"`
	Error struct {
		StatusCode int    `json:"statusCode"`
		Message    string `json:"message"`
	} `json:"error"`
}

func bing(command *bot.Cmd, matches []string) (msg string, err error) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", fmt.Sprintf(searchURL, url.QueryEscape(matches[1])), nil)
	request.Header.Set("Ocp-Apim-Subscription-Key", bot.Config.Bing)

	response, _ := client.Do(request)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	var results SearchResults
	json.Unmarshal(body, &results)

	if err != nil {
		return fmt.Sprintf("No results for %s", matches[1]), nil
	}

	if results.WebPages == nil {
		return fmt.Sprintf("No results for %s", matches[1]), nil
	}

	title := results.WebPages.Value[0].Name
	pageURL := results.WebPages.Value[0].URL

	transport := http.Transport{}

	request, _ = http.NewRequest("HEAD", pageURL, nil)
	response, _ = transport.RoundTrip(request)
	pageURL = response.Header.Get("Location")

	output := fmt.Sprintf("Bing | %s | %s ", title, pageURL)
	return output, nil
}

func init() {
	bot.RegisterCommand(
		"^b(?:ing)? (.+)$",
		bing)

	bot.RegisterCommand(
		"^g(?:oogle)? (.+)$",
		bing)
}
