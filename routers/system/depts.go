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

func InitDeptRouters(r *gin.RouterGroup) gin.IRoutes {
	r = r.Group("system")
	{
		r.POST("/dept/createDept", system.AddDept)  // 添加
		r.GET("/dept/getDeptList", system.ListDept) // 列表
		// r.GET("/dept/:dept_id", dept.InfoDept)   // 详情
		r.POST("/dept/deleteDept", system.DelDept) // 删除
	}
	return r
}
