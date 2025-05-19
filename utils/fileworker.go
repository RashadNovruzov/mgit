package utils

import "os"

func CreateFile(path string, data []byte) {
	CheckErr(os.WriteFile(path, data, 0644))
}

func ReadFile(filePath string) []byte {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil
	}
	file, err := os.ReadFile(filePath)
	CheckErr(err)
	return file
}
