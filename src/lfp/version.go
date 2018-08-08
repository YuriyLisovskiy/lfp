// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package lfp

import (
	"time"

	"github.com/tcnksm/go-latest"
)

const (
	PROGRAM_NAME = "lfp"
	VERSION      = "1.0.1-beta"
	ABOUT        = PROGRAM_NAME + " " + VERSION + "\n\n" +
		"A utility for licensing existing project\n" +
		"Yuriy Lisovskiy (c) 2018 <https://github.com/YuriyLisovskiy>\n\n"
)

// verCheckCh is channel which gets latest.Response
var verCheckCh = make(chan *latest.CheckResponse)

// CheckTimeout is default timeout of latest.Check execution.
var CheckTimeout = 2 * time.Second

func init() {
	go func() {
		githubTag := &latest.GithubTag{
			Owner:      "YuriyLisovskiy",
			Repository: "lfp",
		}

		// Ignore error, because it's not important
		res, _ := latest.Check(githubTag, VERSION)
		if res.Current != VERSION {
			verCheckCh <- res
		}
	}()
}
