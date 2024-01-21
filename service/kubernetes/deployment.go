/*
@auth: Ansu
@source: Ansu
@time: 2024/1/18
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

func GetDeploymentList(cid string, namespace string) (err error) {
	clientSetAny, found := global.ClientSetCache.Get(cid)
	if !found {
		global.TPLogger.Error("当前集群不存在，", err)
		return errors.New("当前集群不存在")
	}
	clientset := clientSetAny.(*kubernetes.Clientset)
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	ctx, cancel = context.WithTimeout(context.TODO(), time.Second*2)
	defer cancel()
	deploymentList, err := clientset.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		global.TPLogger.Error("获取命名空间失败", err)
		return errors.New("获取命名空间失败")
	}
	for _, item := range deploymentList.Items {
		fmt.Println(item.Name)
	}
	return nil
}
