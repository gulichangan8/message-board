package service

import (
	"Project/dao"
	"Project/model"
)

// CheckUsername 检查用户名格式
func CheckUsername(username string) bool {
	ok1 := IsShortName(username)
	ok2 := IsNameErr(username)
	ok3 := IsRepeatUsername(username)
	if ok1 && ok2 {
		if ok3 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

// CheckPassword 检查密码格式
func CheckPassword(password string) bool {
	ok1 := IsPasswordErr(password)
	ok2 := IsShortPassword(password)
	if ok1 && ok2 {
		return true
	} else {
		return false
	}
}

// CheckLogin 登录账号检查
func CheckLogin(username string, password string) bool {
	U := dao.TakeDate("user")
	u, _ := U.(model.Users)
	ok := false
	for _, date := range u {
		if date.Username == username && date.Password == password {
			ok = true
		} else {
			continue
		}
	}
	if ok {
		return true
	} else {
		return false
	}
}

// PasswordProtect 密保问题是否回答正确
func PasswordProtect(username string, trueName string, likeFood string, age int) bool {
	Q := dao.TakeDate("question")
	q, _ := Q.(model.Queses)
	ok := false
	for _, date := range q {
		if date.UserName == username && date.TrueName == trueName && date.LikeFood == likeFood && date.Age == age {
			ok = true
		} else {
			continue
		}
	}
	if ok {
		return true
	} else {
		return false
	}
}

// CheckUsernameExist 判断写给的人是否存在
func CheckUsernameExist(username string) bool {
	user := dao.TakeDate("user")
	U, _ := user.(model.Users)
	ok := false
	for _, date := range U {
		if date.Username == username {
			ok = true
		} else {
			continue
		}
	}
	return ok
}

// CheckWriterExist 判断是否留过言
func CheckWriterExist(writer string) bool {
	M := dao.TakeDate("message")
	m, _ := M.(model.Messes)
	ok := false
	for _, date := range m {
		if date.WriteName == writer {
			ok = true
		} else {
			continue
		}
	}
	return ok
}

// IfHaveSomeoneWriteTo 判断是否有人留过言
func IfHaveSomeoneWriteTo(username string) bool {
	M := dao.TakeDate("message")
	m, _ := M.(model.Messes)
	ok := false
	for _, date := range m {
		if date.OwnerName == username {
			ok = true
		} else {
			continue
		}
	}
	return ok
}

func CheckAuthorExist(username string) bool {
	c := dao.TakeDate("comment")
	C, _ := c.(model.Comments)
	ok := false
	for _, date := range C {
		if date.Author == username {
			ok = true
		} else {
			continue
		}
	}
	return ok
}

func CheckWriteExist(username string) bool {
	c := dao.TakeDate("comment")
	C, _ := c.(model.Comments)
	ok := false
	for _, date := range C {
		if date.Writer == username {
			ok = true
		} else {
			continue
		}
	}
	return ok
}

func CheckCommentLength(comment string) bool {
	var name = []rune(comment)
	if len(name) > 100 {
		return true
	} else {
		return false
	}
}
