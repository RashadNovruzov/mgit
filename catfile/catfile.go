package catfile

import (
	"bytes"
	"log"
	"mgit/constants"
	"mgit/utils"
)

func CatFile(oid string, typeOfObj string) string {
	obj := utils.ReadFile(constants.DefaultMgitObjectsPath + oid)
	parts := bytes.SplitN(obj, []byte{0x00}, 2)
	if len(parts) != 2 {
		log.Fatal("invalid object format: no null byte")
	}

	type_ := string(parts[0]) // decode to string
	content := parts[1]
	if typeOfObj != "" && typeOfObj != type_ {
		log.Fatal("Expected " + type_ + ", got: " + typeOfObj)
	}
	sContent := string(content)
	return sContent
}
