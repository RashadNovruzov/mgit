package log

import (
	"fmt"
	"mgit/commit"
	"mgit/utils"
)

func PrintLog(commitOid string) {
	if commitOid == "" {
		commitOid = commit.GetHead()
	}

	for commitOid != "" {
		commitInfo, err := commit.GetCommit(commitOid)
		utils.CheckErr(err)
		commitOid = commitInfo.ParentOid
		fmt.Println("-----------------------------")
		fmt.Println("Commit message: " + commitInfo.Message)
		fmt.Println("Commit parent: " + commitInfo.ParentOid)
		fmt.Println("Commit tree: " + commitInfo.TreeOid)
		fmt.Println("-----------------------------")
		fmt.Println()
	}
}
