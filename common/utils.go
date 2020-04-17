package common

import (
	"regexp"
	"strings"
	"time"
)

var (
	Timezone, _  = time.LoadLocation("America/Los_Angeles")
	datetimeFmt  = "2006-01-02"
	regexes      = []string{`^[\s\p{Zs}]+|[\s\p{Zs}]+$`, `[\s\p{Zs}]{2,}`, `-[\s\p{Zs}]+`}
	replacements = []string{"", " ", "-"}
)

func SanitizeStr(in string) string {
	out := in
	for i, regex := range regexes {
		out = regexp.MustCompile(regex).ReplaceAllString(out, replacements[i])
	}
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
