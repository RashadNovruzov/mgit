package hashobject

import (
	"log"
	"mgit/constants"
	"mgit/utils"
)

func HashObject(filePath string, objectType string) string {
	if objectType == "" {
		objectType = "blob"
	}
	data := utils.ReadFile(filePath)
	oid := utils.HashData(data, objectType)
	utils.CreateFile(constants.DefaultMgitPath+oid, data)
	log.Println("File ", filePath, " saved with hash ", oid)
	return oid
}
