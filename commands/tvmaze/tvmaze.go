package tvmaze

import (
	"fmt"
	"github.com/0x263b/porygon2"
	"github.com/0x263b/porygon2/web"
	"net/url"
)

const (
	tvMazeURL = "http://api.tvmaze.com/singlesearch/shows?q=%s"
)

type Showinfo struct {
	Name     string `json:"name"`
	Status   string `json:"status"`
	Schedule struct {
		Time string   `json:"time"`
		Days []string `json:"days"`
	} `json:"schedule"`
	Network struct {
		Name string `json:"name"`
	} `json:"network"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Previousepisode struct {
			Href string `json:"href"`
		} `json:"previousepisode"`
		Nextepisode struct {
			Href string `json:"href"`
		} `json:"nextepisode"`
	} `json:"_links"`
}

type Nextepisode struct {
	Season  int    `json:"season"`
	Number  int    `json:"number"`
	Airdate string `json:"airdate"`
	Airtime string `json:"airtime"`
}

func tvmaze(command *bot.Cmd, matches []string) (msg string, err error) {
	results := &Showinfo{}
	err = web.GetJSON(fmt.Sprintf(tvMazeURL, url.QueryEscape(matches[1])), results)
	if err != nil {
		return "TVmaze | Could not find show", nil
	}

	if len(results.Links.Nextepisode.Href) != 0 {
		next := &Nextepisode{}
		err = web.GetJSON(results.Links.Nextepisode.Href, next)
		if err != nil {
			return "TVmaze | Could not find show", nil
		}

		if len(results.Schedule.Days) == 0 {
			results.Schedule.Days = []string{"???"}
		}

		output := fmt.Sprintf("TVmaze | %s | Airtime: %s %s on %s | Status: %s | Next Ep: S%vE%v at %s %s",
			results.Name,
			results.Schedule.Days[0],
			results.Schedule.Time,
			results.Network.Name,
			results.Status,
			next.Season,
			next.Number,
			next.Airtime,
			next.Airdate,
		)
		return output, nil
	}

	if len(results.Schedule.Days) == 0 {
		output := fmt.Sprintf("TVmaze | %s | Status: %s",
			results.Name,
			results.Status,
		)
		return output, nil
	}

	output := fmt.Sprintf("TVmaze | %s | Airtime: %s %s on %s | Status: %s",
		results.Name,
		results.Schedule.Days[0],
		results.Schedule.Time,
		results.Network.Name,
		results.Status,
	)
	return output, nil
}

func init() {
	bot.RegisterCommand(
		"^tv (.+)$",
		tvmaze)
}
