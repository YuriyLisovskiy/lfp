// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package lfp

import (
	"time"

	"github.com/tcnksm/go-latest"
	"github.com/YuriyLisovskiy/lfp/src/config"
)

// verCheckCh is channel which gets latest.Response
var verCheckCh = make(chan *latest.CheckResponse)

// CheckTimeout is default timeout of latest.Check execution.
var checkTimeout = 2 * time.Second

func init() {
	go func() {
		githubTag := &latest.GithubTag{
			Owner:      "YuriyLisovskiy",
			Repository: "lfp",
		}

		// Ignore error, because it's not important
		res, _ := latest.Check(githubTag, config.VERSION)
		if res.Current > config.VERSION {
			verCheckCh <- res
		}
	}()
}
