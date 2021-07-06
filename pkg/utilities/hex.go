// Copyright (C) 2021-2021 The Whirlsplash Collective
// SPDX-License-Identifier: GPL-3.0-only

package utilities

import "unicode/utf8"

// https://stackoverflow.com/a/43461796/14452787
//
// I cannot explain the joy that flew through me once I found this ISO-8859-1
// (Latin-1) to UTF-8 converter. I was looking for a valid solution on how to
// insert the username and password into a hex-escaped string sequence for
// EncodeSessionInitialization for hours and this finally solved all of my
// problems.
func ISO88591ToString(iso string) string {
	var utf []rune
	for i := 0; i < len(iso); i++ {
		r := iso[i]
		if utf == nil {
			if r < utf8.RuneSelf {
				continue
			}
			utf = make([]rune, len(iso))
			for j, r := range iso[:i] {
				utf[j] = rune(r)
			}
		}
		utf[i] = rune(r)
	}
	if utf == nil {
		return string(iso)
	}
	return string(utf)
}
