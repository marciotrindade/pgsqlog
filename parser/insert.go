package parser

import (
	"regexp"
)

var insertRegexp = regexp.MustCompile(`insert into "?([a-z0-9_\.]*)"?(.*)`)

// TableOfInsert extract table name from a query
// It's retuns the table name just if it's match with an insert command
func TableOfInsert(query string) string {
	matches := insertRegexp.FindStringSubmatch(query)

	if len(matches) > 0 {
		return matches[1]
	}
	return ""
}
