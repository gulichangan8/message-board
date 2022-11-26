package model

// User 数据库相关结构体
type User struct {
	Username string
	Password string
}

type Ques struct {
	UserName string
	TrueName string
	LikeFood string
	Age      int
}

type Users []User
type Queses []Ques
