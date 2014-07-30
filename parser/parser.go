package parser

import (
	"regexp"
	"strings"
)

var lineRegexp = regexp.MustCompile(`([0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}) brt \[[0-9]*\]: \[(.*?)\] user=(.*?),db=(.*?) log:  duration: (.*?) ms  (.*?): (.*)`)
var sqlRegexp = regexp.MustCompile(`([a-z]*)(.*)`)

type LogLine struct {
	Action    string
	CreatedAt string
	Database  string
	Duration  string
	Line      string
	Query     string
	Table     string
	Username  string
}

func (parse *LogLine) Parse() {
	matches := lineRegexp.FindStringSubmatch(parse.Line)

	if len(matches) > 0 {
		parse.Username = matches[3]
		parse.Database = matches[4]
		parse.Duration = matches[5]
		parse.Query = strings.Trim(matches[7], " ")
		parse.CreatedAt = matches[1]
		parse.Sql()
	}
}

func (parse *LogLine) Sql() {
	matches := sqlRegexp.FindStringSubmatch(parse.Query)

	parse.Action = matches[1]

	switch matches[1] {
	case "select":
		parse.Table = TableOfSelect(parse.Query)
	case "update":
		parse.Table = TableOfUpdate(parse.Query)
	case "insert":
		parse.Table = TableOfInsert(parse.Query)
	case "copy":
		parse.Table = TableOfCopy(parse.Query)
	case "delete":
		parse.Table = TableOfDelete(parse.Query)
	case "create":
		parse.Action, parse.Table = TableOfCreate(parse.Query)
	case "drop":
		parse.Action, parse.Table = TableOfDrop(parse.Query)
	case "alter":
		parse.Table = TableOfAlter(parse.Query)
	}
}
