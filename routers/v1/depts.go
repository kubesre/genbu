/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package v1

import (
	"genbu/controllers/dept"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitDeptRouters(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	{
		r.Use(authMiddleware.MiddlewareFunc())
		r.POST("/dept/add", dept.AddDept)
		r.GET("/dept/list", dept.ListDept)
		r.GET("/dept/info", dept.InfoDept)
		r.POST("/dept/del", dept.DelDept)
	}
	return r
}
