// A Discord and YouTube bot written in Go
package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// Version of the bot
const Version = "v0.0.1-alpha"

// Authentication Token to be read from environment or command line
var Token = ""

func init() {
	Token = os.Getenv("BOT_TOKEN")
	flag.StringVar(&Token, "t", "", "Discord Bot Authentication Token")
}

func main() {
	// Make sure we have an Authentication Token
	flag.Parse()
	if Token == "" {
		fmt.Printf("missing authentication token!\n")
		return
	}

	// Use Token to authenticate and verify successful authentication
	var session, err = discordgo.New(Token)
	if err != nil {
		fmt.Printf("invalid authentication token: %s\n", err)
		return
	}
	session.State.User, err = session.User("@me")
	if err != nil {
		fmt.Printf("error fetching user information: %s\n", err)
		return
	}
	fmt.Printf("emonk-%s authenticated as %s\n", Version, session.State.User)

	// Register a callback for MessageCreate events
	session.AddHandler(messageCreate)

	// Open a websocket connection
	err = session.Open()
	if err != nil {
		fmt.Println("error opening connection: %s", err)
		return
	}

	// UpdateStatus() requires a websocket, i. e. needs to be called after Open()
	err = session.UpdateStatus(0, "Chasing Nephthys")
	if err != nil {
		fmt.Printf("error setting current game: %s\n", err)
		return
	}

	// Setup a channel to wait for a signal to terminate me
	fmt.Printf("Waiting for signal, e. g. CTRL-C, to terminate me ...\n")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc // now waiting

	session.Close()
}

// Callback for MessageCreate events used to reply/react to certain messages
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// never reply/react to myself
	if m.Author.ID == s.State.User.ID {
		return
	}

	var msg string = strings.TrimSpace(m.Content)
	var reply string = ""
	var reaction string = ""

	switch msg {
	case "ping":
		reply = fmt.Sprintf("<@%s> Pong! :ping_pong:", m.Author.ID)
		reaction = "🔁" // unicode for :repeat: (arrows_clockwise)
	case "pong":
		reply = fmt.Sprintf("<@%s> Ping! :ping_pong:", m.Author.ID)
		reaction = "🔄" // unicode for :arrows_counterclockwise:
	case "hi", "hiho", "hello":
		reply = fmt.Sprintf("<@%s> Hello! :wave:", m.Author.ID)
		reaction = "👋" // unicode for :wave:
	case "no comment":
		reply = "`Real programmers don't write comments.\nIf it was hard to write, it should be hard to read.`"
	}
	if len(reply) > 0 {
		s.ChannelMessageSend(m.ChannelID, reply)
	}
	if len(reaction) > 0 {
		s.MessageReactionAdd(m.ChannelID, m.ID, reaction)
	}
}
