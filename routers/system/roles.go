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

func InitRolesRouters(r *gin.RouterGroup) gin.IRoutes {
	RoleRouter := r.Group("system")
	{
		RoleRouter.GET("/role/getRoleInfo", system.RolesInfo)
		RoleRouter.POST("/role/createRole", system.AddRole)
		RoleRouter.POST("/role/updateRole", system.UpdateRole)
		RoleRouter.POST("/role/deleteRole", system.DelRole)
		RoleRouter.GET("/role/getRoleList", system.ListRole)
	}
	return r
}
