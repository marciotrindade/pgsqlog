package parse

import (
	"regexp"
)

var selectRegexp = regexp.MustCompile(`select(.*?)from "?([a-z0-9_\.]*)"?(.*)`)

func TableOfSelect(query string) string {
	matches := selectRegexp.FindStringSubmatch(query)

	if len(matches) > 0 {
		return matches[2]
	}
	return ""
}
