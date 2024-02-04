package kubernetes

import (
	"context"
	"encoding/json"
	"errors"
	"genbu/common/global"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"time"
)

// 获取ns信息
func (k *k8sCluster) GetK8sNameSpaceList(cid string) (restful []map[string]interface{}, err error) {
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
	//nodeList, err := clientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	nameSpaceList, err := clientSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {

		global.TPLogger.Error("获取namespace失败：", err)
		return nil, errors.New("获取namespace失败")
	}
	//for _, item := range nameSpaceList.Items {
	//	fmt.Println(item.Name)
	//}
	restful = []map[string]interface{}{}
	for _, item := range nameSpaceList.Items {
		restful = append(restful, map[string]interface{}{
			"name":            item.Name,
			"api_version":     item.APIVersion,
			"resourceVersion": item.ResourceVersion,
			"kind":            item.Kind,
			//"namespace":       item.Namespace,
			"create_time": item.CreationTimestamp.Format(time.DateTime), // 格式化成：2006-01-02 15:04:05
			"labels":      item.Labels,
			"status":      item.Status.Phase,
		})
	}
	return restful, nil

}

// 新增ns
func (k *k8sCluster) CreateK8sNameSpace(cid, Namespace string) (*corev1.Namespace, error) {
	clientSetAny, found := global.ClientSetCache.Get(cid)
	if !found {
		global.TPLogger.Error("当前集群不存在：", errors.New(""))
		return nil, errors.New("当前集群不存在")
	}
	clientSet := clientSetAny.(*kubernetes.Clientset)
	namespace := corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: Namespace}}
	ns, err := clientSet.CoreV1().Namespaces().Create(context.Background(), &namespace, metav1.CreateOptions{})
	if err != nil {
		global.TPLogger.Error("创建namespace失败：", err)
		return nil, err
	} else {

		return ns, nil
	}

}

// 删除ns
func (k *k8sCluster) DeleteNamespace(cid, namespace string) (string, error) {
	clientSetAny, found := global.ClientSetCache.Get(cid)
	if !found {
		global.TPLogger.Error("当前集群不存在：", errors.New(""))
		return "", errors.New("当前集群不存在")
	}
	clientSet := clientSetAny.(*kubernetes.Clientset)
	err := clientSet.CoreV1().Namespaces().Delete(context.Background(), namespace, metav1.DeleteOptions{})
	if err != nil {
		global.TPLogger.Error("删除Namespace:%s失败：", namespace, err)
		return "", errors.New("Namespace删除失败")
	}

	return "namespace删除成功", nil
}

// 更新ns
func (k *k8sCluster) UpdateNamespace(cid, NameSpaceName, Annotations, Labels string) (*corev1.Namespace, error) {
	clientSetAny, found := global.ClientSetCache.Get(cid)

	if !found {
		global.TPLogger.Error("当前集群不存在：", errors.New(""))
		return nil, errors.New("当前集群不存在")
	}
	clientSet := clientSetAny.(*kubernetes.Clientset)
	namespace := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: NameSpaceName,
		},
	}
	//if Status != "" {
	//	namespace.Status.Phase = corev1.NamespacePhase(Status)
	//}
	if Annotations != "" {
		if namespace.Annotations == nil {
			namespace.Annotations = make(map[string]string)
		}
		namespace.Annotations["annotations"] = Annotations
	}
	if Labels != "" {
		if namespace.Labels == nil {
			namespace.Labels = make(map[string]string)
		}
		namespace.Labels["labels"] = Labels
	}
	_, err := json.Marshal(namespace)
	if err != nil {
		global.TPLogger.Error("解析失败", err.Error())
		return nil, errors.New("解析失败")
	}
	update, err := clientSet.CoreV1().Namespaces().Update(context.Background(), namespace, metav1.UpdateOptions{})
	if err != nil {
		global.TPLogger.Error("namespace更新失败: ", err)
		return nil, errors.New("namespace更新失败")
	} else {
		return update, nil
	}
}
