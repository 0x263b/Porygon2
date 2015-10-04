package weather

import (
	"time"
)

type yahooWeather struct {
	Query struct {
		Count   int       `json:"count"`
		Created time.Time `json:"created"`
		Lang    string    `json:"lang"`
		Results struct {
			Channel struct {
				Title         string `json:"title"`
				Link          string `json:"link"`
				Description   string `json:"description"`
				Language      string `json:"language"`
				Lastbuilddate string `json:"lastBuildDate"`
				TTL           string `json:"ttl"`
				Location      struct {
					City    string `json:"city"`
					Country string `json:"country"`
					Region  string `json:"region"`
				} `json:"location"`
				Units struct {
					Distance    string `json:"distance"`
					Pressure    string `json:"pressure"`
					Speed       string `json:"speed"`
					Temperature string `json:"temperature"`
				} `json:"units"`
				Wind struct {
					Chill     string `json:"chill"`
					Direction string `json:"direction"`
					Speed     string `json:"speed"`
				} `json:"wind"`
				Atmosphere struct {
					Humidity   string `json:"humidity"`
					Pressure   string `json:"pressure"`
					Rising     string `json:"rising"`
					Visibility string `json:"visibility"`
				} `json:"atmosphere"`
				Astronomy struct {
					Sunrise string `json:"sunrise"`
					Sunset  string `json:"sunset"`
				} `json:"astronomy"`
				Image struct {
					Title  string `json:"title"`
					Width  string `json:"width"`
					Height string `json:"height"`
					Link   string `json:"link"`
					URL    string `json:"url"`
				} `json:"image"`
				Item struct {
					Title     string `json:"title"`
					Lat       string `json:"lat"`
					Long      string `json:"long"`
					Link      string `json:"link"`
					Pubdate   string `json:"pubDate"`
					Condition struct {
						Code string `json:"code"`
						Date string `json:"date"`
						Temp string `json:"temp"`
						Text string `json:"text"`
					} `json:"condition"`
					Description string `json:"description"`
					Forecast    []struct {
						Code string `json:"code"`
						Date string `json:"date"`
						Day  string `json:"day"`
						High string `json:"high"`
						Low  string `json:"low"`
						Text string `json:"text"`
					} `json:"forecast"`
					GUID struct {
						Ispermalink string `json:"isPermaLink"`
						Content     string `json:"content"`
					} `json:"guid"`
				} `json:"item"`
			} `json:"channel"`
		} `json:"results"`
	} `json:"query"`
}
