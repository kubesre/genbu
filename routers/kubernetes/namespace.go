package kubernetes

import (
	"genbu/controllers/kubernetes"
	"github.com/gin-gonic/gin"
)

func InitNameSpaceRouters(r *gin.RouterGroup) gin.IRoutes {
	r = r.Group("/kubernetes/:cid")
	{
		r.GET("/namespace/getNameSpaceList", kubernetes.GetK8sNameSpaceList)
		r.POST("/namespace/createNameSpace", kubernetes.CreateK8sNameSpace)
		r.POST("/namespace/deleteNameSpace", kubernetes.DeleteNamespace)
		r.POST("/namespace/updateNameSpace", kubernetes.UpdateNamespace)

	}
	return r
}
