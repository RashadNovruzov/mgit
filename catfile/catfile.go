package catfile

import (
	"bytes"
	"fmt"
	"log"
	"mgit/utils"
	"os"
)

func CatFile(filepath string, typeOfObj string) {
	obj, err := os.ReadFile(".mgit/objects/" + filepath)
	utils.CheckErr(err)
	parts := bytes.SplitN(obj, []byte{0x00}, 2)
	if len(parts) != 2 {
		log.Fatal("invalid object format: no null byte")
	}

	type_ := string(parts[0]) // decode to string
	content := parts[1]
	if typeOfObj != "" && typeOfObj != type_ {
		log.Fatal("invalid object type: " + type_)
	}
	fmt.Println(string(content))

}
