package api

import (
	"Project/dao"
	"Project/model"
	"Project/service"
	"Project/tool"
	"github.com/gin-gonic/gin"
)

// Register 注册接口
func Register() {
	r := gin.Default()
	r.POST("/register", func(c *gin.Context) {
		u, p := dao.GetUsernamePassword(c)
		ok1 := service.CheckUsername(u)
		ok2 := service.CheckPassword(p)
		U := tool.CreateUser(u, p)
		var use model.Use
		use.User = U
		if ok1 && ok2 {
			c.String(200, "注册成功")
			dao.BringDate("user", "user", use)
		} else {
			c.String(200, "提示：1.用户名不能含有\\字符\n2.用户名不能小于六个字符\n3.密码不能含有\\字符\n4.密码不能小于六个字符\n")
		}
	})
	err := r.Run()
	if err != nil {
		return
	}
}

// Login 登录接口
func Login() {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		u, p := dao.GetUsernamePassword(c)
		ok := service.CheckLogin(u, p)
		if ok {
			c.String(200, "登陆成功")
		} else {
			c.String(200, "登陆失败")
		}
	})
	err := r.Run()
	if err != nil {
		return
	}
}

// Question 设置密保问题
func Question() {
	r := gin.Default()
	r.POST("/question", func(c *gin.Context) {
		u, t, l, a := dao.GetQuestion(c)
		q := tool.CreateQues(u, t, l, a)
		var Q model.Use
		Q.Ques = q
		dao.BringDate("user", "message", Q)
	})
}

// AnswerQuestion 回答密保
func AnswerQuestion() {
	r := gin.Default()
	r.POST("/answer_question", func(c *gin.Context) {
		u, t, l, a := dao.GetQuestion(c)
		ok := service.PasswordProtect(u, t, l, a)
		if ok {
			c.String(200, "密保问题回答正确，请输入新密码")
		} else {
			c.String(200, "密保问题回答错误，请重新回答")
		}
	})
}