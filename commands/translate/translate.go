package translate

import (
	"encoding/json"
	"fmt"
	"github.com/0x263b/porygon2"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	translateURL = "https://api.datamarket.azure.com/Bing/MicrosoftTranslator/v1/Translate?Text='%s'&To='%s'&From='%s'&$format=json"
)

type Translation struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Text string `json:"Text"`
		} `json:"results"`
	} `json:"d"`
}

func translate(command *bot.Cmd, matches []string) (msg string, err error) {

	from := matches[1]
	to := matches[2]

	if from == "cn" {
		from = "zh-CHS"
	} else if from == "tw" {
		from = "zh-CHT"
	} else if to == "cn" {
		to = "zh-CHS"
	} else if to == "tw" {
		to = "zh-CHT"
	}

	client := &http.Client{}
	request, _ := http.NewRequest("GET", fmt.Sprintf(translateURL, url.QueryEscape(matches[3]), to, from), nil)
	request.SetBasicAuth(bot.Config.TranslateClient, bot.Config.TranslateSecret)

	response, _ := client.Do(request)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	var translation Translation
	json.Unmarshal(body, &translation)

	if err != nil {
		return fmt.Sprintf("Translate | %s >> %s | Could not get translation", from, to), nil
	}

	if len(translation.D.Results) == 0 {
		return fmt.Sprintf("Translate | %s >> %s | Could not get translation", from, to), nil
	}

	return fmt.Sprintf("Translate | %s >> %s | %s", from, to, translation.D.Results[0].Text), nil
}

func init() {
	bot.RegisterCommand(
		"^tr(?:anslate)? (\\w+) (\\w+) (.+)$",
		translate)
}
