package bing

import (
	"encoding/json"
	"fmt"
	"github.com/0x263b/Porygon2"
	"github.com/dustin/go-humanize"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	imageURL = "https://api.datamarket.azure.com/Bing/Search/v1/Image?Query='%s'&Adult='Off'&$format=json"
)

type ImageResults struct {
	D struct {
		Results []struct {
			MediaURL    string `json:"MediaUrl"`
			FileSize    string `json:"FileSize"`
			ContentType string `json:"ContentType"`
		} `json:"results"`
		Next string `json:"__next"`
	} `json:"d"`
}

func image(command *bot.Cmd, matches []string) (msg string, err error) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", fmt.Sprintf(imageURL, url.QueryEscape(matches[1])), nil)
	request.SetBasicAuth(bot.Config.TranslateClient, bot.Config.TranslateSecret)

	response, _ := client.Do(request)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	var results ImageResults
	json.Unmarshal(body, &results)

	if err != nil {
		return fmt.Sprintf("No results for %s", matches[1]), nil
	}

	if len(results.D.Results) == 0 {
		return fmt.Sprintf("No results for %s", matches[1]), nil
	}

	size, _ := strconv.ParseUint(results.D.Results[0].FileSize, 10, 64)
	humanize.Bytes(size)

	output := fmt.Sprintf("Bing | %s → %s %s | %s", matches[1],
		results.D.Results[0].ContentType,
		humanize.Bytes(size),
		results.D.Results[0].MediaURL)
	return output, nil
}

func init() {
	bot.RegisterCommand(
		"^img (.+)$",
		image)
}
