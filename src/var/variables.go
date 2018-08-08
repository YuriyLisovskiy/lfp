// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package _var

import (
	"time"

	"github.com/tcnksm/go-latest"
)

// verCheckCh is channel which gets latest.Response
var VerCheckCh = make(chan *latest.CheckResponse)

// CheckTimeout is default timeout of latest.Check execution.
var CheckTimeout = 2 * time.Second
