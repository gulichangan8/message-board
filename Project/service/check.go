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
	U := dao.TakeDate("user", "user")
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

func PasswordProtect(username string, trueName string, likeFood string, age int) bool {
	M := dao.TakeDate("user", "message")
	m, _ := M.(model.Queses)
	ok := false
	for _, date := range m {
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
