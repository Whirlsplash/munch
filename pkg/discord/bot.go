package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"log"
)

var s *discordgo.Session

func SetupBot() {
	var err error
	s, err = discordgo.New("Bot " + viper.GetString("discord.account.token"))
	if err != nil {
		log.Fatalln("Unable to start Munch:", err)
	}

	SetupHandler()

	if err := s.Open(); err != nil {
		log.Fatalln("Unable to connect to Discord gateway:", err)
	}
	defer s.Close()
}
