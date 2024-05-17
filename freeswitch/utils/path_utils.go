package utils

import (
	"path/filepath"
	"regexp"
)

func IsPathValid(path string, extensions ...string) bool {
	if path == "" {
		return false
	}

	ext := filepath.Ext(path)
	path = path[:len(path)-len(ext)]

	pattern := `^(/|./|../)?([\w\s-]+/)*[\w\s-]+$`
	match, _ := regexp.MatchString(pattern, path)

	if !match {
		return false
	}

	if len(extensions) == 0 {
		return true
	}

	for _, extension := range extensions {
		if ext == extension {
			return true
		}
	}

	return false
}
