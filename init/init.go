package init

import (
	"log"
	"mgit/utils"
	"os"
)

func InitializeRepo() {
	err := os.Mkdir(".mgit", 0755)
	utils.CheckErr(err)
	err = os.Mkdir(".mgit/objects", 0755)
	utils.CheckErr(err)
	log.Println("Initialized git repository")
}
