package service

import (
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
