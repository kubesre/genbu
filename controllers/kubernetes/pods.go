package kubernetes

import (
	"fmt"
	"genbu/common/global"
	service "genbu/service/kubernetes"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	"strconv"
)

/*
@Auth: Menah3m
@CreateTime: 2024/1/19
@Desc:
*/

type MyPod = corev1.Pod

func GetPodList(ctx *gin.Context) {
	clusterID := ctx.Param("cid")
	namespace := ctx.Query("namespace")
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

	ret, err := service.NewK8sInterface().GetPodList(clusterID, namespace, page, pageSize)
	if err != nil {
		global.ReturnContext(ctx).Failed("获取结果失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("获取结果成功", ret)
}

func GetPod(ctx *gin.Context) {
	clusterID := ctx.Param("cid")
	namespace := ctx.Query("namespace")
	name := ctx.Query("name")

	ret, err := service.NewK8sInterface().GetPod(clusterID, namespace, name)
	if err != nil {
		global.ReturnContext(ctx).Failed("获取结果失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("获取结果成功", ret)
}

func CreatePod(ctx *gin.Context) {
	clusterID := ctx.Param("cid")
	var myPod MyPod
	err := ctx.ShouldBindJSON(&myPod)
	fmt.Println(myPod)
	ret, err := service.NewK8sInterface().CreatePod(clusterID, &myPod)
	if err != nil {
		global.ReturnContext(ctx).Failed("获取结果失败", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("获取结果成功", ret)
}
