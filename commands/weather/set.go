package weather

import (
	"fmt"
	"github.com/0x263b/porygon2"
	"github.com/0x263b/porygon2/web"
	"net/url"
	"strings"
)

func checkLocation(senderNick string) (location string, coords string) {
	location = bot.GetUserKey(senderNick, "location")
	coords = bot.GetUserKey(senderNick, "coords")
	return string(location), string(coords)
}

func setLocation(command *bot.Cmd, matches []string) (msg string, err error) {
	geo := &Geocode{}
	err = web.GetJSON(fmt.Sprintf(GeocodeURL, url.QueryEscape(strings.TrimSpace(matches[1])), bot.Config.API.Geocode), geo)
	if err != nil {
		return fmt.Sprintf("Could not find %s", strings.TrimSpace(matches[1])), nil
	}
	if geo.Status != "OK" {
		return fmt.Sprintf("Could not find %s", strings.TrimSpace(matches[1])), nil
	}

	coords := fmt.Sprintf("%v,%v", geo.Results[0].Geometry.Location.Lat, geo.Results[0].Geometry.Location.Lng)

	location := geo.Results[0].AddressComponents[0].ShortName
	if len(geo.Results[0].AddressComponents) > 1 {
		if geo.Results[0].AddressComponents[1].Types[0] == "locality" {
			location = geo.Results[0].AddressComponents[1].ShortName
		}
	}

	bot.SetUserKey(command.Nick, "location", location)
	bot.SetUserKey(command.Nick, "coords", coords)
	return fmt.Sprintf("%s: Location updated to: %s", command.Nick, location), nil
}

func init() {
	bot.RegisterCommand(
		"^set location (.+)$",
		setLocation)
}
