// Copyright (C) 2021-2021 The Whirlsplash Collective
// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"github.com/Whirlsplash/munch/pkg/config"
	"github.com/Whirlsplash/munch/pkg/server"
)

func main() {
	config.Setup()
	server.Do()
}
