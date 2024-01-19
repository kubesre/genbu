package kubernetes

import (
	"genbu/common/global"
	"genbu/service/k8s"
	"github.com/gin-gonic/gin"
	"strconv"
)

/*
@Auth: Menah3m
@CreateTime: 2024/1/19
@Desc:
*/
func GetK8sClusterPodList(ctx *gin.Context) {
	clusterID := ctx.Keys["cluster"]
	namespace := ctx.Query("namespace")
	name := ctx.Query("name")
	pageStr := ctx.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		global.ReturnContext(ctx).Failed("获取page参数失败", err.Error())
		return
	}
	pageSizeStr := ctx.Query("pagesize")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		global.ReturnContext(ctx).Failed("获取pageSize参数失败", err.Error())
		return
	}

	err = k8s.NewK8sInterface().GetK8sClusterPodList(clusterID, namespace, name, page, pageSize)
	if err != nil {
		global.ReturnContext(ctx).Failed("获取结果失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("获取结果成功", "success")
}