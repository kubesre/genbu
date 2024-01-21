/*
@auth: Ansu
@source: Ansu
@time: 2024/1/18
*/

package kubernetes

import (
	"genbu/common/global"
	service "genbu/service/kubernetes"
	"github.com/gin-gonic/gin"
)

// get deployment list

func GetDeploymentList(ctx *gin.Context) {
	clusterID := ctx.Param("cid")
	err := service.GetDeploymentList(clusterID, "default")
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", "success")

}
