package weather

import (
	"fmt"
	"github.com/0x263b/porygon2"
	"github.com/0x263b/porygon2/web"
	"math"
	"net/url"
)

const (
	GeocodeURL = "https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s"
	DarkSkyURL = "https://api.forecast.io/forecast/%s/%s?units=auto&exclude=minutely,hourly,alerts"
)

func Emoji(icon string) string {
	if icon == "clear-day" {
		return "☀️"
	} else if icon == "clear-night" {
		return "🌙"
	} else if icon == "rain" {
		return "☔️"
	} else if icon == "snow" {
		return "❄️"
	} else if icon == "sleet" {
		return "☔️❄️"
	} else if icon == "wind" {
		return "💨"
	} else if icon == "fog" {
		return "🌁"
	} else if icon == "cloudy" {
		return "☁️"
	} else if icon == "partly-cloudy-day" {
		return "⛅"
	} else if icon == "partly-cloudy-night" {
		return "⛅"
	} else {
		return ""
	}
}

func Round(f float64) float64 {
	return math.Floor(f + .5)
}

func getCoords(location string) string {
	var err error
	geo := &Geocode{}
	err = web.GetJSON(fmt.Sprintf(GeocodeURL, url.QueryEscape(location), bot.Config.Geocode), geo)
	if err != nil || geo.Status != "OK" {
		return ""
	}
	return fmt.Sprintf("%v,%v", geo.Results[0].Geometry.Location.Lat, geo.Results[0].Geometry.Location.Lng)
}

func weather(command *bot.Cmd, matches []string) (msg string, err error) {

	var location string = matches[1]
	var coords string

	if location == "" {
		location, coords = checkLocation(command.Nick)
		if coords == "" {
			return "Location not provided, nor on file. Use `-set location <location>` to save", nil
		}
	} else {
		coords = getCoords(location)
		if coords == "" {
			return fmt.Sprintf("Could not find %s", location), nil
		}
	}

	data := &Forecast{}
	err = web.GetJSON(fmt.Sprintf(DarkSkyURL, bot.Config.Weather, coords), data)
	if err != nil {
		return fmt.Sprintf("Could not get weather for: %s", location), nil
	}

	units := "°C"
	windspeed := "m/s"
	if data.Flags.Units == "us" {
		units = "°F"
		windspeed = "mph"
	} else if data.Flags.Units == "ca" {
		windspeed = "km/h"
	} else if data.Flags.Units == "uk2" {
		windspeed = "mph"
	}

	return fmt.Sprintf("Weather | %s | Now: %s %s %v%s. Today: %s %v%s/%v%s. Humidity: %v%%. Wind: %v%s. Precipitation: %v%%.",
		location,
		data.Currently.Summary,
		Emoji(data.Currently.Icon),
		Round(data.Currently.Temperature),
		units,
		Emoji(data.Daily.Data[0].Icon),
		Round(data.Daily.Data[0].TemperatureMax),
		units,
		Round(data.Daily.Data[0].TemperatureMin),
		units,
		int(data.Daily.Data[0].Humidity*100),
		data.Daily.Data[0].WindSpeed,
		windspeed,
		int(data.Daily.Data[0].PrecipProbability*100)), nil
}

func init() {
	bot.RegisterCommand(
		"^w(?:e(?:ather)?)?(?: (.+))?$",
		weather)
}
