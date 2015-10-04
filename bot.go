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

// Config must contain the necessary data to connect to an IRC server
type Config struct {
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
}

type ircConnection interface {
	Privmsg(target, message string)
	GetNick() string
	Join(target string)
	Part(target string)
}

var (
	irccon       *irc.Connection
	config       *Config
	ChannelNicks = make(map[string][]string)
)

func logChannel(channel, text, senderNick string) {
	t := time.Now()

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
	log.SetPrefix(t.Format(time.RFC3339))
	log.SetOutput(f)

	line := fmt.Sprintf("\t%s\t%s", senderNick, text)
	log.Println(line)
	return
}

func onPRIVMSG(e *irc.Event) {
	if e.Arguments[0] != irccon.GetNick() {
		channel := strings.Replace(e.Arguments[0], "#", "", -1)
		logChannel(channel, e.Message(), e.Nick)
	}
	messageReceived(e.Arguments[0], e.Message(), e.Nick, irccon)
}

func getServerName() string {
	separatorIndex := strings.LastIndex(config.Server, ":")
	if separatorIndex != -1 {
		return config.Server[:separatorIndex]
	} else {
		return config.Server
	}
}

func connect() {
	irccon = irc.IRC(config.User, config.Nick)
	irccon.Password = config.Password
	irccon.UseTLS = config.UseTLS
	irccon.TLSConfig = &tls.Config{
		ServerName:         getServerName(),
		InsecureSkipVerify: true,
	}
	irccon.VerboseCallbackHandler = config.Debug
	err := irccon.Connect(config.Server)
	if err != nil {
		log.Fatal(err)
	}
}

func onEndOfMotd(e *irc.Event) {
	SetUserKey(config.Owner, "admin", "true")
	irccon.Privmsg("nickserv", "identify "+config.Nickserv)
	irccon.Mode(config.Nick, config.Modes)
	SetChannelKey(config.Channel, "auto_join", true)
	Channels.VisitItemsAscend([]byte(""), true, func(i *gkvlite.Item) bool {
		if GetChannelKey(string(i.Key), "auto_join") {
			time.Sleep(2 * time.Second)
			irccon.Join(string(i.Key))
		}
		return true
	})
}

func GetNames(channel string) []string {
	irccon.SendRaw(fmt.Sprintf("NAMES %v", channel))
	time.Sleep(1 * time.Second)
	return ChannelNicks[channel]
}

func onNames(e *irc.Event) {
	// Strip modes
	r, _ := regexp.Compile("([~&@%+])")
	s := r.ReplaceAllString(e.Arguments[3], "")

	// Combine all responses & remove duplicates
	old := ChannelNicks[strings.ToLower(e.Arguments[2])]
	nu := strings.Split(s, "\\s")
	uniq := removeDuplicates(append(old, nu...))

	ChannelNicks[strings.ToLower(e.Arguments[2])] = uniq
	log.Printf("Names: %v", uniq)
}

func onEndOfNames(e *irc.Event) {
	log.Printf("onEndOfNames: %v", e.Arguments)
}

func onKick(e *irc.Event) {
	if e.Arguments[1] == config.Nick {
		time.Sleep(2 * time.Second)
		irccon.Join(e.Arguments[0])
	}
}

func configureEvents() {
	irccon.AddCallback("376", onEndOfMotd)
	irccon.AddCallback("366", onEndOfNames)
	irccon.AddCallback("353", onNames)
	irccon.AddCallback("KICK", onKick)
	irccon.AddCallback("PRIVMSG", onPRIVMSG)
	irccon.AddCallback("CTCP_ACTION", onPRIVMSG)
}

// Run reads the Config, connect to the specified IRC server and starts the bot.
// The bot will automatically join all the channels specified in the configuration
func Run(c *Config) {
	initkv()
	config = c
	connect()
	configureEvents()
	irccon.Loop()
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
