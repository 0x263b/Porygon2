package translate

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/0x263b/Porygon2"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	authURL      = "https://datamarket.accesscontrol.windows.net/v2/OAuth2-13"
	translateURL = "https://api.microsofttranslator.com/v2/Http.svc/Translate"
)

type Authorization struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	Scope       string `json:"scope"`
}

type String struct {
	Translation string `xml:",chardata"`
	Namespace   string `xml:"xmlns,attr"`
}

func auth() string {
	parameters := url.Values{}
	parameters.Add("grant_type", "client_credentials")
	parameters.Add("client_id", bot.Config.API.TranslateClient)
	parameters.Add("client_secret", bot.Config.API.TranslateSecret)
	parameters.Add("scope", "http://api.microsofttranslator.com")

	client := &http.Client{}
	request, _ := http.NewRequest("POST", authURL, bytes.NewBufferString(parameters.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, _ := client.Do(request)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	var token Authorization
	json.Unmarshal(body, &token)
	return token.AccessToken
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

	transURL, _ := url.Parse(translateURL)
	parameters := url.Values{}
	parameters.Add("from", from)
	parameters.Add("to", to)
	parameters.Add("text", matches[3])
	transURL.RawQuery = parameters.Encode()

	token := auth()

	authorizationHeader := fmt.Sprintf("Bearer %s", token)

	client := &http.Client{}
	request, _ := http.NewRequest("GET", transURL.String(), nil)
	request.Header.Add("Authorization", authorizationHeader)

	response, _ := client.Do(request)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	var dict String
	xml.Unmarshal(body, &dict)

	if dict.Translation == "" {
		return fmt.Sprintf("Translate | %s >> %s | Could not get translation", from, to), nil
	}
	return fmt.Sprintf("Translate | %s >> %s | %s", from, to, dict.Translation), nil
}

func init() {
	bot.RegisterCommand(
		"^tr(?:anslate)? (\\w+) (\\w+) (.+)$",
		translate)
}
