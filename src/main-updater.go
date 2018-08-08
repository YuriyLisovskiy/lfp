// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package main

import (
	"fmt"

	"github.com/YuriyLisovskiy/lfp/src/updater"
)

func main() {
	err := updater.RunCLI()
	if err != nil {
		fmt.Println(err)
	}
}
