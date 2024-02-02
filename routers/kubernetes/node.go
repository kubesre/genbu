package kubernetes

import (
	"genbu/controllers/kubernetes"
	"github.com/gin-gonic/gin"
)

func InitNodeRouters(r *gin.RouterGroup) gin.IRoutes {
	r = r.Group("/kubernetes/:cid")
	{
		r.GET("/node/getNodeList", kubernetes.GetK8sClusterNodeList)
		r.GET("/node/getNodeDetail", kubernetes.GetK8sClusterNodeDetail)
	}
	return r
}
