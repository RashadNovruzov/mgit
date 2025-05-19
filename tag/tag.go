package tag

import (
	"mgit/constants"
	"mgit/utils"
	"os"
	"path/filepath"
)

func CreateTag(name string, oid string) {
	UpdateRef(constants.DefaultRefsFolder+name, oid)
}

func UpdateRef(ref string, oid string) {
	refPath := constants.DefaultMgitPath + ref
	refPathDirs := filepath.Dir(refPath)
	err := os.MkdirAll(refPathDirs, os.ModePerm)
	utils.CheckErr(err)
	utils.CreateFile(refPath, []byte(oid))
}

func GetRef(ref string) string {
	refPath := constants.DefaultMgitPath + constants.DefaultRefsFolder + ref
	fileData := utils.ReadFile(refPath)
	if fileData == nil {
		return ""
	}
	return string(fileData)
}
