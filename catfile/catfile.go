package catfile

import (
	"fmt"
	"mgit/utils"
	"os"
)

func CatFile(filepath string) {
	file, err := os.ReadFile(".mgit/objects/" + filepath)
	utils.CheckErr(err)
	fmt.Println(string(file))
}
