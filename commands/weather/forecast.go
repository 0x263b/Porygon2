package weather

import (
	"fmt"
	"github.com/0x263b/Porygon2"
	"github.com/0x263b/Porygon2/web"
	"net/url"
)

func forecast(command *bot.Cmd, matches []string) (msg string, err error) {

	location := matches[1]

	if location == "" {
		location = checkLocation(command.Nick)
	}

	if location == "" {
		return "Location not provided, nor on file. Use `-set location <location>` to save", nil
	}

	query := fmt.Sprintf("select * from weather.forecast where woeid in (select woeid from geo.places(1) where text=\"%s\") and u=\"c\"", location)
	data := &yahooWeather{}
	err = web.GetJSON(fmt.Sprintf(yahooURL, url.QueryEscape(query), yahooAppID), data)
	if err != nil {
		return fmt.Sprintf("Could not get weather for: %s", location), nil
	}
	if data.Query.Results.Channel.Location.City == "" {
		return fmt.Sprintf("Could not get weather for: %s", location), nil
	}

	output := fmt.Sprintf("Forecast | %s ", data.Query.Results.Channel.Location.City)
	for i := range data.Query.Results.Channel.Item.Forecast {
		output += fmt.Sprintf("| %s: %s %s°%s/%s°%s ",
			data.Query.Results.Channel.Item.Forecast[i].Day,
			data.Query.Results.Channel.Item.Forecast[i].Text,
			data.Query.Results.Channel.Item.Forecast[i].High,
			data.Query.Results.Channel.Units.Temperature,
			data.Query.Results.Channel.Item.Forecast[i].Low,
			data.Query.Results.Channel.Units.Temperature,
		)
	}

	return output, nil
}

func init() {
	bot.RegisterCommand(
		"^f(?:o(?:recast)?)?(?: (.+))?$",
		forecast)
}
