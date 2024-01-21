package kubernetes

import (
	"genbu/controllers/kubernetes"
	"github.com/gin-gonic/gin"
)

/*
   @Auth: Menah3m
   @CreateTime: 2024/1/19
   @Desc:
*/

func InitPodRouters(r *gin.RouterGroup) gin.IRoutes {
	r = r.Group("/kubernetes/:cid")
	{
		r.GET("/pod/getPodList", kubernetes.GetPodList)
		r.GET("/pod/getPod", kubernetes.GetPod)
		r.POST("/pod/createPod", kubernetes.CreatePod)
		r.DELETE("/pod/deletePod", kubernetes.DeletePod)
		r.PUT("/pod/updatePod", kubernetes.UpdatePod)

	}
	return r
}
