/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package v1

import (
	"genbu/controllers/users"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitUserRouters(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	{
		r.Use(authMiddleware.MiddlewareFunc())
		r.GET("/user/info", users.GetUserInfo)
		r.GET("/user/list", users.UserList)
		r.POST("/user/update", users.UserUpdate)
		r.POST("/user/add", users.UserAdd)
	}
	return r
}
