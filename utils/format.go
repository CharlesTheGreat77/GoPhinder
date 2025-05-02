package utils

import (
	"regexp"
	"strings"
)

func StripANSI(input string) string {
	ansiEscape := "\033\\[[0-9;]*m"
	return strings.TrimSpace(
		regexp.MustCompile(ansiEscape).ReplaceAllString(input, ""),
	)
}
