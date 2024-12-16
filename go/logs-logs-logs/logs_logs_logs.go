package logs

import (
	"fmt"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	runes := []rune(log)
	for _, x := range runes {
		s := fmt.Sprintf("%U", x)
		if s == "U+2757" {
			return "recommendation"
		} else if s == "U+1F50D" {
			return "search"
		} else if s == "U+2600" {
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	//runes := []rune(log)
	newLog := ""
	for _, x := range log {
		if x == oldRune {
			newLog += fmt.Sprintf("%c", newRune)
		} else {
			newLog += fmt.Sprintf("%c", x)
		}
	}
	return newLog
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	length := utf8.RuneCountInString(log)
	if length <= limit {
		return true
	}
	return false
}
