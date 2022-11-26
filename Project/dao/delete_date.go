package dao

import "database/sql"

// DeleteMessageDate 删除留言
func DeleteMessageDate(username string) {
	var dns = "root:040818@tcp(127.0.0.1:3306)/message_board?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := sql.Open("mysql", dns)
	_, err := db.Exec("delete from message where write_name=?", username)
	if err != nil {
		return
	}
}
