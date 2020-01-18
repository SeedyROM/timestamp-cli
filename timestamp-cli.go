package main

import (
	"fmt"
	"os"
	"strings"
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

// ParseExpression parses a time expression
func ParseExpression(expression string) {
	fmt.Println(expression)
}

// ParseCommand parses the raw command given to the CLI
func ParseCommand(rawCommand []string) {
	length := len(rawCommand)

	var expression []string

	for i := 0; i < length; i++ {
		part := rawCommand[i]
		if SliceContains(part, []string{AGO, FROM}) != -1 {
			ParseExpression(strings.Join(expression, " "))
			break
		} else {
			expression = append(expression, part)
		}
	}
}

func main() {
	rawCommand := os.Args[1:]

	// Default to now or expect now
	if len(rawCommand) == 0 || rawCommand[0] == NOW {
		fmt.Println(Now())
	} else {
		ParseCommand(rawCommand)
		// fmt.Printf("Unsupported command: %v\n", strings.Join(os.Args[1:], " "))
		// os.Exit(-1)
	}
}
