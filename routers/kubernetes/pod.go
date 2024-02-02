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
		r.GET("/pod/list", kubernetes.GetPodList)
		r.GET("/pod/describe", kubernetes.GetPod)
		r.POST("/pod/create", kubernetes.CreatePod)
		r.DELETE("/pod/delete", kubernetes.DeletePod)
		r.PUT("/pod/update", kubernetes.UpdatePod)
		r.GET("/pod/logs", kubernetes.GetPodLogs)
		// r.GET("/pod/watch", kubernetes.WatchPod)
		// r.GET("/pod/exec", kubernetes.ExecPod)

	}
	return r
}
