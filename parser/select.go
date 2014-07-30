package parser

import (
	"regexp"
)

var selectRegexp = regexp.MustCompile(`select(.*?)from "?([a-z0-9_\.]*)"?(.*)`)

// TableOfSelect extract table name from a query
// It's retuns the table name just if it's match with a select command
func TableOfSelect(query string) string {
	matches := selectRegexp.FindStringSubmatch(query)

	if len(matches) > 0 {
		return matches[2]
	}
	return ""
}
