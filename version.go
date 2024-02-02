package main

import (
	"fmt"
)

const (
	appName = "mikanFixer"
	desc    = "Sources: https://www.github.com/thank243/mikanFixer"
)

var (
	version = "dev"
	date    = "unknown"
)

func getVersion() string {
	return fmt.Sprintf("%s %s, built at %s\n%s", appName, version, date, desc)
}
