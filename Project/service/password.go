package service

import "strings"

func IsPasswordErr(password string) bool {
	if strings.Contains(password, "\\") == true {
		return false
	} else {
		return true
	}
}

func IsShortPassword(password string) bool {
	var name = []rune(password)
	if len(name) > 6 {
		return true
	} else {
		return false
	}
}
