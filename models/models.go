package models

type TreeEntry struct {
	Name       string
	Oid        string
	ObjectType string
}

type Commit struct {
	Message   string
	TreeOid   string
	ParentOid string
}
