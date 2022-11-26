package dao

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

// GetComment 获得发表的评论
func GetComment(c *gin.Context) (string, string, string) {
	username := c.PostForm("username")
	writeTo := c.PostForm("write_to")
	comment := c.PostForm("comment")
	return username, writeTo, comment
}
