package main

import (
	"mgit/cmd"
	"mgit/utils"
)

func main() {
	rootCmd := cmd.NewRootCommand()

	err := rootCmd.Execute()
	utils.CheckErr(err)
}
