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

/*
   @Auth: Menah3m
   @CreateTime: 2024/1/19
   @Desc: pods 相关操作
*/

func (k *k8sCluster) GetK8sClusterPodList(cid interface{}, namespace, name string, page, pagesize int) (err error) {
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
	// nodeList, err := clientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	podsList, err := clientSet.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		global.TPLogger.Error("获取pod失败：", err)
		return errors.New("获取pod失败")
	}
	for _, item := range podsList.Items {
		fmt.Println(item.Name)
	}
	return nil

}
