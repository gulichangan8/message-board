package respond

import (
	"Project/model"
	"github.com/gin-gonic/gin"
)

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

func PublishMessageErr(c *gin.Context) {
	c.String(200, "留言发表失败，查无此人")
}

func ChangeMessageTrue(c *gin.Context) {
	c.String(200, "留言修改成功")
}

func ChangeMessageErr(c *gin.Context) {
	c.String(200, "您还未发表过留言")
}

func DeleteMessageTrue(c *gin.Context) {
	c.String(200, "留言删除成功")
}

func DeleteMessageErr(c *gin.Context) {
	c.String(200, "您还未发表过留言")
}

func ResMessageTrue(c *gin.Context) {
	c.String(200, "回复留言成功")
}

func ResMessageErr(c *gin.Context) {
	c.String(200, "还未有人给您留过言")
}

func ReadMessageTrue(c *gin.Context, mess model.Mess) {
	c.JSON(200, mess)
}

func ReadMessageErr(c *gin.Context) {
	c.String(200, "查看留言失败，用户名不存在")
}

func PublishProductionTrue(c *gin.Context) {
	c.String(200, "作品发布成功")
}

func PublishCommentTrue(c *gin.Context) {
	c.String(200, "评论发表成功")
}

func PublishCommentErr(c *gin.Context) {
	c.String(200, "评论发表失败，未找到作者")
}

func ChangeCommentTrue(c *gin.Context) {
	c.String(200, "评论修改成功")
}

func ChangeCommentErr(c *gin.Context) {
	c.String(200, "未找到评论")
}

func PersonalMessageTrue(c *gin.Context) {
	c.String(200, "个人信息保存成功")
}

func ChangePersonalMessageTrue(c *gin.Context) {
	c.String(200, "个人信息修改成功")
}

func ChangeGoodTrue(c *gin.Context) {
	c.String(200, "点赞成功")
}

func ChangeGoodErr(c *gin.Context) {
	c.String(200, "取消点赞成功")
}
