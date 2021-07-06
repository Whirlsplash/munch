package discord

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func SetupHandler() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println(r.User.Username, "is live!")
	})

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}
