package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	re, err := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)

	if err != nil {
		panic(err)
	}

	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re, err := regexp.Compile(`<[~=*-]*>`)

	if err != nil {
		panic(err)
	}
	return re.Split(text, -1)

}

func CountQuotedPasswords(lines []string) int {
	re, err := regexp.Compile(`(?i)".*password.*"`)

	if err != nil {
		panic(err)
	}
	counter := 0
	for _, line := range lines {
		if re.MatchString(line) {
			counter++
		}
	}
	return counter

}

func RemoveEndOfLineText(text string) string {
	re, err := regexp.Compile(`end-of-line\d+`)

	if err != nil {
		panic(err)
	}
	return re.ReplaceAllString(text, "")

}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`\bUser\s+([a-zA-Z0-9]{6,})\b`)

	var result []string
	for _, line := range lines {
		// Check if the line contains a user name
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			// Prepend the tag "[USR] <UserName>" to the line
			userName := matches[1]
			line = "[USR] " + userName + " " + line
		}
		result = append(result, line)
	}
	return result
}
