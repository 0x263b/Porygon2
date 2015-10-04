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
	clientID     = ""
	clientSecret = ""
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
	parameters.Add("client_id", clientID)
	parameters.Add("client_secret", clientSecret)
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
	transURL, _ := url.Parse(translateURL)
	parameters := url.Values{}
	parameters.Add("from", matches[1])
	parameters.Add("to", matches[2])
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
		return fmt.Sprintf("Translate | %s >> %s | Could not get translation", matches[1], matches[2]), nil
	}
	return fmt.Sprintf("Translate | %s >> %s | %s", matches[1], matches[2], dict.Translation), nil
}

func init() {
	bot.RegisterCommand(
		"^tr(?:anslate)? (\\w+) (\\w+) (.+)$",
		translate)
}
