// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import (
	"os"
	"log"

	"github.com/YuriyLisovskiy/lofp/src/args"
)

// logErr: if DEBUG is true, prints trace, else prints error message.
func logErr(err error) {
	if DEBUG {
		log.Panic(err)
	} else {
		log.Println(err)
	}
}

func RunCLI() {

	// parse command line arguments
	if err := args.Lofp.Parse(os.Args[1:]); err != nil {
		logErr(err)
		return
	}

	arguments := args.Lofp.Args()

	// check if there is no errors in given arguments
	if err := args.ValidateArgs(arguments); err != nil {
		logErr(err)
		return
	}

	// retrieve paths from arg set
	paths := args.RetrievePaths(arguments)

	// check if paths contain arguments
	if err := args.ValidatePaths(paths); err != nil {
		logErr(err)
		return
	}

	// run main process
	process(paths)
}

func process(paths []string) {
	for _, path := range paths {
		println(path)
	}
}
