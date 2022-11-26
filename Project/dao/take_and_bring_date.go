package dao

import (
	"Project/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// TakeDate 取出数据库library中form表中的数据
// 参数：库名 表名	  返回值：结构体数组
func TakeDate(library string, form string) (date interface{}) {
	var u model.User
	var m model.Mess
	var q model.Ques
	var U model.Users
	var M model.Messes
	var Q model.Queses
	var dns = "root:040818@tcp(127.0.0.1:3306)/" + library + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := sql.Open("mysql", dns)
	rows, _ := db.Query("select * from ?", form)
	//库分类
	switch library {
	case "user":
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
		case "message":
			for rows.Next() {
				err := rows.Scan(&m.OwnerName, &m.WriteName, &m.Message, &m.Respond)
				if err != nil {
					fmt.Println(err)
				}
				M = append(M, m)
				date = M
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
		default:
			break
		}
	default:
		break
	}
	return date
}

// BringDate 将数据存入数据库
// 参数：库名 表名 结构体组合	返回值：无
func BringDate(library string, form string, U model.Use) {
	var dns = "root:040818@tcp(127.0.0.1:3306)/" + library + "?charset=utf8mb4&parseTime=True&loc=Local"
	u := U.User
	m := U.Mess
	q := U.Ques
	db, _ := sql.Open("mysql", dns)
	//库分类
	switch library {
	case "user":
		//表分类
		switch form {
		case "user":
			_, err := db.Exec("insert into ? (username,password) value (?,?)",
				form, u.Username, u.Password)
			if err != nil {
				return
			}
		case "message":
			_, err := db.Exec("insert into ? (owner_name,write_name,message,respond) value (?,?,?,?)",
				form, m.OwnerName, m.WriteName, m.Message, m.Respond)
			if err != nil {
				return
			}
		case "question":
			_, err := db.Exec("insert into ? (username,true_name,like_food,age) value (?,?,?,?)",
				form, q.UserName, q.TrueName, q.LikeFood, q.Age)
			if err != nil {
				return
			}
		default:
			break
		}
	default:
		break
	}
}
