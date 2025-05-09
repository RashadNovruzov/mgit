package commit

import (
	"mgit/constants"
	"mgit/utils"
	"mgit/writetree"
)

func Commit(message string) string {
	commit := "tree " + writetree.WriteTree("")
	commit = commit + "\n"
	commit = commit + message
	oid, bytes := utils.HashData([]byte(commit), "commit")
	utils.CreateFile(constants.DefaultMgitPath+oid, bytes)
	return oid
}
