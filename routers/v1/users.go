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
		r.POST("/system/user/logout", authMiddleware.LogoutHandler)   // 退出
		r.POST("/system/user/refresh", authMiddleware.RefreshHandler) // 刷新令牌
		r.POST("/system/user/register", users.Register)               // 注册
		r.GET("/system/user/infoUser", users.GetUserInfo)
		r.GET("/system/user/listUser", users.UserList)
		r.POST("system/user/updateUser", users.UserUpdate)
		//r.POST("/user", users.UserAdd)
	}
	return r
}
