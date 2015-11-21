package weather

type Geocode struct {
	Results []struct {
		AddressComponents []struct {
			LongName  string   `json:"long_name"`
			ShortName string   `json:"short_name"`
			Types     []string `json:"types"`
		} `json:"address_components"`
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Bounds struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"bounds"`
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			LocationType string `json:"location_type"`
			Viewport     struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"viewport"`
		} `json:"geometry"`
		PlaceID string   `json:"place_id"`
		Types   []string `json:"types"`
	} `json:"results"`
	Status string `json:"status"`
}

type Flags struct {
	DarkSkyUnavailable string   `json:"darksky-unavailable"`
	DarkSkyStations    []string `json:"darksky-stations"`
	DataPointStations  []string `json:"datapoint-stations"`
	ISDStations        []string `json:"isds-stations"`
	LAMPStations       []string `json:"lamp-stations"`
	METARStations      []string `json:"metars-stations"`
	METNOLicense       string   `json:"metnol-license"`
	Sources            []string `json:"sources"`
	Units              string   `json:"units"`
}

type DataPoint struct {
	Time                   int64   `json:"time"`
	Summary                string  `json:"summary"`
	Icon                   string  `json:"icon"`
	SunriseTime            float64 `json:"sunriseTime"`
	SunsetTime             float64 `json:"sunsetTime"`
	PrecipIntensity        float64 `json:"precipIntensity"`
	PrecipIntensityMax     float64 `json:"precipIntensityMax"`
	PrecipIntensityMaxTime float64 `json:"precipIntensityMaxTime"`
	PrecipProbability      float64 `json:"precipProbability"`
	PrecipType             string  `json:"precipType"`
	PrecipAccumulation     float64 `json:"precipAccumulation"`
	Temperature            float64 `json:"temperature"`
	TemperatureMin         float64 `json:"temperatureMin"`
	TemperatureMinTime     float64 `json:"temperatureMinTime"`
	TemperatureMax         float64 `json:"temperatureMax"`
	TemperatureMaxTime     float64 `json:"temperatureMaxTime"`
	ApparentTemperature    float64 `json:"apparentTemperature"`
	DewPoint               float64 `json:"dewPoint"`
	WindSpeed              float64 `json:"windSpeed"`
	WindBearing            float64 `json:"windBearing"`
	CloudCover             float64 `json:"cloudCover"`
	Humidity               float64 `json:"humidity"`
	Pressure               float64 `json:"pressure"`
	Visibility             float64 `json:"visibility"`
	Ozone                  float64 `json:"ozone"`
	MoonPhase              float64 `json:"moonPhase"`
}

type DataBlock struct {
	Summary string      `json:"summary"`
	Icon    string      `json:"icon"`
	Data    []DataPoint `json:"data"`
}

type alert struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Time        float64 `json:"time"`
	Expires     float64 `json:"expires"`
	URI         string  `json:"uri"`
}

type Forecast struct {
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timezone  string    `json:"timezone"`
	Offset    float64   `json:"offset"`
	Currently DataPoint `json:"currently"`
	Minutely  DataBlock `json:"minutely"`
	Hourly    DataBlock `json:"hourly"`
	Daily     DataBlock `json:"daily"`
	Alerts    []alert   `json:"alerts"`
	Flags     Flags     `json:"flags"`
	APICalls  int       `json:"apicalls"`
	Code      int       `json:"code"`
}

type Units string

const (
	CA   Units = "ca"
	SI   Units = "si"
	US   Units = "us"
	UK   Units = "uk"
	AUTO Units = "auto"
)
