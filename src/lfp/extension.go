// Copyright (c) 2018 Yuriy Lisovskiy
// Distributed under the MIT software license, see the accompanying
// file LICENSE or https://opensource.org/licenses/MIT

package lfp

import "strings"

func getComments(ext string) (string, string, error) {
	ext = strings.ToLower(ext)
	switch ext {
	case "h", "hpp", "hh", "hxx", "ccpph", "cc", "cxx", "c++", "cp", "cpp", "gvy", "gy", "gsh", "groovy",
		"fsi", "fsx", "fsscript", "fs", "kts", "kt", "mm", "m", "phtml", "php3", "php4", "php5", "php7",
		"phps", "php-s", "php", "scss", "sass", "postcss", "css", "dtsi", "dts", "tesc", "tese", "geom",
		"frag", "comp", "vert", "sc", "scala", "uci", "upkg", "uc", "as", "go", "java", "js", "json", "swift",
		"cfc", "cu", "cuh", "d", "dart", "jai", "jsx", "less", "lds", "qcl", "qml", "rs", "styl", "ts", "tsx",
		"y", "proto", "zig":
		return "//", "", nil
	case "htm", "xhtml", "html", "markdown", "mdown", "mkdn", "mkd", "mdwn", "mdtxt", "mdtext", "text",
		"rmd", "md", "handlebars", "hbs", "xml", "polly", "cshtml", "rhtml":
		return "<!--", "-->", nil
	case "yaml", "yml", "pm", "t", "pod", "pl", "rdata", "rds", "rda", "r", "s", "asm", "makefile", "in", "awk",
		"sh", "csh", "nim", "sls", "tcl", "toml", "py", "rb", "coffee", "jl", "nix", "arr", "tf":
		return "#", "", nil
	case "lhs", "hs", "adb", "ads", "lagda", "agda", "lidr", "idr", "hlean", "lean", "sql", "lua":
		return "--", "", nil
	case "pp", "inc", "pas":
		return "{*", "*}", nil
	case "vbhtml", "asax", "ascx", "ashx", "asmx", "axd", "dbml", "edmx", "resx", "svc", "aspx", "asp":
		return "<%--", "--%>", nil
	case "btm", "cmd", "bat":
		return "REM", "", nil
	case "hrl", "erl", "pro", "p", "sty", "tex", "oz":
		return "%", "", nil
	case "4th", "fr", "frt", "fth", "f83", "fb", "fpm", "e4", "rx", "ft", "forth":
		return "(", ")", nil
	case "for", "ftn", "f77", "pfo", "f", "f08", "f90", "f95", "f03":
		return "!", "", nil
	case "el", "lsp", "scm", "ss", "rkt", "lisp", "ini":
		return ";", "", nil
	case "mli", "ml", "nb", "wl", "v", "sml":
		return "(*", "*)", nil
	case "txt", "hex", "ihex", "rst":
		return "", "", nil
	case "mustache":
		return "{{!", "}}", nil
	}
	return "", "", ErrCommentNotFound
}

func getExtension(path string) string {
	pos := strings.LastIndex(path, ".")
	if pos == -1 {
		pos = strings.LastIndex(path, "/")
		if pos != -1 {
			return path[pos+1:]
		}
	} else {
		return path[pos+1:]
	}
	return path
}
