/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package kubernetes

import (
	"genbu/common/global"
	mod "genbu/models/kubernetes"
	service "genbu/service/kubernetes"
	"github.com/gin-gonic/gin"
)

// 集群添加

func AddK8sCluster(ctx *gin.Context) {
	params := new(mod.Configs)
	if err := ctx.ShouldBind(&params); err != nil {
		global.TPLogger.Error("k8s集群添加参数校验失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	if err := service.NewK8sInterface().AddK8sCluster(params); err != nil {
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
	data, err := service.NewK8sInterface().ListK8sCluster(params.Name, params.Limit, params.Page)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	global.ReturnContext(ctx).Successful("success", data)
}

// 集群删除

func DeleteK8sCluster(ctx *gin.Context) {
	params := new(struct {
		CID []string `json:"c_id"`
	})
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.TPLogger.Error("集群删除数据绑定失败：", err)
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	if err := service.NewK8sInterface().DeleteK8sCluster(params.CID); err != nil {
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	global.ReturnContext(ctx).Successful("success", "集群删除成功")
}

// 集群更新

func UpdateK8sCluster(ctx *gin.Context) {
	params := new(mod.Configs)
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.TPLogger.Error("集群更新数据绑定失败：", err)
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	if err := service.NewK8sInterface().UpdateK8sCluster(params); err != nil {
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	global.ReturnContext(ctx).Successful("success", "集群更新成功")
}

// 刷新集群

func RefreshK8sCluster(ctx *gin.Context) {
	if err := service.NewK8sInterface().RefreshK8sCluster(); err != nil {
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	global.ReturnContext(ctx).Successful("success", "集群刷新成功")
}
