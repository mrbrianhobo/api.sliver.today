package common

import (
	"regexp"
)

func TrimStr(in string) string {
	space := regexp.MustCompile(`\s+`)
	out := space.ReplaceAllString(in, " ")
	return out
}
