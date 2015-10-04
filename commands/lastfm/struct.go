package lastfm

type NowPlaying struct {
	Recenttracks struct {
		Track []struct {
			Artist struct {
				Text string `json:"#text"`
				Mbid string `json:"mbid"`
			} `json:"artist"`
			Name       string `json:"name"`
			Streamable string `json:"streamable"`
			Mbid       string `json:"mbid"`
			Album      struct {
				Text string `json:"#text"`
				Mbid string `json:"mbid"`
			} `json:"album"`
			URL   string `json:"url"`
			Image []struct {
				Text string `json:"#text"`
				Size string `json:"size"`
			} `json:"image"`
			Attr struct {
				Nowplaying string `json:"nowplaying"`
			} `json:"@attr"`
		} `json:"track"`
		Attr struct {
			User       string `json:"user"`
			Page       string `json:"page"`
			Perpage    string `json:"perPage"`
			Totalpages string `json:"totalPages"`
			Total      string `json:"total"`
		} `json:"@attr"`
	} `json:"recenttracks"`
	Error   int           `json:"error"`
	Message string        `json:"message"`
	Links   []interface{} `json:"links"`
}

type ArtistTags struct {
	Toptags struct {
		Tag []struct {
			Count int    `json:"count"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"tag"`
		Attr struct {
			Artist string `json:"artist"`
		} `json:"@attr"`
	} `json:"toptags"`
}

type WeeklyCharts struct {
	Topartists struct {
		Artist []struct {
			Name       string `json:"name"`
			Playcount  string `json:"playcount"`
			Mbid       string `json:"mbid"`
			URL        string `json:"url"`
			Streamable string `json:"streamable"`
			Image      []struct {
				Text string `json:"#text"`
				Size string `json:"size"`
			} `json:"image"`
			Attr struct {
				Rank string `json:"rank"`
			} `json:"@attr"`
		} `json:"artist"`
		Attr struct {
			User       string `json:"user"`
			Page       string `json:"page"`
			Perpage    string `json:"perPage"`
			Totalpages string `json:"totalPages"`
			Total      string `json:"total"`
		} `json:"@attr"`
	} `json:"topartists"`
	Error   int           `json:"error"`
	Message string        `json:"message"`
	Links   []interface{} `json:"links"`
}
