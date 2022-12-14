package api

import (
	"Project/dao"
	"Project/model"
	"Project/respond"
	"Project/service"
	"Project/tool"
	"github.com/gin-gonic/gin"
)

// PublishMessage 发表留言
func PublishMessage(c *gin.Context) {
	u, w, com := service.GetMessage(c)
	ok1 := service.CheckUsernameExist(w)
	ok2 := service.CheckCommentLength(com)
	if ok1 {
		if ok2 {
			Com := tool.CreateMess(w, u, com, "")
			dao.BringMessageDate(Com)
			respond.PublishMessageTrue(c)
		} else {
			respond.PublishMessageErr2(c)
		}
	} else {
		respond.PublishMessageErr1(c)
	}
}

// ChangeMessage 修改留言
func ChangeMessage(c *gin.Context) {
	u, m := service.GetNewMessage(c)
	ok := service.CheckWriterExist(u)
	if ok {
		respond.ChangeMessageTrue(c)
		dao.ChangeMessageDate(u, m)
	} else {
		respond.ChangeMessageErr(c)
	}
}

// DeleteMessage 删除留言
func DeleteMessage(c *gin.Context) {
	u := service.GetUsername(c)
	ok := service.CheckWriterExist(u)
	if ok {
		dao.DeleteMessageDate(u)
		respond.DeleteMessageTrue(c)
	} else {
		respond.DeleteMessageErr(c)
	}
}

// RespondMessage 回复留言
func RespondMessage(c *gin.Context) {
	u, r := service.GetUsernameRespond(c)
	ok := service.IfHaveSomeoneWriteTo(u)
	if ok {
		dao.RespondMessageDate(u, r)
		respond.ResMessageTrue(c)
	} else {
		respond.ResMessageErr(c)
	}
}

// ReadMessage 查看留言回复
func ReadMessage(c *gin.Context) {
	u := service.GetUsername(c)
	m := dao.TakeMessageDate()
	var d model.Mess
	ok := false
	for _, date := range m {
		if date.OwnerName == u {
			d = date
			ok = true
		}
	}
	if ok {
		respond.ReadMessageTrue(c, d)
	} else {
		respond.ReadMessageErr(c)
	}
}
