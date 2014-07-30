package parser

import (
	"regexp"
)

var createRegexp = regexp.MustCompile(`create +(table|index|unique index) "?([[a-z0-9_\.]*)"?(.*)`)

// TableOfCreate extract action and table name from a query
// It's retuns the action and table name just if it's match with a create command
// action can be (table, index or unique index)
func TableOfCreate(query string) (string, string) {
	var table string
	var action string

	matches := createRegexp.FindStringSubmatch(query)

	if len(matches) > 0 {
		table = matches[2]

		switch matches[1] {
		case "table":
			action = "create table"
		case "index":
			action = "create index"
		case "unique index":
			action = "create index"
		}
	}
	return action, table
}
