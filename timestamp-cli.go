package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	// NOW is now
	NOW = "now"
	// AGO is before
	AGO = "ago"
	// FROM is after
	FROM = "from now"
	// DAYS are days
	DAYS = "days"
	// HOURS are hours
	HOURS = "hours"
	// MINUTES are minutes
	MINUTES = "minutes"
)

func now() int64 {
	return time.Now().Round(time.Millisecond).UnixNano() / 1e6
}

func main() {
	rawCommand := os.Args[1:]

	if len(rawCommand) == 0 || rawCommand[0] == NOW {
		fmt.Println(now())
	} else {
		fmt.Printf("Unsupported command: %v", strings.Join(os.Args[1:], " "))
		os.Exit(-1)
	}
}
