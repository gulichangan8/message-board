package api

import (
	"Project/dao"
	"Project/model"
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
	U := tool.CreateUser(u, p)
	var use model.Use
	use.User = U
	if ok1 && ok2 {
		respond.RegisterTrue(c)
		dao.BringDate("user", use)
	} else {
		respond.RegisterErr(c)
	}
}

// Login 登录接口
func Login(c *gin.Context) {
	u, p := service.GetUsernamePassword(c)
	ok := service.CheckLogin(u, p)
	var use model.Use
	if ok {
		respond.LoginTrue(c)
		l := tool.CreateLogin(u, true)
		use.Login = l
		dao.BringDate("login", use)
	} else {
		respond.LoginErr(c)
	}
}

// Question 设置密保接口
func Question(c *gin.Context) {
	u, t, l, a := service.GetQuestion(c)
	q := tool.CreateQues(u, t, l, a)
	var Q model.Use
	Q.Ques = q
	dao.BringDate("message", Q)
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
	var P model.Use
	P.Message = p
	dao.BringDate("person", P)
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
