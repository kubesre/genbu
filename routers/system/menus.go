/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package system

import (
	"genbu/controllers/system"
	"github.com/gin-gonic/gin"
)

func InitMenusRouters(r *gin.RouterGroup) gin.IRoutes {
	r = r.Group("system")
	{
		r.POST("/menu/createMenu", system.AddMenus)
		r.GET("/menu/getMenuList", system.ListMenus)
	}
	return r
}
