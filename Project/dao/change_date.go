package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// ChangePassword 修改密码
func ChangePassword(username string, password string) {
	var dns = "root:040818@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := sql.Open("mysql", dns)
	_, err := db.Exec("update user set password=? where username=?", password, username)
	if err != nil {
		return
	}
}
