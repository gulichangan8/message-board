package model

type Comment struct {
	Author  string
	Writer  string
	Comment string
}

type Praise struct {
	Author string
	Reader string
	Good   bool
}

type Good struct {
	Author string
	Member int
}

type Comments []Comment
type Praises []Praise
type Goods []Good
