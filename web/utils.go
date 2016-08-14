package web

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Short struct {
	URL struct {
		ID             string    `json:"id"`
		ShortURL       string    `json:"short_url"`
		LongURL        string    `json:"long_url"`
		ExpirationDate time.Time `json:"expiration_date"`
	} `json:"url"`
	Status string `json:"status"`
}

func GetBody(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func GetJSON(url string, v interface{}) error {
	body, err := GetBody(url)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

func ShortenURL(url string) string {
	client := &http.Client{}
	request, _ := http.NewRequest("POST", "https://mnn.im/s", bytes.NewBufferString(url))
	response, _ := client.Do(request)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	var short Short
	json.Unmarshal(body, &short)
	return short.URL.ShortURL
}
