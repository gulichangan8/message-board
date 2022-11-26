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
	Com := tool.CreateMess(w, u, com, "")
	var C model.Use
	C.Mess = Com
	dao.BringDate("message", C)
	respond.PublishMessageTrue(c)
}

// ChangeMessage 修改留言
func ChangeMessage(c *gin.Context) {
	u, m := service.GetNewMessage(c)
	respond.ChangeMessageTrue(c)
	dao.ChangeMessageDate(u, m)
}

// DeleteMessage 删除留言
func DeleteMessage(c *gin.Context) {
	u := service.GetUsername(c)
	dao.DeleteMessageDate(u)
	respond.DeleteMessageTrue(c)
}

// RespondMessage 回复留言
func RespondMessage(c *gin.Context) {
	u, r := service.GetUsernameRespond(c)
	dao.RespondMessageDate(u, r)
	respond.ResMessageTrue(c)
}

// ReadMessage 查看留言回复
func ReadMessage(c *gin.Context) {
	u := service.GetUsername(c)
	m := dao.TakeDate("message")
	M, _ := m.(model.Messes)
	var d model.Mess
	ok := false
	for _, date := range M {
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
