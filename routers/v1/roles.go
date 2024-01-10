/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package v1

import (
	"genbu/controllers/role"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitRolesRouters(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	{
		r.Use(authMiddleware.MiddlewareFunc())
		r.GET("/role/info", role.RolesInfo)
		r.POST("/role/add", role.AddRole)
		r.POST("/role/update", role.UpdateRole)
		r.POST("/role/bind_menu", role.AddRelationRoleAndMenu)
		r.POST("/role/del", role.DelRole)
		r.GET("/role/list", role.ListRole)
	}
	return r
}
