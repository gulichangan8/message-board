package dao

import (
	"Project/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// TakeDate 取出数据库user中form表中的数据
func TakeDate(form string) (date interface{}) {
	var u model.User
	var q model.Ques
	var U model.Users
	var Q model.Queses
	var m model.Mess
	var M model.Messes
	var dns = "root:040818@tcp(127.0.0.1:3306)/message_board?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := sql.Open("mysql", dns)
	rows, _ := db.Query("select * from ?", form)
	//表分类
	switch form {
	case "user":
		for rows.Next() {
			err := rows.Scan(&u.Username, &u.Password)
			if err != nil {
				fmt.Println(err)
			}
			U = append(U, u)
			date = U
		}
	case "question":
		for rows.Next() {
			err := rows.Scan(&q.UserName, &q.TrueName, &q.LikeFood, &q.Age)
			if err != nil {
				fmt.Println(err)
			}
			Q = append(Q, q)
			date = Q
		}
	case "message":
		for rows.Next() {
			err := rows.Scan(&m.OwnerName, &m.WriteName, &m.Message, &m.Respond)
			if err != nil {
				fmt.Println(err)
			}
			M = append(M, m)
			date = M
		}
	default:
		break
	}
	return date
}

// BringDate 将数据存入数据库
func BringDate(form string, U model.Use) {
	var dns = "root:040818@tcp(127.0.0.1:3306)/message_board?charset=utf8mb4&parseTime=True&loc=Local"
	u := U.User
	q := U.Ques
	m := U.Mess
	db, _ := sql.Open("mysql", dns)
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
	default:
		break
	}
}

