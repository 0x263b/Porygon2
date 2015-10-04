package weather

import (
	"fmt"
	"github.com/0x263b/Porygon2"
	"github.com/0x263b/Porygon2/web"
	"net/url"
)

const (
	yahooURL   = "https://query.yahooapis.com/v1/public/yql?format=json&q=%s&appid=%s&env=store://datatables.org/alltableswithkeys"
	yahooAppID = ""
)

func weather(command *bot.Cmd, matches []string) (msg string, err error) {

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

	output := fmt.Sprintf(
		"Weather | %s | %s %s°%s. Wind chill: %s°%s. Humidity: %s%%",
		data.Query.Results.Channel.Location.City,
		data.Query.Results.Channel.Item.Condition.Text,
		data.Query.Results.Channel.Item.Condition.Temp,
		data.Query.Results.Channel.Units.Temperature,
		data.Query.Results.Channel.Wind.Chill,
		data.Query.Results.Channel.Units.Temperature,
		data.Query.Results.Channel.Atmosphere.Humidity)

	return output, nil
}

func init() {
	bot.RegisterCommand(
		"^w(?:e(?:ather)?)?(?: (.+))?$",
		weather)
}
