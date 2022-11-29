package dao

import (
	"Project/model"
	"Project/tool"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// TakeDate 取出数据库中form表中的数据
func TakeDate(form string) (date interface{}) {
	var u model.User
	var q model.Ques
	var U model.Users
	var Q model.Queses
	var m model.Mess
	var M model.Messes
	var c model.Comment
	var C model.Comments
	var p model.Message
	var P model.Messages
	var g model.Praise
	var G model.Praises
	var o model.Good
	var O model.Goods
	var l model.Login
	var L model.Logins
	var dns = "root:040818@tcp(127.0.0.1:3306)/message_board?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := sql.Open("mysql", dns)
	rows, _ := db.Query("select * from ?", form)
	//表分类
	switch form {
	case "user":
		for rows.Next() {
			err := rows.Scan(&u.Username, &u.Password)
			if err != nil {
				return
			}
			U = append(U, u)
			date = U
		}
	case "question":
		for rows.Next() {
			err := rows.Scan(&q.UserName, &q.TrueName, &q.LikeFood, &q.Age)
			if err != nil {
				return
			}
			Q = append(Q, q)
			date = Q
		}
	case "message":
		for rows.Next() {
			err := rows.Scan(&m.OwnerName, &m.WriteName, &m.Message, &m.Respond)
			if err != nil {
				return
			}
			M = append(M, m)
			date = M
		}
	case "comments":
		for rows.Next() {
			err := rows.Scan(&c.Author, &c.Writer, &c.Comment)
			if err != nil {
				return
			}
			C = append(C, c)
			date = C
		}
	case "person":
		for rows.Next() {
			err := rows.Scan(&p.Username, &p.Age, &p.Birthday, &p.Constellation, &p.Sex)
			if err != nil {
				return
			}
			P = append(P, p)
			date = P
		}
	case "praise":
		for rows.Next() {
			err := rows.Scan(&g.Author, &g.Reader, &g.Good)
			if err != nil {
				return
			}
			G = append(G, g)
			date = G
		}
	case "good-member":
		for rows.Next() {
			err := rows.Scan(&o.Author, &o.Member)
			if err != nil {
				return
			}
			O = append(O, o)
			date = O
		}
	case "login":
		for rows.Next() {
			err := rows.Scan(&l.Username, &l.Login)
			if err != nil {
				return
			}
			L = append(L, l)
			date = L
		}
	default:
		break
	}
	return date
}

// BringDate 将数据存入数据库
func BringDate(form string, U model.Use) {
	var dns = "root:040818@tcp(127.0.0.1:3306)/message_board?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := sql.Open("mysql", dns)
	u := U.User
	q := U.Ques
	m := U.Mess
	c := U.Comment
	p := U.Message
	g := U.Praise
	o := U.Good
	l := U.Login
	//表分类
	switch form {
	case "user":
		_, err := db.Exec("insert into ? (username,password) value (?,?)",
			form, u.Username, u.Password)
		if err != nil {
			return
		}
	case "question":
		_, err := db.Exec("insert into ? (username,true_name,like_food,age) value (?,?,?,?)",
			form, q.UserName, q.TrueName, q.LikeFood, q.Age)
		if err != nil {
			return
		}
	case "message":
		_, err := db.Exec("insert into ? (owner_name,write_name,message,respond) value (?,?,?,?)",
			form, m.OwnerName, m.WriteName, m.Message, m.Respond)
		if err != nil {
			return
		}
	case "comments":
		_, err := db.Exec("insert into ? (author,writer,comment) value (?,?,?)",
			form, c.Author, c.Writer, c.Comment)
		if err != nil {
			return
		}
	case "person":
		_, err := db.Exec("insert into ? (username,age,birthday,constellation,sex) value (?,?,?,?,?)",
			form, p.Username, p.Age, p.Birthday, p.Constellation, p.Sex)
		if err != nil {
			return
		}
	case "praise":
		_, err := db.Exec("insert into ? (author,reader,good) value (?,?,?)",
			form, g.Author, g.Reader, g.Good)
		if err != nil {
			return
		}
	case "good-member":
		_, err := db.Exec("insert into ? (author,member) value (?,?)",
			form, o.Author, o.Member)
		if err != nil {
			return
		}
	case "login":
		_, err := db.Exec("insert into ? (username,login) value (?,?)",
			form, l.Username, l.Login)
		if err != nil {
			return
		}
	default:
		break
	}
}

// TakeOutAuthorComment 将comments表中对应作者的评论数据按照级别依次取出
func TakeOutAuthorComment(author string) model.BiTree {
	var dns = "root:040818@tcp(127.0.0.1:3306)/message_board?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := sql.Open("mysql", dns)
	rows1, _ := db.Query("select * from comments where author=?", author)
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

// BringComments 将数据存入comments表，id自动加一
func BringComments(parentId int, author string, writer string, comment string) {
	var dns = "root:040818@tcp(127.0.0.1:3306)/message_board?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := sql.Open("mysql", dns)
	_, err := db.Exec("insert into comments (parent_id,author,writer,comment) value (?,?,?,?)",
		parentId, author, writer, comment)
	if err != nil {
		return
	}
}

func TakeOutId(writer string) []model.Com {
	var dns = "root:040818@tcp(127.0.0.1:3306)/message_board?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := sql.Open("mysql", dns)
	rows, _ := db.Query("select * from comments where writer=?", writer)
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
