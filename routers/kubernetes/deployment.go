/*
@auth: Ansu
@source: Ansu
@time: 2024/1/18
*/

package kubernetes

import (
	"genbu/controllers/kubernetes"
	"github.com/gin-gonic/gin"
)

func InitDeploymentRouters(r *gin.RouterGroup) gin.IRouter {
	r = r.Group("/kubernetes/cluster/:cid")
	{
		r.GET("/deployment/listDeployment", kubernetes.GetDeploymentList)
		r.GET("/deployment/getDeploymentDetails")
		r.GET("/deployment/getDeploymentYaml")
		r.POST("/deployment/deleteDeployment")
		r.POST("/deployment/createDeployment2yaml")
		r.POST("/deployment/createDeployment2arg")
		r.POST("/deployment/updateDeployment2Yaml")
		r.POST("/deployment/updateDeployment2arg")
		//r.Use(middles.K8sClientCache())
	}
	return r
}
