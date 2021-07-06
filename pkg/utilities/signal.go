// Copyright (C) 2021-2021 The Whirlsplash Collective
// SPDX-License-Identifier: GPL-3.0-only

package utilities

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func SetupSignalHandler() {
  // https://stackoverflow.com/a/18158859/14452787

	c := make(chan os.Signal)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c

		log.Println("Killing Munch with SignalHandler")

		os.Exit(1)
	}()
}