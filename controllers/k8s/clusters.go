/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package k8s

import (
	"genbu/common/global"
	"genbu/models/k8s"
	sk "genbu/service/k8s"
	"github.com/gin-gonic/gin"
)

// 集群添加

func AddK8sCluster(ctx *gin.Context) {
	params := new(k8s.Configs)
	if err := ctx.ShouldBind(&params); err != nil {
		global.TPLogger.Error("k8s集群添加参数校验失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	if err := sk.NewK8sInterface().AddK8sCluster(params); err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", "集群添加成功！！！")
}

// 集群列表

func ListK8sCluster(ctx *gin.Context) {
	params := new(struct {
		Name  string `form:"name"`
		Limit int    `form:"limit"`
		Page  int    `form:"page"`
	})
	if err := ctx.ShouldBindQuery(&params); err != nil {
		global.TPLogger.Error("集群列表查询数据绑定失败：", err)
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	data, err := sk.NewK8sInterface().ListK8sCluster(params.Name, params.Limit, params.Page)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	global.ReturnContext(ctx).Successful("success", data)
}
