package kubernetes

import (
	"context"
	"errors"
	"fmt"
	"genbu/common/global"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"time"
)

/*
   @Auth: Menah3m
   @CreateTime: 2024/1/19
   @Desc: pods 相关操作
*/

func (k *k8sCluster) GetK8sClusterPodList(cid string, namespace, name string, page, pageSize int) (ret interface{}, err error) {
	clientSetAny, found := global.ClientSetCache.Get(cid)
	if !found {
		global.TPLogger.Error("当前集群不存在：", err)
		return nil, errors.New("当前集群不存在")
	}
	clientSet := clientSetAny.(*kubernetes.Clientset)
	// 获取config
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	ctx, cancel = context.WithTimeout(context.TODO(), time.Second*2)
	defer cancel()
	if name == "" {
		podsList, err := clientSet.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
		if err != nil {
			global.TPLogger.Error("获取pod失败：", err)
			return nil, errors.New("获取pod失败")
		}
		for _, item := range podsList.Items {
			fmt.Println(item.Name)
		}
		return podsList, nil
	} else {
		pod, err := clientSet.CoreV1().Pods(namespace).Get(ctx, name, metav1.GetOptions{})
		if err != nil {
			global.TPLogger.Error("获取pod失败：", err)
			return nil, errors.New("获取pod失败")
		}
		fmt.Println(pod.Name, pod.Namespace)
		return pod, nil
	}

}
