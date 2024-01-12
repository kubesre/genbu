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
		r.POST("/dept/addDept", dept.AddDept)  // 添加
		r.GET("/dept/listDept", dept.ListDept) // 列表
		// r.GET("/dept/:dept_id", dept.InfoDept)   // 详情
		r.POST("/dept/deleteDept", dept.DelDept) // 删除
	}
	return r
}
