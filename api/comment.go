package api

import (
	"Project/dao"
	"Project/respond"
	"Project/service"
	"Project/tool"
	"github.com/gin-gonic/gin"
)

// 此处皆为一级评论，parent-id自动为1

// PublishComment 发表评论
func PublishComment(c *gin.Context) {
	a, w, com := service.GetUsernameCommentWriteTo(c)
	ok := service.CheckAuthorExist(a)
	if ok {
		C := tool.CreateComment(a, w, com)
		dao.BringCommentDate(C)
		respond.PublishCommentTrue(c)
	} else {
		respond.PublishCommentErr(c)
	}
}

// ChangeComment 修改评论
//func ChangeComment(c *gin.Context) {
//	a, w, com := service.GetUsernameCommentWriteTo(c)
//	ok := service.CheckAuthorExist(a) && service.CheckWriteExist(w)
//	if ok {
//		dao.ChangeCommentDate(com, a)
//		respond.ChangeCommentTrue(c)
//	} else {
//		respond.ChangeCommentErr(c)
//	}
//}
//
// DeleteComment 删除评论
//func DeleteComment(c *gin.Context) {
//	w := service.GetUsername(c)
//	ok := service.CheckAuthorExist(w)
//	if ok {
//		dao.ChangeCommentDate("", w)
//		respond.DeleteCommentTrue(c)
//	} else {
//		respond.ChangeCommentErr(c)
//	}
//}

// ChangeGood 改变点赞状态
func ChangeGood(c *gin.Context) {
	a, r, g := service.GetGood(c)
	p := tool.CreatePraise(a, r, g)
	dao.BringPraiseDate(p)
	if g {
		respond.ChangeGoodTrue(c)
	} else {
		respond.ChangeGoodErr(c)
	}
}

// NoNameComment 匿名评论
//func NoNameComment(c *gin.Context) {
//	a, _, com := service.GetUsernameCommentWriteTo(c)
//	ok := service.CheckAuthorExist(a)
//	if ok {
//		C := tool.CreateComment(a, "", com)
//		dao.BringCommentDate(C)
//		respond.PublishCommentTrue(c)
//	} else {
//		respond.PublishCommentErr(c)
//	}
//}

// GetGoodMember 查看点赞数
func GetGoodMember(c *gin.Context) {
	member := 0
	p := dao.TakePraiseDate()
	a := service.GetUsername(c)
	for _, date := range p {
		if date.Author == a && date.Good {
			member++
		} else {
			continue
		}
	}
	g := tool.CreateGood(a, member)
	dao.BringGoodMemberDate(g)
	respond.GetGoodMemberTrue(c, g)
}
