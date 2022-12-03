package model

type Use struct {
	User struct {
		Username string
		Password string
	}
	Ques struct {
		UserName string
		TrueName string
		LikeFood string
		Age      int
	}
	Mess struct {
		OwnerName string
		WriteName string
		Message   string
		Respond   string
	}
	Comment struct {
		Author  string
		Writer  string
		Comment string
	}
	Message struct {
		Username      string
		Age           int
		Birthday      float64
		Constellation string
		Sex           string
	}
	Praise struct {
		Author string
		Reader string
		Good   bool
	}
	Good struct {
		Author string
		Member int
	}
	Login struct {
		Username string
		Login    bool
	}
}
