package bot

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"
)

// Cmd holds the parsed user's input for easier handling of commands
type Cmd struct {
	Raw     string   // Raw is full string passed to the command
	Channel string   // Channel where the command was called
	Nick    string   // User who sent the message
	Message string   // Full string without the prefix
	Command string   // Command is the first argument passed to the bot
	FullArg string   // Full argument as a single string
	Args    []string // Arguments as array
}

// PassiveCmd holds the information which will be passed to passive commands when receiving a message on the channel
type PassiveCmd struct {
	Raw     string // Raw message sent to the channel
	Channel string // Channel which the message was sent to
	Nick    string // Nick of the user which sent the message
}

type customCommand struct {
	Version int
	Cmd     *regexp.Regexp
	CmdFunc activeCmdFunc
}

type incomingMessage struct {
	Channel        string
	Text           string
	SenderNick     string
	BotCurrentNick string
}

const (
	commandNotAvailable   = "Command %v not available."
	noCommandsAvailable   = "No commands available."
	errorExecutingCommand = "Error executing %s: %s"
)

type passiveCmdFunc func(cmd *PassiveCmd) (string, error)
type activeCmdFunc func(cmd *Cmd, matches []string) (string, error)

var (
	commands        = make(map[string]*customCommand)
	passiveCommands = make(map[string]passiveCmdFunc)
)

func RegisterCommand(command string, cmdFunc activeCmdFunc) {
	commands[command] = &customCommand{
		Cmd:     regexp.MustCompile(command),
		CmdFunc: cmdFunc,
	}
}

func RegisterPassiveCommand(command string, cmdFunc func(cmd *PassiveCmd) (string, error)) {
	passiveCommands[command] = cmdFunc
}

func IsPrivateMsg(channel, currentNick string) bool {
	return channel == currentNick
}

func IsIgnored(senderNick string) bool {
	state := GetUserKey(senderNick, "ignore")
	if state == "true" {
		return true
	} else {
		return false
	}
}

func IsAdmin(senderNick string) bool {
	state := GetUserKey(senderNick, "admin")
	if state == "true" {
		return true
	} else {
		return false
	}
}

func messageReceived(channel, text, senderNick string, conn ircConnection) {
	if IsPrivateMsg(channel, conn.GetNick()) {
		channel = senderNick // should reply in private
	}

	if IsIgnored(senderNick) {
		return
	}

	command := parse(text, channel, senderNick)
	if command == nil {
		executePassiveCommands(&PassiveCmd{
			Raw:     text,
			Channel: channel,
			Nick:    senderNick,
		}, conn)
		return
	}

	switch command.Command {
	case joinCommand:
		join(command, channel, senderNick, conn)
	case partCommand:
		part(command, channel, senderNick, conn)
	default:
		handleCmd(command, conn)
	}

}

func executePassiveCommands(cmd *PassiveCmd, conn ircConnection) {
	var wg sync.WaitGroup

	for k, v := range passiveCommands {
		cmdName := k
		cmdFunc := v

		wg.Add(1)

		go func() {
			defer wg.Done()

			log.Println("Executing passive command: ", cmdName)
			result, err := cmdFunc(cmd)
			if err != nil {
				log.Println(err)
			} else if result != "" {
				conn.Privmsg(cmd.Channel, result)
			}
		}()
	}

	wg.Wait()
}

func handleCmd(c *Cmd, conn ircConnection) {
	for _, k := range commands {
		if matches := k.Cmd.FindStringSubmatch(c.Message); len(matches) > 0 {
			message, err := k.CmdFunc(c, matches)
			checkCmdError(err, c, conn)

			msg_arr := strings.Split(message, "\n")
			for i, v := range msg_arr {
				if v != "" {
					// Flood Protection
					if i > 0 {
						time.Sleep(time.Millisecond * 800)
					}
					conn.Privmsg(c.Channel, v)
				}
			}
		}
	}

	// log.Printf("HandleCmd %v %v", c.Command, c.FullArg)

	return
}

func checkCmdError(err error, c *Cmd, conn ircConnection) {
	if err != nil {
		errorMsg := fmt.Sprintf(errorExecutingCommand, c.Command, err.Error())
		log.Printf(errorMsg)
		conn.Privmsg(c.Channel, errorMsg)
	}
}
