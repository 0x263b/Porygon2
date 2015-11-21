package weather

import (
	"fmt"
	"github.com/0x263b/Porygon2"
	"github.com/0x263b/Porygon2/web"
	"math"
	"time"
)

func forecast(command *bot.Cmd, matches []string) (msg string, err error) {

	var location string = matches[1]
	var coords string

	if location == "" {
		coords = checkLocation(command.Nick)
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
	err = web.GetJSON(fmt.Sprintf(DarkSkyURL, bot.Config.API.Weather, coords), data)
	if err != nil {
		return fmt.Sprintf("Could not get weather for: %s", location), nil
	}

	units := "°C"
	if data.Flags.Units == "us" {
		units = "°F"
	}

	output := fmt.Sprintf("Forecast | %s ", location)

	for i := range data.Daily.Data[0:4] {
		tm := time.Unix(data.Daily.Data[i].Time, 0)
		day := tm.Weekday()
		output += fmt.Sprintf("| %s: %s %v%s/%v%s ",
			day,
			Emoji(data.Daily.Data[i].Icon),
			math.Ceil(data.Daily.Data[i].TemperatureMax),
			units,
			math.Ceil(data.Daily.Data[i].TemperatureMin),
			units,
		)
	}

	return output, nil
}

func init() {
	bot.RegisterCommand(
		"^f(?:o(?:recast)?)?(?: (.+))?$",
		forecast)
}
