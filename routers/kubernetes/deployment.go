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
		r.GET("/deployment/getDeploymentDetails", kubernetes.GetDeploymentDetails)
		r.GET("/deployment/getDeploymentYaml", kubernetes.GetDeploymentYaml)
		r.POST("/deployment/deleteDeployment", kubernetes.DeleteDeployment)
		r.POST("/deployment/createOrUpdateDeployment2Yaml", kubernetes.CreateOrUpdateDeployment2Yaml)
		r.POST("/deployment/createOrUpdateDeployment2Arg", kubernetes.CreateOrUpdateDeployment2Arg)
		//r.Use(middles.K8sClientCache())
	}
	return r
}
