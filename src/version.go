// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import (
	"time"

	"github.com/tcnksm/go-latest"
	)

const (
	PROGRAM_NAME = "lfp"
	VERSION      = "1.0-beta"
	ABOUT        = PROGRAM_NAME + " " + VERSION + "\n\n" +
		"A utility for licensing existing project\n" +
		"Yuriy Lisovskiy (c) 2018 <https://github.com/YuriyLisovskiy>\n\n"
)

// verCheckCh is channel which gets go-latest.Response
var verCheckCh = make(chan *latest.CheckResponse)

// CheckTimeout is default timeout of go-latest.Check execution.
var CheckTimeout = 2 * time.Second

func init() {

	go func() {
		githubTag := &latest.GithubTag{
			Owner:      "YuriyLisovskiy",
			Repository: "lfp",
		}

		// Ignore error, because it's not important
		res, _ := latest.Check(githubTag, VERSION)
		verCheckCh <- res
	}()
}
