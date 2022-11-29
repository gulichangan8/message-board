package tool

import (
	"Project/model"
	"database/sql"
)

func Tree(id []int) ([]int, model.BiTree) {
	var dns = "root:040818@tcp(127.0.0.1:3306)/message_board?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := sql.Open("mysql", dns)
	var bt model.BtNode
	var bi model.BiTree
	var com model.Com
	var comes []model.Com
	ok := true
	var ID []int
	for _, Id := range id {
		rows2, _ := db.Query("select * from comments where parent_id=?", Id)
		for rows2.Next() {
			err := rows2.Scan(&bt.Id, &bt.ParentId, &bt.Author, &bt.Writer, &bt.Comment)
			if err != nil {
				ok = false
			}
			ID = append(ID, bt.Id)
			com.Id = bt.Id
			com.Author = bt.Author
			com.Writer = bt.Writer
			com.Comments = bt.Comment
			comes = append(comes, com)
			bi.Root = &comes
		}
	}
	if ok {
		ID, bi = Tree(ID)
	}
	return ID, bi
}
