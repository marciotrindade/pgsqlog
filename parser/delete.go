package parse

import (
	"regexp"
)

var deleteRegexp = regexp.MustCompile(`delete from "?([[a-z0-9_\.]*)"?(.*)`)

func TableOfDelete(query string) string {
	matches := selectRegexp.FindStringSubmatch(query)

	if len(matches) > 0 {
		return matches[1]
	}
	return ""
}
