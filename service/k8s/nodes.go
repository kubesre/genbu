/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/10
*/

package k8s

import (
	"context"
	"errors"
	"fmt"
	"genbu/common/global"
	"genbu/dao/k8s"
	"genbu/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

// 获取node节点信息

func (k *k8sCluster) GetK8sClusterNodeList(cid interface{}) (err error) {
	cidStr, _ := cid.(string)
	cluster, err := k8s.NewK8sInterface().ActiveK8sCluster(cidStr)
	if err != nil {
		global.TPLogger.Error("获取可用集群信息失败: ", err)
		return errors.New("获取可用集群信息失败")
	}
	configStr := cluster.Text
	decodeConfig, err := utils.DecodeBase64(configStr)
	if err != nil {
		global.TPLogger.Error("集群config文件加载失败: ", err)
		return errors.New("集群config文件加载失败")
	}
	clientSet, err := global.NewClientInterface().NewClientSet(decodeConfig)
	if err != nil {
		global.TPLogger.Error("初始化client失败：", err)
		return errors.New("初始化client失败")
	}
	// 获取config
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	ctx, cancel = context.WithTimeout(context.TODO(), time.Second*2)
	defer cancel()
	nodeList, err := clientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		global.TPLogger.Error("获取node失败：", err)
		return errors.New("获取node失败")
	}
	for _, item := range nodeList.Items {
		fmt.Println(item.Name)
	}
	return nil

}
