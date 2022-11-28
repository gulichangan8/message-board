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

type Message struct {
	Username      string
	Age           int
	Birthday      float64
	Constellation string
	Sex           string
}

type Login struct {
	Username string
	Login    bool
}

type Users []User
type Queses []Ques
type Messages []Message
type Logins []Login
