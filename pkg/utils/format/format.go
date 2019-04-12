package format

import (
	"regexp"
	"strings"
)

var Services = `
			Service: %s
			Prefix: %s
			Upstream Url: %s
			Methods: %v
			Allow Cross Origin: %t
			Middlewares: %s
`

func Methods(m []string) string {
	if len(m) == 0 {
		return "[*]"
	}

	return joinArray(m)
}

func Middleware(m map[string]bool) string {
	var s []string

	if len(m) == 0 {
		return "None"
	}

	for key, value := range m {
		if value {
			s = append(s, strings.Title(key))
		}
	}

	return joinArray(s)
}

func SingleJoiningSlash(target, path string) string {
	t := strings.TrimSuffix(target, "/")
	p := strings.TrimPrefix(path, "/")
	return t + "/" + p
}

func AddSlashes(prefix string) string {
	return "/" + prefix + "/"
}

func CamelToSnakeUnderscore(camelCase string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")

	snakeUnderscore := matchFirstCap.ReplaceAllString(camelCase, "${1}_${2}")
	snakeUnderscore = matchAllCap.ReplaceAllString(snakeUnderscore, "${1}_${2}")
	return strings.ToUpper(snakeUnderscore)
}

func joinArray(s []string) string {
	return "[" + strings.Join(s, ", ") + "]"
}
