// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import (
	"os"
	"io/ioutil"

	"github.com/YuriyLisovskiy/lofp/src/static"
)

func RunCLI() error {

	// if there are no arguments
	if len(os.Args) == 1 {
		print(static.ABOUT)
		lfp.Usage()
		return nil
	}

	// parse command line arguments
	if err := lfp.Parse(os.Args[1:]); err != nil {
		return nil
	}
	arguments := lfp.Args()

	// check if there is no errors in given arguments
	if err := validateArgs(arguments); err != nil {
		return err
	}

	if *versionPtr {
		println(static.VERSION)
	} else if *helpPtr {
		lfp.Usage()
	} else {

		// read and parse config file
		cfgData, err := ioutil.ReadFile(*configPtr)
		if err != nil {
			return err
		}
		cfg, err := parseConfig(cfgData)
		if err != nil {
			return err
		}

		// validate and normalize configuration
		err = cfg.validate()
		if err != nil {
			return err
		}
		cfg, err = cfg.normalize()
		if err != nil {
			return err
		}

		// run main process
		err = process(cfg)
		if err != nil {
			return err
		}
	}
	return nil
}

func process(cfg Config) error {
	return processPaths(cfg)
}
