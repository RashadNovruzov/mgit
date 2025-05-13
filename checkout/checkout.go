package checkout

import (
	"mgit/commit"
	"mgit/readtree"
)

func Checkout(commitId string) {
	commitEntry, _ := commit.GetCommit(commitId)
	treeOid := commitEntry.TreeOid
	readtree.ReadTree(treeOid)
	commit.SetHead(commitId)
}
