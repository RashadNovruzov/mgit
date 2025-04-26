package init

import (
	"log"
	"os"
)

func InitializeRepo() {
	err := os.Mkdir(".mgit", 0755)
	if err != nil {
		log.Fatal("Error while initializing repository", err)
	}
	log.Println("Initialized git repository")
}
