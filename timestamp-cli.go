package main

import (
	"errors"
	"fmt"
	"github.com/nleeper/goment"
	"os"
	"strconv"
	"strings"
)

// Error constants
const (
	Success        = 0
	InvalidCommand = 1 - iota
	ParseError
	Exited
)

// Time constants
const (
	// Now is now
	Now = "now"
	// Ago is before
	Ago = "ago"
	// From is after
	From = "from"
	// Minutes are minutes
	Minutes = "minutes"
	// MinuteAmount is the amount of milliseconds in a minute
	MinuteAmount = 1000 * 60
	// Hours are hours
	Hours = "hours"
	// HourAmount is the amount of milliseconds in an hour
	HourAmount = MinuteAmount * 60
	// Days are days
	Days = "days"
	// DayAmount is the amount of milliseconds in an day
	DayAmount = HourAmount * 24
	// Months are months
	Months = "months"
	// MonthAmount is the amount of milliseconds in a month
	MonthAmount = DayAmount * 30
	// Years are years
	Years = "years"
	// YearAmount is the amount of milliseconds in a year
	YearAmount = MonthAmount * 12
)

// ParseExpression gets the amount or calculation from the expression
func ParseExpression(expression string) (string, string, error) {
	parts := strings.Split(expression, " ")
	length := len(parts)
	var amount []string

	for i := 0; i < length; i++ {
		part := parts[i]

		if SliceContains(part, []string{Years, Months, Days, Hours, Minutes}) == -1 {
			amount = append(amount, parts[i])
		} else {
			return strings.Join(amount, " "), parts[i], nil
		}
	}

	return "", "", errors.New("Invalid expression")
}

// GetTimeStamp gets a timestamp from a total value and an amount
func GetTimeStamp(total string, amount string) int64 {
	var amountInMillis int64

	switch amount {
	case Minutes:
		amountInMillis = MinuteAmount
	case Hours:
		amountInMillis = HourAmount
	case Days:
		amountInMillis = DayAmount
		break
	case Months:
		amountInMillis = MonthAmount
	case Years:
		amountInMillis = YearAmount
	}

	totalInt, err := strconv.Atoi(total)

	if err != nil {
		fmt.Printf("Invalid amount: %v\n", amount)
		os.Exit(ParseError)
	}

	return int64(totalInt) * amountInMillis
}

// HandleCommand parses a time expression
func HandleCommand(expression string, when string) int64 {
	total, amount, err := ParseExpression(expression)

	// Handle invalid expressions
	if err != nil {
		fmt.Printf("Invalid expression: %v\n", expression)
		os.Exit(ParseError)
	}

	// Get the amount of time in milliseconds from the command
	result := GetTimeStamp(total, amount)

	result := goment.New(GetNow())

	// If it's in the past subtract
	if when == Ago {
		return GetNow() - result
	}
	// If it's in the future add
	return GetNow() + result
}

// EvaluateCommand parses the raw command given to the CLI
func EvaluateCommand(rawCommand []string) (int64, error) {
	length := len(rawCommand)

	var expression []string

	for i := 0; i < length; i++ {
		part := rawCommand[i]

		if part == "-h" || part == "--help" {
			usage()
			os.Exit(Success)
		}

		if SliceContains(part, []string{Ago, From}) != -1 {
			return HandleCommand(strings.Join(expression, " "), part), nil
		}

		expression = append(expression, part)
	}

	return -1, fmt.Errorf("Invalid command: %v", strings.Join(rawCommand, " "))
}

// Display the program's version and source package information
func printVersion() {
	fmt.Println("github.com/SeedyROM/timestamp v1.0.0")
}

func help() {
	fmt.Println("Usage:")
	fmt.Println("\ttimestamp 5 minutes from now")
	fmt.Println("\ttimestamp 10 hours ago")
}

// Pretty print the usage for the command
func usage() {
	description := "A command to generate timestamps (in milliseconds or seconds)\nwith a human readable interface."

	// Print information about the command
	printVersion()

	// TODO: Woah can you actually push a struct on the stack?
	(&HeaderOptions{
		Text:    description,
		Pattern: "-",
		Padding: 1,
	}).PrintHeader()

	// Print the help screen
	help()

	// End the pretty output
	HorizontalRule(description, "-")
}

// Where the magic happens
func main() {
	rawCommand := os.Args[1:]

	// Default to now or expect now
	if len(rawCommand) == 0 || rawCommand[0] == Now {
		fmt.Println(GetNow())
	} else {
		// Evaluate a command given to the program
		result, err := EvaluateCommand(rawCommand)
		if err != nil {
			fmt.Printf("Invalid command: %v", rawCommand)
			os.Exit(InvalidCommand)
		}

		fmt.Println(result)
	}
}
