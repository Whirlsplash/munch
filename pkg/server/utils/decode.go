// Copyright (C) 2021-2021 The Whirlsplash Collective
// SPDX-License-Identifier: GPL-3.0-only

package utils

import "strconv"

type Text struct {
	Author  string
	Message string
}

// Decode a TEXT command from hexadecimal to a Text struct
func DecodeText(data []byte) Text {
	authorLength, _ := strconv.ParseInt(strconv.Itoa(int(data[4])), 16, 32)
	authorLast := authorLength + 5
	messageLength, _ := strconv.ParseInt(strconv.Itoa(int(data[authorLast])), 16, 32)
	messageStart := authorLast + 1

	return Text{
		Author:  string(data[5:(authorLast)]),
		Message: string(data[messageStart : messageStart+(messageLength)]),
	}
}
