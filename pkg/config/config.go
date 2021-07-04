// Copyright (C) 2021-2021 The Whirlsplash Collective
// SPDX-License-Identifier: GPL-3.0-only

package config

import (
	"github.com/spf13/viper"
	"log"
)

func Setup() {
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Panicln("Cannot read configuration file:", err)
		} else {
			log.Panicln("Read configuration file but an error occurred anyway:", err)
		}
	}

	viper.WatchConfig()
}
