// Package bot provides a simple to use IRC bot
package bot

import (
	"crypto/tls"
	"fmt"
	"github.com/steveyen/gkvlite"
	"github.com/thoj/go-ircevent"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// Configure must contain the necessary data to connect to an IRC server
type Configure struct {
	Server        string // IRC server:port. Ex: irc.rizon.net:6697
	Channel       string // Initial channel to connect. Ex: "#channel"
	User          string // The IRC username the bot will use
	Nick          string // The nick the bot will use
	Nickserv      string // Nickserv password
	Password      string // Server password
	Modes         string // User modes. Ex: GRp
	UseTLS        bool   // Should connect using TLS? (yes)
	TLSServerName string // Must supply if UseTLS is true
	Debug         bool   // This will log all IRC communication to standad output
	Prefix        string // Prefix used to identify a command. !hello whould be identified as a command
	Owner         string // Owner of the bot. Used for admin-only commands
	API
}

type API struct {
	Lastfm                string
	Giphy                 string
	Translate             string
	TwitterConsumerKey    string
	TwitterConsumerSecret string
	Weather               string
	Wolfram               string
	Youtube               string
	Geocode               string
	Bing                  string
}

type ircConnection interface {
	Privmsg(target, message string)
	GetNick() string
	Join(target string)
	Part(target string)
}

var (
	Conn         *irc.Connection
	Config       *Configure
	ChannelNicks = make(map[string][]string)
)

func logChannel(channel, text, senderNick string, action bool) {
	t := time.Now()
	channel = strings.Replace(channel, "/", "–", -1)

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	if _, err := os.Stat(fmt.Sprintf("%s/logs", dir)); os.IsNotExist(err) {
		os.MkdirAll(fmt.Sprintf("%s/logs", dir), 0711)
	}
	mo := fmt.Sprintf("%v", int(t.Month()))
	if len(mo) < 2 {
		mo = fmt.Sprintf("0%s", mo)
	}
	f, err := os.OpenFile(fmt.Sprintf("%s/logs/%s.%v.%v.log", dir, channel, t.Year(), mo), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetFlags(0)
	log.SetOutput(f)

	var line string

	if action {
		line = fmt.Sprintf("[%s] * %s %s", t.Format(time.RFC3339), senderNick, text)
	} else {
		line = fmt.Sprintf("[%s] <%s> %s", t.Format(time.RFC3339), senderNick, text)
	}

	log.Println(line)
	return
}

func onPRIVMSG(e *irc.Event) {
	if e.Arguments[0] != Conn.GetNick() {
		channel := strings.Replace(e.Arguments[0], "#", "", -1)
		logChannel(channel, e.Message(), e.Nick, false)
	}
	messageReceived(e.Arguments[0], e.Message(), e.Nick, Conn)
}

func onACTION(e *irc.Event) {
	if e.Arguments[0] != Conn.GetNick() {
		channel := strings.Replace(e.Arguments[0], "#", "", -1)
		logChannel(channel, e.Message(), e.Nick, true)
	}
	// messageReceived(e.Arguments[0], e.Message(), e.Nick, Conn)
}

func getServerName() string {
	separatorIndex := strings.LastIndex(Config.Server, ":")
	if separatorIndex != -1 {
		return Config.Server[:separatorIndex]
	} else {
		return Config.Server
	}
}

func connect() {
	Conn = irc.IRC(Config.User, Config.Nick)
	Conn.Password = Config.Password
	Conn.UseTLS = Config.UseTLS
	Conn.TLSConfig = &tls.Config{
		ServerName:         getServerName(),
		InsecureSkipVerify: true,
	}
	Conn.Version = "Porygon2 → https://github.com/0x263b/Porygon2"
	Conn.VerboseCallbackHandler = Config.Debug
	err := Conn.Connect(Config.Server)
	if err != nil {
		log.Fatal(err)
	}
}

func onEndOfMotd(e *irc.Event) {
	SetUserKey(Config.Owner, "admin", "true")
	Conn.Privmsg("nickserv", "identify "+Config.Nickserv)
	Conn.Mode(Config.Nick, Config.Modes)
	SetChannelKey(Config.Channel, "auto_join", true)
	Channels.VisitItemsAscend([]byte(""), true, func(i *gkvlite.Item) bool {
		if GetChannelKey(string(i.Key), "auto_join") {
			time.Sleep(2 * time.Second)
			Conn.Join(string(i.Key))
		}
		return true
	})
}

func GetNames(channel string) []string {
	Conn.SendRawf("NAMES %v", channel)
	return ChannelNicks[channel]
}

func onNames(e *irc.Event) {
	// Strip modes
	r, _ := regexp.Compile("([~&@%+])")
	s := r.ReplaceAllString(e.Arguments[3], "")

	// Combine all responses & remove duplicates
	old := ChannelNicks[strings.ToLower(e.Arguments[2])]
	nu := strings.Split(s, " ")
	uniq := removeDuplicates(append(old, nu...))

	ChannelNicks[strings.ToLower(e.Arguments[2])] = uniq
}

func onKick(e *irc.Event) {
	if e.Arguments[1] == Config.Nick {
		time.Sleep(2 * time.Second)
		Conn.Join(e.Arguments[0])
	}
}

func ConfigureEvents() {
	Conn.AddCallback("376", onEndOfMotd)
	Conn.AddCallback("353", onNames)
	Conn.AddCallback("KICK", onKick)
	Conn.AddCallback("PRIVMSG", onPRIVMSG)
	Conn.AddCallback("CTCP_ACTION", onACTION)
}

// Run reads the Config, connect to the specified IRC server and starts the bot.
// The bot will automatically join all the channels specified in the Configuration
func Run(c *Configure) {
	initkv()
	Config = c
	connect()
	ConfigureEvents()
	Conn.Loop()
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
