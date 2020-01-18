package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

// GetTimeStamp gets a timestamp from an total and an amount
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
		os.Exit(-1)
	}

	return int64(totalInt) * amountInMillis
}

// HandleCommand parses a time expression
func HandleCommand(expression string, when string) {
	total, amount, err := ParseExpression(expression)

	// Handle invalid expressions
	if err != nil {
		fmt.Printf("Invalid expression: %v\n", expression)
		os.Exit(-1)
	}

	// Get the amount of time in milliseconds from the command
	result := GetTimeStamp(total, amount)

	// If it's in the past subtract, in the future add
	if when == Ago {
		fmt.Println(GetNow() - result)
	} else {
		fmt.Println(GetNow() + result)
	}
}

// ParseCommand parses the raw command given to the CLI
func ParseCommand(rawCommand []string) {
	length := len(rawCommand)

	var expression []string

	for i := 0; i < length; i++ {
		part := rawCommand[i]

		if part == "-h" || part == "--help" {
			usage()
			os.Exit(0)
		}

		if SliceContains(part, []string{Ago, From}) != -1 {
			HandleCommand(strings.Join(expression, " "), part)
			break
		} else {
			expression = append(expression, part)
		}
	}
}

func usage() {
	fmt.Println("github.com/SeedyROM/timestamp v1.0.0")
	fmt.Println("Usage:")
	fmt.Println("\ttimestamp 5 minutes from now")
	fmt.Println("\ttimestamp 10 hours ago")
}

func main() {
	rawCommand := os.Args[1:]

	// Default to now or expect now
	if len(rawCommand) == 0 || rawCommand[0] == Now {
		fmt.Println(GetNow())
	} else {
		ParseCommand(rawCommand)
	}
}
