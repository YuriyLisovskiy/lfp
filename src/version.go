// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import (
	"github.com/tcnksm/go-latest"
	"github.com/YuriyLisovskiy/lfp/src/var"
)

func init() {
	go func() {
		githubTag := &latest.GithubTag{
			Owner:      "YuriyLisovskiy",
			Repository: "lfp",
		}

		// Ignore error, because it's not important
		res, _ := latest.Check(githubTag, _var.VERSION)
		if res.Current > _var.VERSION {
			_var.VerCheckCh <- res
		}
	}()
}
