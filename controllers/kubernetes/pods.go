package kubernetes

import (
	"fmt"
	"genbu/common/global"
	service "genbu/service/kubernetes"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	corev1 "k8s.io/api/core/v1"
	"log"
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
	ret, err := service.NewK8sInterface().CreatePod(clusterID, &myPod)
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
	// fmt.Println(pods)
	ret, err := service.NewK8sInterface().DeletePod(clusterID, namespace, pods)
	if err != nil {
		global.ReturnContext(c).Failed("删除Pod失败", err.Error())
		return
	}
	global.ReturnContext(c).Successful("删除Pod成功", ret)
}

func GetPodLogs(c *gin.Context) {
	clusterID := c.Param("cid")
	namespace := c.Query("namespace")
	name := c.Query("name")
	follow := c.Query("follow")

	followB, err := strconv.ParseBool(follow)
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}

	ret, err := service.NewK8sInterface().GetPodLogs(clusterID, namespace, name, followB)
	if err != nil {
		global.ReturnContext(c).Failed("获取Log失败", err.Error())
		return
	}
	global.ReturnContext(c).Successful("获取Pod成功", ret)
}

func WatchPod(c *gin.Context) {
	clusterID := c.Param("cid")
	namespace := c.Query("namespace")

	ws, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	if err != nil {
		global.ReturnContext(c).Failed("创建websocket失败", err.Error())
		return
	}
	pw, _ := service.NewK8sInterface().WatchPod(clusterID, namespace)

	fmt.Println(pw)
	// for event := range pw.ResultChan() {

	// 	// c.Writer.Write([]byte(message))
	// 	// err = ws.WriteMessage(websocket.TextMessage, []byte(message))
	// 	// if err != nil {
	// 	// 	log.Println("Error writing to WebSocket:", err.Error())
	// 	// 	return
	// 	// }
	// }
	ws.Close()

}

func ExecPod(c *gin.Context) {
	clusterID := c.Param("cid")
	namespace := c.Query("namespace")

	ws, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	if err != nil {
		global.ReturnContext(c).Failed("创建websocket失败", err.Error())
		return
	}
	fmt.Println(ws)
	pw, err := service.NewK8sInterface().WatchPod(clusterID, namespace)
	fmt.Println(pw)
	for event := range pw.ResultChan() {
		pod, ok := event.Object.(*corev1.Pod)
		if !ok {
			continue
		}

		message := fmt.Sprintf("Pod %s [%s]: %s", pod.Name, event.Type, pod.Status.Phase)
		err = ws.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("Error writing to WebSocket:", err.Error())
			return
		}
	}

}
