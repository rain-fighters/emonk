// A Discord and YouTube bot written in go
package main

import (
  "flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
  "github.com/bwmarrin/discordgo"
)

// Version is a constant that stores the Disgord version information.
const Version = "v0.0.0-alpha"

// Session is declared in the global space so it can be easily used
// throughout this program.
// In this use case, there is no error that would be returned.
var Session, _ = discordgo.New()

// Read in all configuration options from both environment variables and
// command line arguments.
func init() {

	// Discord Authentication Token
	Session.Token = os.Getenv("BOT_TOKEN")
	if Session.Token == "" {
		flag.StringVar(&Session.Token, "t", "", "Discord Authentication Token")
	}
}

func main() {
  // Declare any variables needed later.
	var err error
  
  // Print out our bot signature
	fmt.Printf("emonk - %s\n", Version)

  // Parse command line arguments
	flag.Parse()

  // Verify a Token was provided
	if Session.Token == "" {
		fmt.Printf("You must provide a Discord authentication token.\n")
		return
	}
  
  // Verify the Token is valid and grab user information
	Session.State.User, err = Session.User("@me")
	if err != nil {
		fmt.Printf("error fetching user information, %s\n", err)
    return
	}
  fmt.Printf("User: %s\n", Session.State.User)

  // Register the messageCreate func as a callback for MessageCreate events.
	Session.AddHandler(messageCreate)
  
  // Open a websocket connection to Discord and begin listening.
	err = Session.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
  
  // Wait for a CTRL-C
	fmt.Printf("Now running. Press CTRL-C to exit ...\n")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Clean up
	Session.Close()

	// Exit Normally.
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
