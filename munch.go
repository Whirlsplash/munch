// Copyright (C) 2021-2021 The Whirlsplash Collective
// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"github.com/Whirlsplash/munch/pkg/config"
	"github.com/Whirlsplash/munch/pkg/discord"
	"github.com/Whirlsplash/munch/pkg/server"
	"sync"
)

func main() {
	config.Setup()

	var wg sync.WaitGroup

	// Discord bot
	wg.Add(1)
	go func() {
		discord.Do()
		wg.Done()
	}()
	wg.Add(1)

	// Worlds bot
	go func() {
		server.Do()
		wg.Done()
	}()

	wg.Wait()
}
