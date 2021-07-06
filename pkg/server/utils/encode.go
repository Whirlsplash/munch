// Copyright (C) 2021-2021 The Whirlsplash Collective
// SPDX-License-Identifier: GPL-3.0-only

package utils

import "github.com/Whirlsplash/munch/pkg/utilities"

const (
	AutoServer = iota
	RoomServer
)

func EncodeSessionInitialization(username string, password string, serverType int) []byte {
	data := ""

	// Data length
	if serverType == AutoServer {
		data += "\x2b"
	} else {
		data += "\x2c"
	}

	// Command header and other stuff we don't need to worry about
	data += "\x01\x06\x03\x02\x32\x34\x09\x0a\x32\x30\x32\x30\x30\x33\x31" +
		"\x32\x30\x30"

	if serverType == RoomServer {
		data += "\x07\x02\x32\x34"
	}

	data +=
		// Username
		"\x02" + string(rune(len(username))) + utilities.ISO88591ToString(username) +

			// Password
			"\x06" + string(rune(len(password))) + utilities.ISO88591ToString(password)

	if serverType == AutoServer {
		data += "\x0c\x01\x31"
	}

	return []byte(data)
}

func EncodeBuddyListUpdate(buddy string) []byte {
	// Command header
	data := "\x01\x1d"

	// Buddy UTF-8 length and UTF-8
	data += string(rune(len(buddy))) + utilities.ISO88591ToString(buddy)

	// Buddy "add"
	data += "\x01"

	// Data length
	data = (string(rune(len(data) + 1))) + data

	return []byte(data)
}

func EncodePropertyUpdate(avatar string) []byte {
	// Command header and extra stuff we don't need to worry about
	data := "\x01\x0f\x00\x05\x40\x01"

	// Avatar UTF-8 length and UTF-8
	data += string(rune(len("avatar:"+avatar))) + utilities.ISO88591ToString("avatar:"+avatar)

	// Data length
	data = (string(rune(len(data) + 1))) + data

	return []byte(data)
}
