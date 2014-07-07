package parser

import (
	"regexp"
)

var updateRegexp = regexp.MustCompile(`update "?([a-z0-9_\.]*)"?(.*)`)

func TableOfUpdate(query string) string {
	matches := updateRegexp.FindStringSubmatch(query)

	if len(matches) > 0 {
		return matches[1]
	}
	return ""
}
