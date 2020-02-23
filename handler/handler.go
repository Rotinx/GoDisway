package handler

import (
	"github.com/bwmarrin/discordgo"
)

// Command struct contains all the commands available for the bot.
type Command struct {
	Name string
	Description string
	Owner bool
	Run func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)
}

// Commands struct contains all the commands available for the bot.
var Commands = make(map[string]Command);



