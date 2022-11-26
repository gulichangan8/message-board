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
	_, err := db.Exec("update message set message=? where owner_name=?", message, username)
	if err != nil {
		return
	}
}
