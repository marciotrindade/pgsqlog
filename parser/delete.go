package parser

import (
	"regexp"
)

var deleteRegexp = regexp.MustCompile(`delete from "?([[a-z0-9_\.]*)"?(.*)`)

// TableOfDelete extract table name from a query
// It's retuns the table name just if it's match with a delete command
func TableOfDelete(query string) string {
	matches := deleteRegexp.FindStringSubmatch(query)

	if len(matches) > 0 {
		return matches[1]
	}
	return ""
}
