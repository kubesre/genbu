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
	r = r.Group("kubernetes/:cid")
	{
		r.GET("/config/getConfigList", kubernetes.ListK8sConfigMap)   // 获取节点上的所有ConfigMap
		r.GET("/config/getConfigInfo", kubernetes.Get8sConfigMapInfo) // 获取某个指定的ConfigMap
		r.POST("/config/deleteConfig", kubernetes.DeleteConfig)       // 删除指定的ConfigMap
		r.POST("/config/deleteConfigs", kubernetes.DeleteConfigs)     // 删除多个ConfigMap
		r.POST("/config/createConfig", kubernetes.CreateConfigMap)    // 创建ConfigMap
		r.POST("/config/updateConfig", kubernetes.UpdateConfigMap)    // 更新ConfigMap
	}
	return r
}
