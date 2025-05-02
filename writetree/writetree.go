package writetree

import (
	"fmt"
	"mgit/constants"
	"mgit/hashobject"
	"mgit/utils"
	"os"
	"path/filepath"
	"strings"
)

var IGNORED_PATHS = []string{".git", ".hg", ".svn", ".mgit", ".idea"}

type Entry struct {
	Name       string
	Oid        string
	ObjectType string
}

func WriteTree(folderPath string) string {
	var hashedEntries []Entry

	if len(folderPath) == 0 {
		folderPath = constants.DEFAULT_FOLDER_PATH
	}
	entries, err := os.ReadDir(folderPath)
	utils.CheckErr(err)

	var objectType, oid string

	for _, entry := range entries {
		fullPath := filepath.Join(folderPath, entry.Name())
		if isIgnored(fullPath) {
			continue
		}
		if entry.IsDir() {
			objectType = "tree"
			oid = WriteTree(fullPath)
		} else {
			objectType = "blob"
			oid = hashobject.HashObject(fullPath, objectType)
		}
		hashedEntries = append(hashedEntries, Entry{entry.Name(), oid, objectType})
	}

	var builder strings.Builder
	for _, entry := range hashedEntries {
		builder.WriteString(fmt.Sprintf("%s %s %s\n", entry.ObjectType, entry.Oid, entry.Name))
	}
	tree := builder.String()
	oidTree := utils.HashData([]byte(tree), objectType)
	utils.CreateFile(constants.DEFAULT_MGIT_PATH+oidTree, []byte(tree))

	return oidTree
}

func isIgnored(path string) bool {
	parts := strings.Split(filepath.ToSlash(path), "/")
	for _, part := range parts {
		for _, ignored := range IGNORED_PATHS {
			if strings.HasPrefix(part, ignored) {
				return true
			}
		}
	}
	return false
}
