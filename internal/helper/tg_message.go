package helper

import "strings"

func Command(message string) string {
	if strings.HasPrefix(message, "/") {
		return strings.Fields(message)[0][1:]
	}

	return ""
}

func CommandArguments(message string) []string {
	if strings.HasPrefix(message, "/") {
		return strings.Fields(message)[1:]
	}

	return nil
}
