package bot

import (
	"strings"
)

const (
	joinCommand = "join"
	partCommand = "part"
)

func join(c *Cmd, channel, senderNick string, conn ircConnection) {
	if !IsAdmin(senderNick) {
		return
	}

	channelToJoin := strings.TrimSpace(c.FullArg)
	channelToJoin = strings.ToLower(channelToJoin)

	if channelToJoin == "" {
		return
	}

	SetChannelKey(channelToJoin, "auto_join", true)
	conn.Join(channelToJoin)
}

func part(c *Cmd, channel, senderNick string, conn ircConnection) {
	if !IsAdmin(senderNick) {
		return
	}

	channelToPart := strings.TrimSpace(c.FullArg)
	channelToPart = strings.ToLower(channelToPart)

	if channelToPart == "" {
		channelToPart = channel
	}

	SetChannelKey(channelToPart, "auto_join", false)
	conn.Part(channelToPart)
}
