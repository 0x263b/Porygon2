package lastfm

type NowPlaying struct {
	Recenttracks struct {
		Track []struct {
			Artist struct {
				Text string `json:"#text"`
				Mbid string `json:"mbid"`
			} `json:"artist"`
			Name  string `json:"name"`
			Album struct {
				Text string `json:"#text"`
				Mbid string `json:"mbid"`
			} `json:"album"`
			Attr struct {
				Nowplaying string `json:"nowplaying"`
			} `json:"@attr"`
		} `json:"track"`
		Attr struct {
			Total string `json:"total"`
		} `json:"@attr"`
	} `json:"recenttracks"`
	Error int `json:"error"`
}

type ArtistTags struct {
	Toptags struct {
		Tag []struct {
			Name string `json:"name"`
		} `json:"tag"`
	} `json:"toptags"`
}

type WeeklyCharts struct {
	Topartists struct {
		Artist []struct {
			Name      string `json:"name"`
			Playcount string `json:"playcount"`
			Attr      struct {
				Rank string `json:"rank"`
			} `json:"@attr"`
		} `json:"artist"`
		Attr struct {
			Total string `json:"total"`
		} `json:"@attr"`
	} `json:"topartists"`
	Error int `json:"error"`
}
