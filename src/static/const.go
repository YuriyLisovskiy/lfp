// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package static

const (
	PROGRAM_NAME = "lofp"
	VERSION      = "v0.0.1-alpha"
	ABOUT        = PROGRAM_NAME + " " + VERSION + "\n" +
		"A utility for licensing existing project\n" +
		"Yuriy Lisovskiy (c) 2018 <https://github.com/YuriyLisovskiy>\n\n"
	USAGE = "USAGE:\n" +
		"    " + PROGRAM_NAME + " [FLAGS] [OPTIONS]\n\n" +
		"FLAGS:\n" +
		"    -p\t\tSets paths to add license notice\n" +
		"    -l\tSpecify license\n" +
		"    -r\tSets project root directory\n\n" +
		"OPTIONS:\n" +
		"    -lf\tAdd LICENSE file to the project root folder\n" +
		"    -j\t\tAdd license notice to specified file(s) and(or) folder(s)\n" +
		"    -h\t\tPrint usage\n" +
		"    -v\t\tPrint version\n"
)
