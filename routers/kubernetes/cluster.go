/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package kubernetes

import (
	"genbu/controllers/kubernetes"
	"genbu/middles"
	"github.com/gin-gonic/gin"
)

func InitClusterRouters(r *gin.RouterGroup) gin.IRoutes {
	r = r.Group("kubernetes")
	{
		r.POST("/cluster/createCluster", kubernetes.AddK8sCluster)
		r.GET("/cluster/getClusterList", kubernetes.ListK8sCluster)
		r.POST("/cluster/deleteCluster", kubernetes.DeleteK8sCluster)
		r.POST("/cluster/updateCluster", kubernetes.UpdateK8sCluster)
		r.POST("/cluster/refreshCluster", kubernetes.RefreshK8sCluster)
		r.Use(middles.K8sClientCache())
	}
	return r
}
