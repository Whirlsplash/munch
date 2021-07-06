// Copyright (C) 2021-2021 The Whirlsplash Collective
// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"github.com/Whirlsplash/munch/pkg/config"
	"github.com/Whirlsplash/munch/pkg/discord"
	"github.com/Whirlsplash/munch/pkg/server"
	"github.com/spf13/viper"
	"sync"
)

func main() {
	config.Setup()

	var wg sync.WaitGroup

	// Discord bot
	if viper.GetBool("discord.enable") {
		wg.Add(1)
		go func() {
			discord.Do()
			wg.Done()
		}()
	}

	// Worlds bot
	if viper.GetBool("worlds.enable") {
		wg.Add(1)
		go func() {
			server.Do()
			wg.Done()
		}()
	}

	wg.Wait()
}
