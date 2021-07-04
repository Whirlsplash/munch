// Copyright (C) 2021-2021 The Whirlsplash Collective
// SPDX-License-Identifier: GPL-3.0-only

package server

func Do() {
	// The Distributor isn't actually needed for Munch to work, the only reason
	// that we use it is because it gives us the correct Hub port to connect to.
	doHub(doDistributor())
}
