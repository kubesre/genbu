package k8s

import (
	"context"
	"errors"
	"genbu/common/global"
	"genbu/dao/k8s"
	"genbu/utils"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

/*var Ks ks

type ks struct {
	ClientSet *kubernetes.Clientset
}

func (k *ks) Init(cid string) (err error) {
	cluster, err := k8s.NewK8sInterface().ActiveK8sCluster(cid)
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
	k.ClientSet, err = global.NewClientInterface().NewClientSet(decodeConfig)
	if err != nil {
		global.TPLogger.Error("初始化client失败：", err)
		return errors.New("初始化client失败")
	}
	return nil
}*/

func (k *k8sCluster) InitClientSet(cid string) (clientSet *kubernetes.Clientset, err error) {
	cluster, err := k8s.NewK8sInterface().ActiveK8sCluster(cid)
	if err != nil {
		global.TPLogger.Error("获取可用集群信息失败: ", err)
		return nil, errors.New("获取可用集群信息失败")
	}
	configStr := cluster.Text
	decodeConfig, err := utils.DecodeBase64(configStr)
	if err != nil {
		global.TPLogger.Error("集群config文件加载失败: ", err)
		return nil, errors.New("集群config文件加载失败")
	}
	clientSet, err = global.NewClientInterface().NewClientSet(decodeConfig)
	if err != nil {
		global.TPLogger.Error("初始化client失败：", err)
		return nil, errors.New("初始化client失败")
	}
	return clientSet, nil
}

//获取statefulset列表
func (k *k8sCluster) GetStatefulSetList(cid, namespace string) (statefulSetList *v1.StatefulSetList, err error) {
	clientSet, err := k.InitClientSet(cid)
	statefulSetList, err = clientSet.AppsV1().StatefulSets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		global.TPLogger.Error("获取statefulSet列表失败：", err)
		return nil, errors.New("获取statefulSet列表失败")
	}

	return statefulSetList, nil
}
