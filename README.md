# Porygon2

An IRC Bot written in Go

### Install
```
# Dependencies
go get -u github.com/thoj/go-ircevent github.com/steveyen/gkvlite github.com/PuerkitoBio/goquery github.com/dustin/go-humanize github.com/kennygrant/sanitize gopkg.in/xmlpath.v2 github.com/kurrik/oauth1a github.com/kurrik/twittergo

# Porygon2
go get -u github.com/0x263b/porygon2
```

### Example
```Go
package main

import (
	"github.com/0x263b/porygon2"
	_ "github.com/0x263b/porygon2/commands/8ball"
	_ "github.com/0x263b/porygon2/commands/admin"
	_ "github.com/0x263b/porygon2/commands/bing"
	_ "github.com/0x263b/porygon2/commands/choose"
	_ "github.com/0x263b/porygon2/commands/lastfm"
	_ "github.com/0x263b/porygon2/commands/opengraph"
	_ "github.com/0x263b/porygon2/commands/translate"
	_ "github.com/0x263b/porygon2/commands/twitter"
	_ "github.com/0x263b/porygon2/commands/urbandictionary"
	_ "github.com/0x263b/porygon2/commands/weather"
	_ "github.com/0x263b/porygon2/commands/wolfram"
	_ "github.com/0x263b/porygon2/commands/youtube"
)

func main() {
	config := newConfig()
	bot.Run(config)
}

func newConfig() *bot.Configure {
	return &bot.Configure{
		Server:   "irc.rizon.net:6697", // "server:port"
		Channel:  "#Porygon2",          // "#channel" or "#channel key"
		User:     "Porygon2",           // "bot"
		Nick:     "Porygon2",           // "bot"
		Nickserv: "some password",      // leave as "" if none
		Modes:    "GRp",                // "GRp"
		UseTLS:   true,                 // true/false
		Debug:    false,                // true/false
		Prefix:   "!",                  // "!"
		Owner:    "joe",                // your nick
		API: bot.API{
			Bing:                  "",
			Geocode:               "",
			Giphy:                 "",
			Lastfm:                "",
			Translate:             "",
			TwitterConsumerKey:    "",
			TwitterConsumerSecret: "",
			Weather:               "",
			Wolfram:               "",
			Youtube:               "",
		},
	}
}
```

#### Ubuntu service

Save as `/etc/init/porygon2.conf` and run `service porygon2 start`

```
# Upstart Configuration

description     "Porygon2"
author          "Black Smiling Face"

start on (net-device-up
          and local-filesystems
          and runlevel [2345])

stop on runlevel [016]

respawn

exec /path/to/porygon2
```

#### APIs

* [Bing](https://www.microsoft.com/cognitive-services/en-us/bing-web-search-api): [commands/bing](commands/bing)
	*  Sign up for the free tier then get your [API Key](https://www.microsoft.com/cognitive-services/en-US/subscriptions)
* [Giphy](https://github.com/Giphy/GiphyAPI): [commands/giphy](commands/giphy/giphy.go)
	* The public beta key is good enough
* [Google Geocode](https://developers.google.com/maps/documentation/geocoding/intro): [commands/weather](commands/weather/weather.go)
	* Key can be found on the [Google API Console](https://console.developers.google.com/apis/credentials)
* [Last.fm](http://www.last.fm/api): [commands/lastfm](commands/lastfm/)
* [Forecast.io](https://developer.forecast.io/docs/v2): [commands/weather](commands/weather/weather.go)
	* [Create and view your API key here](https://developer.forecast.io/)
* [Microsoft Translator](https://msdn.microsoft.com/en-us/library/hh454949.aspx): [commands/translate](commands/translate/translate.go)
	* Go to the [Azure Marketplace](https://datamarket.azure.com/dataset/bing/microsofttranslator) and subscribe to Microsoft Translator, then test it with the [Service Explorer](https://datamarket.azure.com/dataset/explore/bing/microsofttranslator). Top of the page, you'll see the `Primary Account Key`
* [Twitter](https://dev.twitter.com/rest/public): [commands/twitter](commands/twitter/twitter.go)
	* [Create an app](https://apps.twitter.com/) (Read only). Under `Keys and Access`, copy the `Consumer Key` and `Consumer Secret`
* [Wolfram Alpha](http://products.wolframalpha.com/api/): [commands/wolfram](commands/wolfram/wolfram.go)
	* Register, then [create and App](https://developer.wolframalpha.com/portal/myapps/index.html), and copy the `AppID`
* [Youtube](https://developers.google.com/youtube/v3/): [commands/youtube](commands/youtube/youtube.go)
	* Key can be found on the [Google API Console](https://console.developers.google.com/apis/credentials)

