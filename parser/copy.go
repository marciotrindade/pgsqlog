package parser

import (
	"regexp"
)

var copyRegexp = regexp.MustCompile(`copy "?([a-z0-9_\.]*)"?(.*)`)

// TableOfCopy extract table name from a query
// It's retuns the table name just if it's match with a copy command
func TableOfCopy(query string) string {
	matches := copyRegexp.FindStringSubmatch(query)

	if len(matches) > 0 {
		return matches[1]
	}
	return ""
}
