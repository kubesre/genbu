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
	params := new(struct {
		Name  string `form:"name"`
		Page  int    `form:"page"`
		Limit int    `form:"limit"`
	})

	clusterID := ctx.Param("cid")

	if err := ctx.Bind(params); err != nil {
		global.ReturnContext(ctx).Failed("绑定参数失败", err.Error())
		return
	}

	data, err := service.NewK8sInterface().GetK8sClusterNodeList(clusterID, params.Name, params.Page, params.Limit)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", data)
}
