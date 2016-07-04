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
	imageURL = "https://api.cognitive.microsoft.com/bing/v5.0/images/search?q=%s&count=1&mkt=en-us&safeSearch=Off"
)

type ImageResults struct {
	Value []struct {
		ContentURL     string `json:"contentUrl"`
		ContentSize    string `json:"contentSize"`
		EncodingFormat string `json:"encodingFormat"`
	} `json:"value"`
}

func image(command *bot.Cmd, matches []string) (msg string, err error) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", fmt.Sprintf(imageURL, url.QueryEscape(matches[1])), nil)
	request.Header.Set("Ocp-Apim-Subscription-Key", bot.Config.Bing)

	response, _ := client.Do(request)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	var results ImageResults
	json.Unmarshal(body, &results)

	if err != nil {
		return fmt.Sprintf("No results for %s", matches[1]), nil
	}

	if len(results.Value) == 0 {
		return fmt.Sprintf("No results for %s", matches[1]), nil
	}

	output := fmt.Sprintf("Bing | %s | %s", matches[1], results.Value[0].ContentURL)
	return output, nil
}

func init() {
	bot.RegisterCommand(
		"^img (.+)$",
		image)
}
