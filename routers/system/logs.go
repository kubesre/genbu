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

func InitLogRouters(r *gin.RouterGroup) gin.IRoutes {
	r = r.Group("system")
	{
		r.GET("/log", system.GetOperationLogList)
	}
	return r
}
