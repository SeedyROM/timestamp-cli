package main

import "time"

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
