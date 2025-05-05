package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func HashData(data []byte, objectType string) (string, []byte) {
	obj := append([]byte(objectType), 0x00)
	obj = append(obj, data...)
	hash := sha1.New()
	hash.Write(obj)
	oid := hex.EncodeToString(hash.Sum(nil))
	return oid, obj
}
