// Copyright (C) 2021-2021 The Whirlsplash Collective
// SPDX-License-Identifier: GPL-3.0-only

package discord

func Do() {
	SetupBot()
	CleanupCommands()
	SetupCommands()

	select {}
}
