package utils

import (
	"fmt"
	"log"
	"mgit/constants"
	"strings"
)

func IsIgnoredPath(path string) bool {
	fmt.Println(path)

	parts := strings.Split(path, "/")
	log.Println(parts)
	for _, ignoredPath := range constants.IgnoredPaths {
		for _, part := range parts {
			if strings.HasPrefix(part, ignoredPath) {
				return true
			}
		}
	}
	return false
}
