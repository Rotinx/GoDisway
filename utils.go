package main

import (
	command "bot/commands"
	event "bot/events"
	"bot/handler"
	"encoding/json"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)


// Configuration contains data from config file.
type configuration struct {
	Token string
	Prefix string
}

var token = ""

// LoadConfig setups up the discord bots configs.
func LoadConfig() {	
	configuration := configuration{}
	file, err := os.Open("./config.json") 
	if err != nil { 
		log.Fatal(err)
	}  
	decoder := json.NewDecoder(file) 
	err = decoder.Decode(&configuration) 
	if err != nil { 
		log.Fatal(err)
	}  

	token = configuration.Token
	event.Prefix = configuration.Prefix
}

// Init discord client
func Init() {
	client, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("error creating Discord session,", err)
		return
	}

	err = client.Open()
	if err != nil {
		log.Fatal("error opening Discord session,", err)
		return
	}

	client.AddHandler(event.MessageCreate)

	u, err := client.User("@me")
	if err != nil {
		log.Fatal("error fetching Discord user,", err)
		return
	}
	
	log.Printf("Bot logged in. %s", u)	

	// Load commands.
	handler.New("love","Test command", false, command.Hello)
}