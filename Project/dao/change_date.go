package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// ChangePasswordDate 修改密码
func ChangePasswordDate(username string, password string) {
	var dns = "root:040818@tcp(127.0.0.1:3306)/message_board?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := sql.Open("mysql", dns)
	_, err := db.Exec("update user set password=? where username=?", password, username)
	if err != nil {
		return
	}
}

// ChangeMessageDate 修改留言
func ChangeMessageDate(username string, message string) {
	var dns = "root:040818@tcp(127.0.0.1:3306)/message_board?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := sql.Open("mysql", dns)
	_, err := db.Exec("update message set message=? where write_name=?", message, username)
	if err != nil {
		return
	}
}

func RespondMessageDate(username string, respond string) {
	var dns = "root:040818@tcp(127.0.0.1:3306)/message_board?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := sql.Open("mysql", dns)
	_, err := db.Exec("update message set respond=? where owner_name=?", respond, username)
	if err != nil {
		return
	}
}

// DeleteMessageDate 删除留言
func DeleteMessageDate(username string) {
	var dns = "root:040818@tcp(127.0.0.1:3306)/message_board?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := sql.Open("mysql", dns)
	_, err := db.Exec("update message set message=? where write_name=?", "留言已删除", username)
	if err != nil {
		return
	}
}

// ChangeCommentDate 修改评论
func ChangeCommentDate(comment string, username string) {
	var dns = "root:040818@tcp(127.0.0.1:3306)/message_board?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := sql.Open("mysql", dns)
	_, err := db.Exec("update comment set `comment`=? where writer=?", comment, username)
	if err != nil {
		return
	}
}

func ChangePersonalMessage(username string, age int, birthday float64, constellation string, sex string) {
	var dns = "root:040818@tcp(127.0.0.1:3306)/message_board?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := sql.Open("mysql", dns)
	_, err := db.Exec("update person set age=?,birthday=?,constellation=?,sex=? where username=?", age, birthday, constellation, sex, username)
	if err != nil {
		return
	}
}
