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
	searchURL = "https://api.datamarket.azure.com/Bing/Search/v1/Web?Query='%s'&Options='DisableLocationDetection'&Market='en-US'&$format=json"
)

type SearchResults struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ID          string `json:"ID"`
			Title       string `json:"Title"`
			Description string `json:"Description"`
			DisplayURL  string `json:"DisplayUrl"`
			URL         string `json:"Url"`
		} `json:"results"`
		Next string `json:"__next"`
	} `json:"d"`
}

func bing(command *bot.Cmd, matches []string) (msg string, err error) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", fmt.Sprintf(searchURL, url.QueryEscape(matches[1])), nil)
	request.SetBasicAuth(bot.Config.TranslateClient, bot.Config.TranslateSecret)

	response, _ := client.Do(request)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	var results SearchResults
	json.Unmarshal(body, &results)

	if err != nil {
		return fmt.Sprintf("No results for %s", matches[1]), nil
	}

	if len(results.D.Results) == 0 {
		return fmt.Sprintf("No results for %s", matches[1]), nil
	}

	output := fmt.Sprintf("Bing | %s | %s ",
		results.D.Results[0].Title,
		results.D.Results[0].URL)
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
