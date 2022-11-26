package model

// User 数据库相关结构体
type User struct {
	Username string
	Password string
}

type Mess struct {
	OwnerName string
	WriteName string
	Message   string
	Respond   string
}

type Ques struct {
	UserName string
	TrueName string
	LikeFood string
	Age      int
}

type Use struct {
	User struct {
		Username string
		Password string
	}
	Mess struct {
		OwnerName string
		WriteName string
		Message   string
		Respond   string
	}
	Ques struct {
		UserName string
		TrueName string
		LikeFood string
		Age      int
	}
}

type Users []User
type Messes []Mess
type Queses []Ques
