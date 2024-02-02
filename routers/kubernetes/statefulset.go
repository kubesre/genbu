package kubernetes

/*
@auth: Meersburg
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/2/02
*/
import (
	"genbu/controllers/kubernetes"
	"genbu/middles"
	"github.com/gin-gonic/gin"
)

func InitStatefulSetRouters(r *gin.RouterGroup) gin.IRoutes {
	r = r.Group("/kubernetes/:cid")
	{
		r.GET("/statefulset/listStatefulSet", kubernetes.ListstatefulSet)
		r.GET("/statefulset/getStatefulSetDetail", kubernetes.DetailStatefulSet)
		r.POST("/statefulset/deleteStatefulSet", kubernetes.DelectStatefulSet)
		r.POST("/statefulset/createStatefulSetYaml", kubernetes.CreateStatefulSetYaml)
		r.POST("/statefulset/updateStatefulSetYaml", kubernetes.UpdateStatefulSetYaml)
		r.Use(middles.K8sClientCache())
	}
	return r
}
