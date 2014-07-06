package parse

import (
	"regexp"
)

var alterRegexp = regexp.MustCompile(`alter table "?([[a-z0-9_\.]*)"?(.*)`)

func TableOfAlter(query string) string {
	matches := selectRegexp.FindStringSubmatch(query)

	if len(matches) > 0 {
		return matches[1]
	}
	return ""
}