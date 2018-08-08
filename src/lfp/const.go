// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package lfp

const (
	PROGRAM_NAME = "lfp"
	VERSION      = "1.0.3"
	ABOUT        = PROGRAM_NAME + " " + VERSION + "\n\n" +
		"A utility for licensing the project\n" +
		"Copyright (c) 2018 Yuriy Lisovskiy, <https://github.com/YuriyLisovskiy>\n\n"
	LICENSE_HELP =
`
			Available Licenses

Apache:
 - 'apache-2.0' - Apache License Version 2.0

> Requires <authors> field.


BSD:
 - 'bsd-2-clause' - BSD 2-Clause License
 - 'bsd-3-clause' - BSD 3-Clause License

> Requires <authors> field.


Eclipse:
 - 'epl-2.0' - Eclipse Public License - v2.0


GNU:
 - 'gpl-2.0' - GNU General Public License v2.0
 - 'gpl-3.0' - GNU General Public License v3.0
 - 'agpl-3.0' - GNU Affero General Public License v3.0
 - 'lgpl-2.1' - GNU Lesser General Public License v2.1
 - 'lgpl-3.0' - GNU Lesser General Public License v3.0

> Requires <program description> and <authors> fields except lgpl-3.0.


MIT:
 - 'mit' - MIT License

> Requires <authors> field.


Mozilla:
 - 'mpl-2.0' - Mozilla Public License Version 2.0

Unlicense:
 - 'unlicense'

`
)
