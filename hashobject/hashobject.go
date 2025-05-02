package hashobject

import (
	"crypto/sha1"
	"encoding/hex"
	"mgit/utils"
	"os"
)

func HashObject(filePath string, typeOfObj string) {
	if typeOfObj == "" {
		typeOfObj = "blob"
	}
	data := readFile(filePath)
	obj := append([]byte(typeOfObj), 0x00)
	obj = append(obj, data...)
	hash := sha1.New()
	hash.Write(obj)
	hashedString := hex.EncodeToString(hash.Sum(nil))
	createFile(hashedString, obj)
}

func createFile(hashedString string, data []byte) {
	utils.CheckErr(os.WriteFile(".mgit/objects/"+hashedString, data, 0644))
}

func readFile(filePath string) []byte {
	file, err := os.ReadFile(filePath)
	utils.CheckErr(err)
	return file
}
