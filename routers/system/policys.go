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

func InitPolicyRouters(r *gin.RouterGroup) gin.IRoutes {
	r = r.Group("system")
	{
		r.POST("/policy/createPolicy", system.AddCasbin)
		r.POST("/policy/deletePolicy", system.DelPolicy)
		r.GET("/policy/getPolicyList", system.ListPolicy)
		r.POST("/policy/menu/createPolicy", system.AddRelationRoleAndMenu)
	}
	return r
}
