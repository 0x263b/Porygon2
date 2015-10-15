package tvmaze

import (
	"fmt"
	"github.com/0x263b/Porygon2"
	"github.com/0x263b/Porygon2/web"
	"net/url"
)

const (
	tvMazeURL = "http://api.tvmaze.com/singlesearch/shows?q=%s"
)

type Showinfo struct {
	ID        int      `json:"id"`
	URL       string   `json:"url"`
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Language  string   `json:"language"`
	Genres    []string `json:"genres"`
	Status    string   `json:"status"`
	Runtime   int      `json:"runtime"`
	Premiered string   `json:"premiered"`
	Schedule  struct {
		Time string   `json:"time"`
		Days []string `json:"days"`
	} `json:"schedule"`
	Rating struct {
		Average float64 `json:"average"`
	} `json:"rating"`
	Weight  int `json:"weight"`
	Network struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Country struct {
			Name     string `json:"name"`
			Code     string `json:"code"`
			Timezone string `json:"timezone"`
		} `json:"country"`
	} `json:"network"`
	Webchannel interface{} `json:"webChannel"`
	Externals  struct {
		Tvrage  int `json:"tvrage"`
		Thetvdb int `json:"thetvdb"`
	} `json:"externals"`
	Image struct {
		Medium   string `json:"medium"`
		Original string `json:"original"`
	} `json:"image"`
	Summary string `json:"summary"`
	Updated int    `json:"updated"`
	Links   struct {
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
	ID       int    `json:"id"`
	URL      string `json:"url"`
	Name     string `json:"name"`
	Season   int    `json:"season"`
	Number   int    `json:"number"`
	Airdate  string `json:"airdate"`
	Airtime  string `json:"airtime"`
	Airstamp string `json:"airstamp"`
	Runtime  int    `json:"runtime"`
	Image    struct {
		Medium   string `json:"medium"`
		Original string `json:"original"`
	} `json:"image"`
	Summary string `json:"summary"`
	Links   struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
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
