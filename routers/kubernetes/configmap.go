/*
@auth: 啷个办
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/17
*/
package kubernetes

import (
	"genbu/controllers/kubernetes"

	"github.com/gin-gonic/gin"
)

func InitConfigRouters(r *gin.RouterGroup) gin.IRoutes {
	r = r.Group("kubernetes")

	{
		r.GET("/config/getConfigList", kubernetes.ListK8sConfigMap)
	}
	return r
}
