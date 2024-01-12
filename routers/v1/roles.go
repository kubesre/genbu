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
		r.GET("/role/infoRole/:rid", role.RolesInfo)
		r.POST("/role/addRole", role.AddRole)
		r.POST("/role/updateRole", role.UpdateRole)
		r.POST("/role/deleteRole", role.DelRole)
		r.GET("/role/listRole", role.ListRole)
	}
	return r
}
