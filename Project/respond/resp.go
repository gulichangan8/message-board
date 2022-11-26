package respond

import "github.com/gin-gonic/gin"

func RegisterTrue(c *gin.Context) {
	c.String(200, "注册成功")
}

func RegisterErr(c *gin.Context) {
	c.String(200, "提示：1.用户名不能含有\\字符\n2.用户名不能小于六个字符\n3.密码不能含有\\字符\n4.密码不能小于六个字符\n")
}

func LoginTrue(c *gin.Context) {
	c.String(200, "登陆成功")
}

func LoginErr(c *gin.Context) {
	c.String(200, "登陆失败")
}

func QuestionTrue(c *gin.Context) {
	c.String(200, "设置密保成功")
}

func AnswerQuestionTrue(c *gin.Context) {
	c.String(200, "密保问题回答正确，请输入新密码")
}

func AnswerQuestionErr(c *gin.Context) {
	c.String(200, "密保问题回答错误，请重新回答")
}

func ChangePasswordTrue(c *gin.Context) {
	c.String(200, "修改成功")
}

func ChangePasswordErr(c *gin.Context) {
	c.String(200, "修改失败，密码格式不正确\n提示：1.密码不能含有\\字符\n2.密码不能小于六个字符\n")
}

func PublishMessageTrue(c *gin.Context) {
	c.String(200, "留言发表成功")
}

func ChangeMessageTrue(c *gin.Context) {
	c.String(200, "留言修改成功")
}

func DeleteMessageTrue(c *gin.Context) {
	c.String(200, "留言删除成功")
}
