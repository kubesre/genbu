package kubernetes

import (
	"genbu/controllers/kubernetes"
	"github.com/gin-gonic/gin"
)

func InitSecretRouters(r *gin.RouterGroup) gin.IRouter {
	r = r.Group("kubernetes/:cid")
	{
		r.GET("/secret/getSecretList", kubernetes.ListK8sSecret)
		r.POST("/secret/createSecret", kubernetes.CreateK8sSecret)
	}
	return r
}
