package parser

import (
	"regexp"
)

var dropRegexp = regexp.MustCompile(`drop (table|schema)( if exists)? "?([[a-z0-9_\.]*)"?(.*)`)

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
