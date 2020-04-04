package command

import (
	"log"

	util "bot/handler"

	"github.com/bwmarrin/discordgo"
)

func LoveMe(session *discordgo.Session, message *discordgo.MessageCreate, args []string)  {
	session.ChannelMessageSend(message.ChannelID, "Love Me")
	
	user := []util.UserSchema{}

	err := util.Db.Select(&user, "select * from users")
	
	log.Println(user[0].ID);
	

	if err != nil {
		log.Fatal(err)
	}
}