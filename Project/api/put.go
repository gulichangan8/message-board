package api

import (
	"Project/dao"
	"Project/service"
	"github.com/gin-gonic/gin"
)

// Change 修改密码的接口
func Change() {
	r := gin.Default()
	r.PUT("/change_password", func(c *gin.Context) {
		u, p := dao.GetUsernamePassword(c)
		ok := service.CheckPassword(p)
		if ok {
			dao.ChangePassword(u, p)
			c.String(200, "修改成功")
		} else {
			c.String(200, "修改失败，密码格式不正确\n提示：1.密码不能含有\\字符\n2.密码不能小于六个字符\n")
		}
	})
}
