package utils

import "os"

func CreateFile(path string, data []byte) {
	CheckErr(os.WriteFile(path, data, 0644))
}

func ReadFile(filePath string) []byte {
	file, err := os.ReadFile(filePath)
	CheckErr(err)
	return file
}
