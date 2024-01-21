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

func GetPodList(c *gin.Context) {
	clusterID := c.Param("cid")
	namespace := c.Query("namespace")
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		global.ReturnContext(c).Failed("获取page参数失败", err.Error())
		return
	}
	pageSizeStr := c.Query("pagesize")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		global.ReturnContext(c).Failed("获取pageSize参数失败", err.Error())
		return
	}

	ret, err := service.NewK8sInterface().GetPodList(clusterID, namespace, page, pageSize)
	if err != nil {
		global.ReturnContext(c).Failed("获取Pod List失败", err.Error())
		return
	}
	global.ReturnContext(c).Successful("获取Pod List成功", ret)
}

func GetPod(c *gin.Context) {
	clusterID := c.Param("cid")
	namespace := c.Query("namespace")
	name := c.Query("name")

	ret, err := service.NewK8sInterface().GetPod(clusterID, namespace, name)
	if err != nil {
		global.ReturnContext(c).Failed("获取Pod失败", err.Error())
		return
	}
	global.ReturnContext(c).Successful("获取Pod成功", ret)
}

func CreatePod(c *gin.Context) {
	clusterID := c.Param("cid")
	var myPod MyPod
	err := c.ShouldBindJSON(&myPod)
	fmt.Println(myPod)
	ret, err := service.NewK8sInterface().UpdatePod(clusterID, &myPod)
	if err != nil {
		global.ReturnContext(c).Failed("创建Pod失败", err.Error())
		return
	}
	global.ReturnContext(c).Successful("创建Pod成功", ret)
}

func UpdatePod(c *gin.Context) {
	clusterID := c.Param("cid")
	var myPatch MyPod
	err := c.ShouldBindJSON(&myPatch)
	fmt.Println(myPatch)
	ret, err := service.NewK8sInterface().UpdatePod(clusterID, &myPatch)
	if err != nil {
		global.ReturnContext(c).Failed("更新Pod失败", err.Error())
		return
	}
	global.ReturnContext(c).Successful("更新Pod成功", ret)
}

func DeletePod(c *gin.Context) {
	clusterID := c.Param("cid")
	namespace := c.Query("namespace")
	pods := c.QueryArray("pod")
	fmt.Println(pods)
	ret, err := service.NewK8sInterface().DeletePod(clusterID, namespace, pods)
	if err != nil {
		global.ReturnContext(c).Failed("删除Pod失败", err.Error())
		return
	}
	global.ReturnContext(c).Successful("删除Pod成功", ret)
}
