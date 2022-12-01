package dao

import (
	"Project/model"
	"database/sql"
	"github.com/gin-gonic/gin"
	"strconv"
)

// TakeUserDate 将user表中的数据取出
func TakeUserDate() model.Users {
	rows, _ := dB.Query("select * from ?", "user")
	var u model.User
	var U model.Users
	for rows.Next() {
		err := rows.Scan(&u.Username, &u.Password)
		if err != nil {
			panic(err)
		}
		U = append(U, u)
	}
	return U
}

// TakeQuestionDate 将question表中的数据取出
func TakeQuestionDate() model.Queses {
	rows, _ := dB.Query("select * from ?", "user")
	var q model.Ques
	var Q model.Queses
	for rows.Next() {
		err := rows.Scan(&q.UserName, &q.TrueName, &q.LikeFood, &q.Age)
		if err != nil {
			panic(err)
		}
		Q = append(Q, q)
	}
	return Q
}

// TakePersonDate 将person表中的数据取出
func TakePersonDate() model.Messages {
	rows, _ := dB.Query("select * from ?", "user")
	var p model.Message
	var P model.Messages
	for rows.Next() {
		err := rows.Scan(&p.Username, &p.Age, &p.Birthday, &p.Constellation, &p.Sex)
		if err != nil {
			panic(err)
		}
		P = append(P, p)
	}
	return P
}

// TakeLoginDate 将login表中的数据取出
func TakeLoginDate() model.Logins {
	rows, _ := dB.Query("select * from ?", "user")
	var l model.Login
	var L model.Logins
	for rows.Next() {
		err := rows.Scan(&l.Username, &l.Login)
		if err != nil {
			panic(err)
		}
		L = append(L, l)
	}
	return L
}

// BringUserDate 将数据存入user表
func BringUserDate(u model.User) {
	_, err := dB.Exec("insert into user (username,password) value (?,?)",
		u.Username, u.Password)
	if err != nil {
		return
	}
}

// BringQuestionDate 将数据存入question表
func BringQuestionDate(q model.Ques) {
	_, err := dB.Exec("insert into question (username,true_name,like_food,age) value (?,?,?,?)",
		q.UserName, q.TrueName, q.LikeFood, q.Age)
	if err != nil {
		return
	}
}

// BringPersonDate 将数据存入person表
func BringPersonDate(p model.Message) {
	_, err := dB.Exec("insert into person (username,age,birthday,constellation,sex) value (?,?,?,?,?)",
		p.Username, p.Age, p.Birthday, p.Constellation, p.Sex)
	if err != nil {
		return
	}
}

// BringLoginDate 将数据存入login表
func BringLoginDate(l model.Login) {
	_, err := dB.Exec("insert into login (username,login) value (?,?)",
		l.Username, l.Login)
	if err != nil {
		return
	}
}

// ChangePasswordDate 修改密码
func ChangePasswordDate(username string, password string) {
	var dns = "root:040818@tcp(127.0.0.1:3306)/message_board?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := sql.Open("mysql", dns)
	_, err := db.Exec("update user set password=? where username=?", password, username)
	if err != nil {
		return
	}
}

// ChangePersonalMessage 修改个人信息
func ChangePersonalMessage(username string, age int, birthday float64, constellation string, sex string) {
	var dns = "root:040818@tcp(127.0.0.1:3306)/message_board?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := sql.Open("mysql", dns)
	_, err := db.Exec("update person set age=?,birthday=?,constellation=?,sex=? where username=?", age, birthday, constellation, sex, username)
	if err != nil {
		return
	}
}

// ChangeLogin 修改登陆状态
func ChangeLogin(username string) {
	var dns = "root:040818@tcp(127.0.0.1:3306)/message_board?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := sql.Open("mysql", dns)
	_, err := db.Exec("update login set login=? where username=?", true, username)
	if err != nil {
		return
	}
}

// GetUsernamePassword 获得用户名密码
func GetUsernamePassword(c *gin.Context) (string, string) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	return username, password
}

// GetQuestion 获得密保问题答案
func GetQuestion(c *gin.Context) (string, string, string, int) {
	username := c.PostForm("username")
	trueName := c.PostForm("true_name")
	likeFood := c.PostForm("like_food")
	age, _ := strconv.Atoi(c.PostForm("age"))
	return username, trueName, likeFood, age
}
