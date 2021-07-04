// Copyright (C) 2021-2021 The Whirlsplash Collective
// SPDX-License-Identifier: GPL-3.0-only

package server

import (
	"github.com/Whirlsplash/munch/pkg/server/utils"
	"github.com/spf13/viper"
	"log"
	"net"
)

func doHub(port string) {
	tcpAddr, _ := net.ResolveTCPAddr(
		"tcp",
		(viper.GetString("worlds.server.ip") + ":" + port),
	)
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)

	// PROPREQ
	conn.Write([]byte("\x03\xff\x0a"))
	reply := make([]byte, 1024)
	conn.Read(reply)
	log.Println("RoomServer: PROPREQ")

	// SESSINIT
	conn.Write(utils.EncodeSessionInitialization(
		viper.GetString("worlds.account.username"),
		viper.GetString("worlds.account.password"),
		utils.RoomServer,
	))
	reply = make([]byte, 1024)
	conn.Read(reply)
	log.Println("RoomServer: SESSINIT")

	// SUB-DIST
	conn.Write([]byte(
		"\x0d\x01\x16\x00\x02\x00\x00\x00\xfa\x00\x00\x0a\xf8" +
			"\x0d\x01\x16\x00\x01\x00\x00\x02\xee\x00\x00\x0c\xf9" +
			"\x0d\x01\x16\x00\x07\x00\x00\x36\xaf\x00\x00\x09\xbd" +
			"\x0d\x01\x16\x00\x06\xf1\x77\x00\x00\x01\x90\x0d\xcf\x0d\x01\x16" +
			"\x98\xee\x00\x00\x00\xc8\x00\x28\x0e\x67\x0d\x01\x16\x00\x04\x00" +
			"\x00\x00\xfa\x00\x00\x0c\xb7\x0d\x01\x16\x00\x03\x00\x00\x02\xee" +
			"\x00\x00\x0c\x30\x0d\x01\x16\x00\x09\x01\xf4\x03\xe8\x01\xf4\x09" +
			"\xd5\x0d\x01\x16\x00\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x0f\x01" +
			"\x12\x00\x08\x00\x01\x04\xcc\x09\xa1\x00\x00\x01\x25\x07\x01\x18" +
			"\x00\x01\x0c\xf9\x07\x01\x18\x00\x03\x0c\x30\x07\x01\x18\x00\x09" +
			"\x09\xd5\x07\x01\x18\x00\x07\x09\xbd\x07\x01\x18\x00\x06\x0d\xcf" +
			"\x07\x01\x18\x98\xee\x0e\x67\x07\x01\x18\x00\x04\x0c\xb7\x07\x01" +
			"\x18\x00\x02\x0a\xf8",
	))
	reply = make([]byte, 1024)
	conn.Read(reply)
	log.Println("RoomServer: SUBSCRIB")

	// Listen for whispers
	for {
		reply = make([]byte, 1024)
		conn.Read(reply)
		if reply[2] == 17 {
			decodedText := utils.DecodeText(reply)
			log.Printf(
				"RoomServer: WHISPER: %s: %s\n",
				decodedText.Author, decodedText.Message,
			)
		}
	}

	conn.Close()
}
