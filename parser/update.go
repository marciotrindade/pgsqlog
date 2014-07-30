package parser

import (
	"regexp"
)

var updateRegexp = regexp.MustCompile(`update "?([a-z0-9_\.]*)"?(.*)`)

// TableOfUpdate extract table name from a query
// It's retuns the table name just if it's match with an update command
func TableOfUpdate(query string) string {
	matches := updateRegexp.FindStringSubmatch(query)

	if len(matches) > 0 {
		return matches[1]
	}
	return ""
}
