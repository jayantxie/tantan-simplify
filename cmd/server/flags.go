package main

import (
	"flag"
	"fmt"
	"os"
	"tantan-simplify/version"
)

type optionFlags struct {
	confFile       string
	showVersion    bool
	BackTrackLevel string
}

var flags optionFlags

func parseFlags() {
	flagSet := flag.NewFlagSet("matrix-batch", flag.ExitOnError)
	flagSet.StringVar(&flags.confFile, "c", "conf/dev/config.toml", "config file path")
	flagSet.BoolVar(&flags.showVersion, "v", false, "show current version")
	flagSet.StringVar(&flags.BackTrackLevel, "bt", "warn", "show backtrack in the log >= {level}")
	flagSet.Parse(os.Args[1:])
}

func showVersion() {
	if flags.showVersion {
		printVersion()
		return
	}
}

func printVersion() {
	fmt.Printf("Version: %s\nBuildDate: %s\nBuildCommit: %s\n", version.Version, version.BuildDate, version.BuildCommit)
}
