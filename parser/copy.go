package parse

import (
	"regexp"
)

var copyRegexp = regexp.MustCompile(`copy "?([a-z0-9_\.]*)"?(.*)`)

func TableOfCopy(query string) string {
	matches := selectRegexp.FindStringSubmatch(query)

	if len(matches) > 0 {
		return matches[1]
	}
	return ""
}