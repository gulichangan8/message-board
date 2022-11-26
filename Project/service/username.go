package service

import (
	"Project/dao"
	"Project/model"
	"strings"
)

// IsRepeatUsername 检查用户名是否满足要求
// 参数：想要检查的用户名    返回值：bool ”用户名已存在“或者“设置成功”
func IsRepeatUsername(username string) bool {
	U := dao.TakeDate("user", "user")
	u, _ := U.(model.Users)
	var usernames []string
	ok := false
	for _, date := range u {
		usernames = append(usernames, date.Username)
	}
	for _, value := range usernames {
		if value == username {
			ok = true
		} else {
			continue
		}
	}
	if ok {
		return ok
	} else {
		return !ok
	}
}

func IsShortName(username string) bool {
	var name = []rune(username)
	if len(name) > 6 {
		return true
	} else {
		return false
	}
}

func IsNameErr(username string) bool {
	if strings.Contains(username, "\\") == true {
		return false
	} else {
		return true
	}
}
