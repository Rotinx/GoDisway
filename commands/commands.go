package command

import (
	// "log"
	"github.com/bwmarrin/discordgo"
)

func Hello(session *discordgo.Session, message *discordgo.MessageCreate, args []string)  {
	session.ChannelMessageSend(message.ChannelID, "Love Me")
}