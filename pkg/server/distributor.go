// Copyright (C) 2021-2021 The Whirlsplash Collective
// SPDX-License-Identifier: GPL-3.0-only

package server

import (
	"encoding/binary"
	"github.com/Whirlsplash/munch/pkg/server/utils"
	"github.com/spf13/viper"
	"log"
	"net"
	"strconv"
)

func doDistributor() string {
	tcpAddr, _ := net.ResolveTCPAddr(
		"tcp",
		viper.GetString("worlds.server.ip")+":"+viper.GetString("worlds.server.port"),
	)
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)

	// PROPREQ
	conn.Write([]byte("\x03\xff\x0a"))
	reply := make([]byte, 1024)
	conn.Read(reply)
	log.Println("AutoServer: PROPREQ")

	// SESSINIT
	conn.Write(utils.EncodeSessionInitialization(
		viper.GetString("worlds.account.username"),
		viper.GetString("worlds.account.password"),
		utils.AutoServer,
	))
	reply = make([]byte, 1024)
	conn.Read(reply)
	log.Println("AutoServer: SESSINIT")

	// PROPUPD
	conn.Write(utils.EncodePropertyUpdate(viper.GetString("worlds.account.avatar")))
	reply = make([]byte, 1024)
	conn.Read(reply)
	log.Println("AutoServer: PROPUPD")

	// ROOMIDRQ
	conn.Write([]byte(
		"\x26\x01\x14\x22\x47\x72\x6f\x75\x6e\x64\x5a\x65\x72\x6f\x23\x73" +
			"\x74\x61\x69\x72\x63\x61\x73\x65\x31\x3c\x64\x69\x6d\x65\x6e\x73" +
			"\x69\x6f\x6e\x2d\x31\x3e",
	))
	reply = make([]byte, 1024)
	conn.Read(reply)
	port := strconv.Itoa(int(binary.BigEndian.Uint16(reply[44:46])))
	log.Println("AutoServer: ROOMIDRQ")

	conn.Close()

	return port
}
