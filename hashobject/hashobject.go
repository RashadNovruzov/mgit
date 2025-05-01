package hashobject

import (
	"crypto/sha1"
	"encoding/hex"
	"mgit/utils"
	"os"
)

func HashObject(filePath string) {
	data := readFile(filePath)
	hash := sha1.New()
	hash.Write(data)
	hashedString := hex.EncodeToString(hash.Sum(nil))
	createFile(hashedString, data)
}

func createFile(hashedString string, data []byte) {
	utils.CheckErr(os.WriteFile(".mgit/objects/"+hashedString, data, 0644))
}

func readFile(filePath string) []byte {
	file, err := os.ReadFile(filePath)
	utils.CheckErr(err)
	return file
}
