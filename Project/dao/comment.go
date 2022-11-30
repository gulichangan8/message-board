package dao

import (
	"Project/model"
	"Project/tool"
	"database/sql"
)

// TakeCommentDate 将comment表中的数据取出
func TakeCommentDate() model.Comments {
	rows, _ := dB.Query("select * from ?", "user")
	var c model.Comment
	var C model.Comments
	for rows.Next() {
		err := rows.Scan(&c.Author, &c.Writer, &c.Comment)
		if err != nil {
			panic(err)
		}
		C = append(C, c)
	}
	return C
}

// TakeOutAuthorComment 将comments表中对应作者的评论数据按照级别依次取出
func TakeOutAuthorComment(author string) model.BiTree {
	rows1, _ := dB.Query("select * from comments where author=?", author)
	var id []int
	var bt model.BtNode
	var bi model.BiTree
	var com model.Com
	var comes []model.Com
	for rows1.Next() {
		err := rows1.Scan(&bt.Id, &bt.ParentId, &bt.Author, &bt.Writer, &bt.Comment)
		if err != nil {
			break
		}
		id = append(id, bt.Id)
		com.Id = bt.Id
		com.Author = bt.Author
		com.Writer = bt.Writer
		com.Comments = bt.Comment
		comes = append(comes, com)
		bi.Root = &comes
	}
	_, bi = tool.Tree(id)
	return bi
}

// TakeOutId 取出自己评论中的id
func TakeOutId(writer string) []model.Com {
	rows, _ := dB.Query("select * from comments where writer=?", writer)
	var bt model.BtNode
	var com model.Com
	var comes []model.Com
	for rows.Next() {
		err := rows.Scan(&bt.Id, &bt.ParentId, &bt.Author, &bt.Writer, &bt.Comment)
		if err != nil {
			break
		}
		com.Id = bt.Id
		com.Author = bt.Author
		com.Writer = bt.Writer
		com.Comments = bt.Comment
		comes = append(comes, com)
	}
	return comes
}

// BringComments 将数据存入comments表，id自动加一
func BringComments(parentId int, author string, writer string, comment string) {
	_, err := dB.Exec("insert into comments (parent_id,author,writer,comment) value (?,?,?,?)",
		parentId, author, writer, comment)
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

// ChangeComments 修改嵌套评论
func ChangeComments(id int, com string) {
	var dns = "root:040818@tcp(127.0.0.1:3306)/message_board?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := sql.Open("mysql", dns)
	_, err := db.Exec("update comments set comment=? where id=?", com, id)
	if err != nil {
		return
	}
}
