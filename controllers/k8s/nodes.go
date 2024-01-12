/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/10
*/

package k8s

import (
	"genbu/common/global"
	"genbu/service/k8s"
	"github.com/gin-gonic/gin"
)

func GetK8sClusterNodeList(ctx *gin.Context) {
	clusterID := ctx.Keys["cluster"]
	err := k8s.NewK8sInterface().GetK8sClusterNodeList(clusterID)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", "success")
}
