/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/10
*/

package kubernetes

import (
	"genbu/common/global"
	service "genbu/service/kubernetes"
	"github.com/gin-gonic/gin"
)

func GetK8sClusterNodeList(ctx *gin.Context) {
	clusterID := ctx.Param("cid")
	err := service.NewK8sInterface().GetK8sClusterNodeList(clusterID)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", "success")
}
