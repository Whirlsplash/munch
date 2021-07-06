// Copyright (C) 2021-2021 The Whirlsplash Collective
// SPDX-License-Identifier: GPL-3.0-only

package discord

import (
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/spf13/viper"
)

func Do() {
	commands := &Bot{}

	bot.Run(
		viper.GetString("discord.token"),
		commands,
		func(ctx *bot.Context) error {
			ctx.HasPrefix = bot.NewPrefix(viper.GetString("discord.prefix"))

			return nil
		},
	)
}
