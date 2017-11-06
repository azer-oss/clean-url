package cleanurl

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

func Clean(raw string) string {
	parsed, err := url.Parse(raw)
	if err != nil {
		return raw
	}

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

	values, err := url.ParseQuery(query)
	if err != nil {
		return ""
	}

	result := []string{}
	for k, v := range values {
		if k != "ref" && !strings.HasPrefix(k, "utm_") && len(v) > 0 && len(v[0]) > 0 {
			result = append(result, fmt.Sprintf("%s=%s", k, v[0]))
		}
	}

	if len(result) == 0 {
		return ""
	}

	return fmt.Sprintf("?%s", strings.Join(result, "&"))
}
