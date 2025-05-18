package writetree

import (
	"fmt"
	"mgit/constants"
	"mgit/hashobject"
	"mgit/models"
	"mgit/utils"
	"os"
	"path/filepath"
	"strings"
)

func WriteTree(folderPath string) string {
	var hashedEntries []models.TreeEntry

	if len(folderPath) == 0 {
		folderPath = constants.DefaultFolderPath
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
		hashedEntries = append(hashedEntries, models.TreeEntry{Name: entry.Name(), Oid: oid, ObjectType: objectType})
	}

	var builder strings.Builder
	for _, entry := range hashedEntries {
		builder.WriteString(fmt.Sprintf("%s %s %s\n", entry.ObjectType, entry.Oid, entry.Name))
	}
	tree := builder.String()
	oidTree, data := utils.HashData([]byte(tree), "tree")
	utils.CreateFile(constants.DefaultMgitObjectsPath+oidTree, data)

	return oidTree
}

func isIgnored(path string) bool {
	parts := strings.Split(filepath.ToSlash(path), "/")
	for _, part := range parts {
		for _, ignored := range constants.IgnoredPaths {
			if strings.HasPrefix(part, ignored) {
				return true
			}
		}
	}
	return false
}
