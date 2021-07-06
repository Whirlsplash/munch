// Copyright (C) 2021-2021 The Whirlsplash Collective
// SPDX-License-Identifier: GPL-3.0-only

package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"log"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Pong!",
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ping": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Pong!",
				},
			})
		},
	}
)

func CleanupCommands() {
	cmds, err := s.ApplicationCommands(s.State.User.ID, "")
	if err != nil {
		log.Panicln("Unable to obtain applications commands:", err)
	}

	for _, e := range cmds {
		if err := s.ApplicationCommandDelete(s.State.User.ID, "", e.ID); err != nil {
			log.Panicf("Cannot delete command '%v': %v", e.Name, err)
		} else {
			log.Printf("Deleted command '%v'", e.Name)
		}
	}
}

func SetupCommands() {
	for _, v := range commands {
		if _, err := s.ApplicationCommandCreate(s.State.User.ID, "", v); err != nil {
			log.Panicf("Cannot create command '%v': %v", v.Name, err)
		} else {
			log.Printf("Created command '%v'", v.Name)
		}
	}

	if viper.GetBool("discord.account.cleanup_app_commands") {
		CleanupCommands()
	}
}
