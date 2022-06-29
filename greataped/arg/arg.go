package arg

import (
	"flag"
	"os"
)

var (
	Values []string
)

var (
	Verbose bool
	VeryVerbose bool
	VeryVeryVerbose bool
	VeryVeryVeryVerbose bool
	VeryVeryVeryVeryVerbose bool
	VeryVeryVeryVeryVeryVerbose bool
)

var (
	help bool
)


func init() {
	flag.BoolVar(&Verbose,                     "v",      false,                          "verbose logs outputted")
	flag.BoolVar(&VeryVerbose,                 "vv",     false,                     "very verbose logs outputted")
	flag.BoolVar(&VeryVeryVerbose,             "vvv",    false,                "very very verbose logs outputted")
	flag.BoolVar(&VeryVeryVeryVerbose,         "vvvv",   false,           "very very very verbose logs outputted")
	flag.BoolVar(&VeryVeryVeryVeryVerbose,     "vvvvv",  false,      "very very very very verbose logs outputted")
	flag.BoolVar(&VeryVeryVeryVeryVeryVerbose, "vvvvvv", false, "very very very very very verbose logs outputted")

	flag.BoolVar(&help, "help", false, "outputs help message")

	flag.Parse()

	Values = flag.Args()

	// --help
	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}
}
