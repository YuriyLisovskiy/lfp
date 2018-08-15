// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package lfp

import (
	"os"
	"fmt"
	"time"
	"io/ioutil"
)

func RunCLI() error {

	// if there are no arguments
	if len(os.Args) == 1 {
		print(ABOUT)
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

	if *licensesPtr {
		return processLicenseHelp()
	}

	return processLicensing()
}

func processVersion() error {
	fmt.Printf("%s version %s\n", PROGRAM_NAME, VERSION)
	select {
	case <-time.After(checkTimeout):
		// Do nothing
	case res := <-verCheckCh:
		if res != nil {
			fmt.Printf("Latest version of %s is %s, please update the %s tool\n",
				PROGRAM_NAME, res.Current, PROGRAM_NAME,
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
	cfg, err := parseConfig(cfgData, getExtension(*configPtr))
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
	fmt.Println("Licensing the project:", cfg.ProjectRoot)

	// run main process
	err = process(cfg)
	if err != nil {
		return err
	}
	fmt.Println("Done.")
	return nil
}

func processLicenseHelp() error {
	fmt.Print(LICENSE_HELP)
	return nil
}
