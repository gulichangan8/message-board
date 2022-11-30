package api

import (
	"Project/dao"
	"Project/respond"
	"Project/service"
	"Project/tool"
	"github.com/gin-gonic/gin"
)

// Register 注册接口
func Register(c *gin.Context) {
	u, p := service.GetUsernamePassword(c)
	ok1 := service.CheckUsername(u)
	ok2 := service.CheckPassword(p)
	P := service.Hmac("a", p)
	U := tool.CreateUser(u, P)
	if ok1 && ok2 {
		respond.RegisterTrue(c)
		dao.BringUserDate(U)
	} else {
		respond.RegisterErr(c)
	}
}

// Login 登录接口
func Login(c *gin.Context) {
	u, p := service.GetUsernamePassword(c)
	P := service.Hmac("a", p)
	ok := service.CheckLogin(u, P)
	if ok {
		respond.LoginTrue(c)
		l := tool.CreateLogin(u, true)
		dao.BringLoginDate(l)
	} else {
		respond.LoginErr(c)
	}
}

// Question 设置密保接口
func Question(c *gin.Context) {
	u, t, l, a := service.GetQuestion(c)
	q := tool.CreateQues(u, t, l, a)
	dao.BringQuestionDate(q)
	respond.QuestionTrue(c)
}

// AnswerQuestion 回答密保接口
func AnswerQuestion(c *gin.Context) {
	u, t, l, a := service.GetQuestion(c)
	ok := service.PasswordProtect(u, t, l, a)
	if ok {
		respond.AnswerQuestionTrue(c)
	} else {
		respond.AnswerQuestionErr(c)
	}
}

// ChangePassword 修改密码接口
func ChangePassword(c *gin.Context) {
	u, p := service.GetUsernamePassword(c)
	ok := service.CheckPassword(p)
	if ok {
		dao.ChangePasswordDate(u, p)
		respond.ChangePasswordTrue(c)
	} else {
		respond.ChangePasswordErr(c)
	}
}

// PersonalMessage 注册个人信息
func PersonalMessage(c *gin.Context) {
	u, a, b, con, s := service.GetPersonalMessage(c)
	p := tool.CreateMessage(u, a, b, con, s)
	dao.BringPersonDate(p)
	respond.PersonalMessageTrue(c)
}

// ChangePersonalMessage 更改个人信息
func ChangePersonalMessage(c *gin.Context) {
	u, a, b, con, s := service.GetPersonalMessage(c)
	dao.ChangePersonalMessage(u, a, b, con, s)
	respond.ChangePersonalMessageTrue(c)
}

// DeleteLogin 注销登陆状态
func DeleteLogin(c *gin.Context) {
	u := service.GetUsername(c)
	dao.ChangeLogin(u)
	respond.DeleteLoginTrue(c)
}
