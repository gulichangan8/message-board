package dao

import (
	"Project/model"
	"database/sql"
)

// TakeMessageDate 将message表中的数据取出
func TakeMessageDate() model.Messes {
	rows, _ := dB.Query("select * from ?", "message")
	var m model.Mess
	var M model.Messes
	for rows.Next() {
		err := rows.Scan(&m.OwnerName, &m.WriteName, &m.Message, &m.Respond)
		if err != nil {
			panic(err)
		}
		M = append(M, m)
	}
	return M
}

// BringMessageDate 将数据存入message表
func BringMessageDate(m model.Mess) {
	_, err := dB.Exec("insert into message (owner_name, message, write_name, respond) value (?,?,?,?)",
		m.OwnerName, m.WriteName, m.Message, m.Respond)
	if err != nil {
		return
	}
}

// BringCommentDate 将数据存入comment表
func BringCommentDate(c model.Comment) {
	_, err := dB.Exec("insert into comments (author,writer,comment) value (?,?,?)",
		c.Author, c.Writer, c.Comment)
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

// RespondMessageDate 回复留言
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
