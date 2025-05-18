package commit

import (
	"bytes"
	"errors"
	"mgit/constants"
	"mgit/models"
	"mgit/utils"
	"mgit/writetree"
	"os"
	"strings"
)

func Commit(message string) string {
	commit := "tree " + writetree.WriteTree("") + "\n"

	head := GetHead()
	if head != "" {
		commit = commit + "parent " + head + "\n"
	}
	commit = commit + "\n" + message
	oid, commitBytes := utils.HashData([]byte(commit), "commit")
	utils.CreateFile(constants.DefaultMgitObjectsPath+oid, commitBytes)
	SetHead(oid)
	return oid
}

func SetHead(oid string) {
	utils.CreateFile(constants.DefaultMgitObjectsPath+"HEAD", []byte(oid))
}

func GetHead() string {
	_, err := os.Stat(constants.DefaultMgitObjectsPath + "HEAD")
	if err != nil {
		return ""
	}
	return string(utils.ReadFile(constants.DefaultMgitObjectsPath + "HEAD"))
}

func GetCommit(oid string) (*models.Commit, error) {

	commitBytes := utils.ReadFile(constants.DefaultMgitObjectsPath + oid)
	commitString := string(bytes.SplitN(commitBytes, []byte{0x00}, 2)[1])

	var commit models.Commit
	commit.ParentOid = ""

	lines := strings.Split(commitString, "\n")
	i := 0

	for ; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			i++ // Skip the empty line
			break
		}

		parts := strings.SplitN(line, " ", 2)

		key, value := parts[0], parts[1]
		switch key {
		case "tree":
			commit.TreeOid = value
		case "parent":
			commit.ParentOid = value
		default:
			return &commit, errors.New("unknown field: " + key)
		}
	}

	commit.Message = strings.Join(lines[i:], "\n")

	return &commit, nil

}
