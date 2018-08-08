// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import (
	"os"
	"fmt"
	"time"
	"io/ioutil"

	"github.com/YuriyLisovskiy/lfp/src/var"
	"github.com/YuriyLisovskiy/lfp/src/updater"
)

func RunCLI() error {

	// if there are no arguments
	if len(os.Args) == 1 {
		print(_var.ABOUT)
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

	if *updatePtr != "" {
		return processUpdate()
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

func processUpdate() error {
	return updater.StartUpdate(*updatePtr)
}

func processVersion() error {
	fmt.Printf("%s version %s\n", _var.PROGRAM_NAME, _var.VERSION)
	select {
	case <-time.After(_var.CheckTimeout):
		// Do nothing
	case res := <-_var.VerCheckCh:
		if res != nil {
			fmt.Printf("Latest version of %s is %s, please update the %s tool\n",
				_var.PROGRAM_NAME, res.Current, _var.PROGRAM_NAME,
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
