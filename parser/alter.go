package parser

import (
	"regexp"
)

var alterRegexp = regexp.MustCompile(`alter table "?([[a-z0-9_\.]*)"?(.*)`)

// TableOfAlter extract table name from a query
// It's retuns the table name just if it's match with an alter command
func TableOfAlter(query string) string {
	matches := alterRegexp.FindStringSubmatch(query)

	if len(matches) > 0 {
		return matches[1]
	}
	return ""
}
