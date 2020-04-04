package handler

import (
	"log"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Db variable
var Db *sqlx.DB

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

// InitDb connects database
func InitDb() {
	Db = sqlx.MustConnect("mysql", "user:password@(IP:PORT)/db");
}