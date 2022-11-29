package service

import (
	"Project/dao"
	"Project/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

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

// GetMessage 获得发表的留言
func GetMessage(c *gin.Context) (string, string, string) {
	username := c.PostForm("username")
	writeTo := c.PostForm("write_to")
	message := c.PostForm("message")
	return username, writeTo, message
}

// GetNewMessage 获得修改的留言
func GetNewMessage(c *gin.Context) (string, string) {
	username := c.PostForm("username")
	message := c.PostForm("message")
	return username, message
}

// GetUsername 获得用户名
func GetUsername(c *gin.Context) string {
	username := c.PostForm("username")
	return username
}

// GetUsernameRespond 获得用户名回复
func GetUsernameRespond(c *gin.Context) (string, string) {
	username := c.PostForm("username")
	respond := c.PostForm("respond")
	return username, respond
}

// GetUsernameCommentWriteTo 获得用户名评论，作品作者
func GetUsernameCommentWriteTo(c *gin.Context) (string, string, string) {
	username := c.PostForm("username")
	author := c.PostForm("write_to")
	comment := c.PostForm("comment")
	return author, username, comment
}

// GetPersonalMessage 获得个人信息
func GetPersonalMessage(c *gin.Context) (string, int, float64, string, string) {
	username := c.PostForm("username")
	age, _ := strconv.Atoi(c.PostForm("age"))
	b, _ := strconv.Atoi(c.PostForm("age"))
	birthday := float64(b)
	constellation := c.PostForm("constellation")
	sex := c.PostForm("sex")
	return username, age, birthday, constellation, sex
}

// GetGood 获得点赞信息
func GetGood(c *gin.Context) (string, string, bool) {
	reader := c.PostForm("username")
	author := c.PostForm("author")
	g := c.PostForm("good")
	good, _ := strconv.ParseBool(g)
	return author, reader, good
}

// GetAuthorComment 获得该作者下的所有评论
func GetAuthorComment(c *gin.Context) model.BiTree {
	author := c.PostForm("author")
	bi := dao.TakeOutAuthorComment(author)
	return bi
}

func GetCom(c *gin.Context) (int, string, string, string) {
	parentId, _ := strconv.Atoi(c.PostForm("id"))
	comment := c.PostForm("comment")
	writer := c.PostForm("username")
	author := c.PostForm("author")
	return parentId, author, writer, comment
}
