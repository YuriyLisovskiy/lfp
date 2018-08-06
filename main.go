// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package main

import (
	"log"

	"github.com/YuriyLisovskiy/lfp/src"
)

func main() {
	if err := src.RunCLI(); err != nil {
		if src.DEBUG {
			log.Panic(err)
		} else {
			log.Println(err)
		}
	}
}
