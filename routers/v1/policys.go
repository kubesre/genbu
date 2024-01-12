/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package v1

import (
	"genbu/controllers/casbin"
	"genbu/controllers/role"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitPolicyRouters(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	{
		r.Use(authMiddleware.MiddlewareFunc())
		r.POST("/policy/role/api/addPolicy", casbin.AddCasbin)
		r.POST("/policy/role/api/deletePolicy", casbin.DelPolicy)
		r.GET("/policy/role/api/listPolicy", casbin.ListPolicy)
		r.POST("/policy/role/menu/addPolicy", role.AddRelationRoleAndMenu)
	}
	return r
}
