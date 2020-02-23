package event

import (
	"bot/handler"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Prefix is a global variable which contains the bots command prefix.
var Prefix = ""

// MessageCreate is a message event :shrug:
func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if(message.Author.Bot || !strings.HasPrefix(message.Content, Prefix)) {
		return
	}

	args := strings.Fields(strings.TrimPrefix(message.Content, Prefix))
	
	for _, command := range handler.Commands {
		if(command.Name == args[0]) {
			command.Run(session, message, args)
		}
	}
}