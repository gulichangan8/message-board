package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var dB *sql.DB

func InitMySql() {
	db, err := sql.Open("mysql",
		"root:040818@tcp(127.0.0.1:3306)/message_board?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	dB = db
}
