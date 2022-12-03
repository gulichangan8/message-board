package model

// BtNode 这只是一个评论树
type BtNode struct {
	Id       int
	ParentId int
	Author   string
	Writer   string
	Comment  string
}

type BiTree struct {
	Root     *[]Com
	Children *[]Com
}

type Com struct {
	Id       int
	Author   string
	Writer   string
	Comments string
}
