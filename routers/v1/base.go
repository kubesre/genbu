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

// 基础路由

func InitBaseRouters(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	{
		r.POST("/login", authMiddleware.LoginHandler)     // 登录
		r.POST("/logout", authMiddleware.LogoutHandler)   // 退出
		r.POST("/refresh", authMiddleware.RefreshHandler) // 刷新令牌
		r.POST("/register", users.Register)               // 注册
	}
	return r
}
