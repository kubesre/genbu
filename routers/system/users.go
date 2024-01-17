/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package system

import (
	"genbu/controllers/system"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitUserRouters(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	r = r.Group("system")
	{
		r.POST("/user/logout", authMiddleware.LogoutHandler)   // 退出
		r.POST("/user/refresh", authMiddleware.RefreshHandler) // 刷新令牌
		//r.POST("/system/user/addUser", users.Register)                // 注册
		r.GET("/user/getUserInfo", system.GetUserInfo)
		r.GET("/user/getUserList", system.UserList)
		r.POST("/user/updateUser", system.UserUpdate)
		r.POST("/user/createUser", system.UserAdd)
	}
	return r
}
