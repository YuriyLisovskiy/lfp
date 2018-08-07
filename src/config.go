// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package src

import "strings"

type Author struct {
	Name string `yaml:"name"`
	Year string `yaml:"year"`
}

// Configuration for licensing the project, see sample/sample.yml
type Config struct {
	// Author
	Authors []Author `yaml:"authors"`

	// Program name
	ProgramName string `yaml:"program_name"`

	// Short description of the program
	ProgramDescription string `yaml:"program_description"`

	// Paths to add license notice
	Paths []string `yaml:"paths"`

	// License name, see docs/licenses.md for details.
	License string `yaml:"license"`

	// An absolute project root path.
	ProjectRoot string `yaml:"project_root"`

	// Create custom license notice.
	CustomLicenseNotice string `yaml:"custom_license_notice"`

	// Add LICENSE file to project root.
	AddLicenseFile bool `yaml:"add_license_file"`

	// Add license notice to every specified path.
	AddLicenseNotice bool `yaml:"add_license_notice"`
}

func (cfg Config) validate() error {
	if len(cfg.Authors) == 0 {
		return ConfigErrAuthorRequired
	}
	for _, author := range cfg.Authors {
		if author.Name == "" || author.Year == "" {
			return ConfigErrYearsAuthors
		}
	}
	if len(cfg.Paths) == 0 {
		return ConfigErrPathsRequired
	}
	if cfg.License == "" {
		return ConfigErrLicenseRequired
	}
	if cfg.ProjectRoot == "" {
		return ConfigErrProjectRootRequired
	}
	if !cfg.AddLicenseFile && !cfg.AddLicenseNotice {
		return ConfigErrAddLicenseFileNoticeRequired
	}
	return nil
}

func removeItem(items []string, pos int) ([]string, error) {
	if pos >= len(items) || pos < 0 {
		return nil, ErrIndexOutOfRange
	}
	if pos+1 < len(items) {
		return append(items[:pos], items[pos+1:]...), nil
	}
	return items[:pos], nil
}

func (cfg Config) normalize() (Config, error) {
	if strings.HasSuffix(cfg.ProjectRoot, "/") {
		cfg.ProjectRoot = cfg.ProjectRoot[:len(cfg.ProjectRoot)-1]
	}
	if cfg.ProgramName == "" {
		cfg.ProgramName = cfg.ProjectRoot[strings.LastIndexByte(cfg.ProjectRoot, byte('/'))+1:]
	}
	var err error
	for i, path := range cfg.Paths {
		if strings.HasPrefix(path, "/") {
			if len(path) > 1 {
				cfg.Paths[i] = path[1:]
			} else {
				cfg.Paths, err = removeItem(cfg.Paths, i)
				if err != nil {
					return Config{}, err
				}
			}
		} else if strings.HasPrefix(path, "./") {
			if len(path) > 2 {
				cfg.Paths[i] = path[2:]
			} else {
				cfg.Paths, err = removeItem(cfg.Paths, i)
				if err != nil {
					return Config{}, err
				}
			}
		}
	}
	return cfg, nil
}
