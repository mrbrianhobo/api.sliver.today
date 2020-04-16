package common

import (
	"regexp"
	"strings"
)

func TrimStr(in string) string {
	space := regexp.MustCompile(`\s+`)
	out := space.ReplaceAllString(in, " ")
	return out
}

func MatchKeywords(in string, keywords []string) bool {
	for _, word := range keywords {
		if strings.EqualFold(in, word) {
			return true
		}
	}

	return false
}
