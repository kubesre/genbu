/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/10
*/

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

// 获取node节点信息

func (k *k8sCluster) GetK8sClusterNodeList(cid string) (err error) {
	clientSetAny, found := global.ClientSetCache.Get(cid)
	if !found {
		global.TPLogger.Error("当前集群不存在：", err)
		return errors.New("当前集群不存在")
	}
	clientSet := clientSetAny.(*kubernetes.Clientset)
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
