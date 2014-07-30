package parser

import (
	"regexp"
)

var dropRegexp = regexp.MustCompile(`drop (table|schema)( if exists)? "?([[a-z0-9_\.]*)"?(.*)`)

// TableOfDrop extract action and table name from a query
// It's retuns the action and table name just if it's match with a drop command
// action can be (drop table, drop index or drop schema)
func TableOfDrop(query string) (string, string) {
	var table string
	var action string

	matches := dropRegexp.FindStringSubmatch(query)

	if len(matches) > 0 {
		table = matches[3]
		switch matches[1] {
		case "table":
			action = "drop table"
		case "schema":
			action = "drop schema"
		}
	}
	return action, table
}
