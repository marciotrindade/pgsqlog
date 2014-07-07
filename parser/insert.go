package parser

import (
	"regexp"
)

var insertRegexp = regexp.MustCompile(`insert into "?([a-z0-9_\.]*)"?(.*)`)

func TableOfInsert(query string) string {
	matches := insertRegexp.FindStringSubmatch(query)

	if len(matches) > 0 {
		return matches[1]
	}
	return ""
}
