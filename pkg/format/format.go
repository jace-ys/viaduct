package format

import (
	"strings"
)

func SingleJoiningSlash(target, path string) string {
	t := strings.TrimSuffix(target, "/")
	p := strings.TrimPrefix(path, "/")
	return t + "/" + p
}

func AddSlashes(prefix string) string {
	return "/" + prefix + "/"
}
