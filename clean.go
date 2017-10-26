package cleanurl

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

func Clean(raw string) string {
	parsed, _ := url.Parse(raw)
	return strings.TrimSpace(fmt.Sprintf("%s%s%s",
		CleanHost(parsed.Host),
		CleanPath(parsed.Path),
		CleanQuery(parsed.RawQuery)))
}

func CleanHost(host string) string {
	return regexp.MustCompile(`^www\.`).ReplaceAllString(host, "")
}

func CleanPath(path string) string {
	if len(path) == 0 {
		return ""
	}

	if path[len(path)-1:] == "/" {
		path = path[:len(path)-1]
	}

	return path
}

func CleanQuery(query string) string {
	if len(query) == 0 {
		return ""
	}

	result := regexp.MustCompile(`(\&|^)utm_[\w]+\=[^\&]+`).ReplaceAllString(query, "")
	result = regexp.MustCompile(`(\&|^)ref\=[^\&]+`).ReplaceAllString(result, "")

	if string(result[0]) == "&" {
		result = result[1:]
	}

	return "?" + result
}
