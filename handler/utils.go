package handler

import (
	"log"
	"github.com/bwmarrin/discordgo"
)

// New adds a new command to the commands struct.
func New(name string, description string, owner bool, run func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)) {
	Commands[name] = Command{
		Name: name,
		Description: description,
		Owner: owner,
		Run: run,
	}

	log.Printf("Loaded new command [%s]", name)
}