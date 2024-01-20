package kubernetes

import (
	"context"
	"errors"
	"fmt"
	"genbu/common/global"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"time"
)

/*
   @Auth: Menah3m
   @CreateTime: 2024/1/19
   @Desc: pods 相关操作
*/

func (k *k8sCluster) GetPodList(cid string, namespace string, page, pageSize int) (ret interface{}, err error) {
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

	podsList, err := clientSet.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		global.TPLogger.Error("获取pod失败：", err)
		return nil, errors.New("获取pod失败")
	}
	for _, item := range podsList.Items {
		fmt.Println(item.Name)
	}
	return podsList, nil

}

func (k *k8sCluster) GetPod(cid string, namespace, name string) (ret interface{}, err error) {
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

	pod, err := clientSet.CoreV1().Pods(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		global.TPLogger.Error("获取pod失败：", err)
		return nil, errors.New("获取pod失败")
	}
	fmt.Println(pod.Name, pod.Namespace)
	return pod, nil

}

func (i *k8sCluster) CreatePod(cid string, pod interface{}) (ret interface{}, err error) {
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

	podPtr, ok := pod.(*corev1.Pod)
	if !ok {
		// 类型断言失败，处理错误情况
	}
	fmt.Println(podPtr.Namespace)
	ret, err = clientSet.CoreV1().Pods(podPtr.Namespace).Create(ctx, podPtr, metav1.CreateOptions{})
	if err != nil {
		global.TPLogger.Error("创建pod失败：", err)
		return nil, errors.New("创建pod失败")
	}
	return ret, nil
}
