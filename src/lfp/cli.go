// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package lfp

import (
	"os"
	"fmt"
	"time"
	"io/ioutil"

	"github.com/YuriyLisovskiy/lfp/src/config"
)

func RunCLI() error {

	// if there are no arguments
	if len(os.Args) == 1 {
		print(config.ABOUT)
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
		return processVersion()
	}

	if *helpPtr {
		lfp.Usage()
		return nil
	}
	return processLicensing()
}

func processVersion() error {
	fmt.Printf("%s version %s\n", config.PROGRAM_NAME, config.VERSION)
	select {
	case <-time.After(checkTimeout):
		// Do nothing
	case res := <-verCheckCh:
		if res != nil {
			fmt.Printf("Latest version of %s is %s, please update the %s tool\n",
				config.PROGRAM_NAME, res.Current, config.PROGRAM_NAME,
			)
		}
	}
	return nil
}

func processLicensing() error {

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
	return nil
}
