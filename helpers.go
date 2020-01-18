package main

import (
	"fmt"
	"strings"
	"time"
)

// SliceContains find a value in a slice and returns the index and bool
func SliceContains(value string, slice []string) int {
	for i, item := range slice {
		if item == value {
			return i
		}
	}

	return -1
}

// GetNow gets the time right now in milliseconds
func GetNow() int64 {
	return time.Now().Round(time.Millisecond).UnixNano() / 1e6
}

// HeaderOptions stores a pretty headers configuration
type HeaderOptions struct {
	Text    string
	Pattern string
	Padding int
}

// HorizontalRule finds the max length of each line of a string
// and returns back a string of pattern length
func HorizontalRule(text string, pattern string) {
	lines := strings.Split(text, "\n")

	var max int = 0
	for _, value := range lines {
		if len(value) > max {
			max = len(value)
		}
	}

	fmt.Println(strings.Repeat(pattern, max))
}

// PrintHeader prints a pretty printed header
func (opts *HeaderOptions) PrintHeader() {
	if opts.Text == "" {
		HorizontalRule(opts.Text, opts.Pattern)

		return
	}
	HorizontalRule(opts.Text, opts.Pattern)
	fmt.Println(opts.Text)
	HorizontalRule(opts.Text, opts.Pattern)
}
