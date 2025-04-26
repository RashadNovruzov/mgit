package main

import (
	"log"
	"mgit/cmd"
)

func main() {
	rootCmd := cmd.NewRootCommand()

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
