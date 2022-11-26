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
	if ok {
		respond.LoginTrue(c)
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
