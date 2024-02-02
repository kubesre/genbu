/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package v1

import (
	"genbu/controllers/k8s"
	"genbu/middles"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitK8sRouters(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	{
		r.Use(authMiddleware.MiddlewareFunc())
		r.POST("/k8s/cluster/addCluster", k8s.AddK8sCluster)
		r.GET("/k8s/cluster/listCluster", k8s.ListK8sCluster)
		r.POST("/k8s/cluster/deleteCluster", k8s.DeleteK8sCluster)
		r.POST("/k8s/cluster/updateCluster", k8s.UpdateK8sCluster)
		r.POST("/k8s/cluster/refreshCluster", k8s.RefreshK8sCluster)
		r.Use(middles.K8sClientCache())
		r.GET("/k8s/cluster/:cid/node/listNode", k8s.GetK8sClusterNodeList)
		r.GET("/k8s/cluster/:cid/statefulset/listStatefulSet", k8s.ListstatefulSet)
	}
	return r
}
