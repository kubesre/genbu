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
	"genbu/utils"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/util/retry"
	"time"
)

type DeployResp struct {
	Item  []appsv1.Deployment `json:"items"`
	Total int                 `json:"total"`
}

func (k *k8sCluster) GetDeploymentList(cid, namespace string, page, limit int) (deployResp *DeployResp, err error) {
	clientSetAny, found := global.ClientSetCache.Get(cid)
	if !found {
		global.TPLogger.Error("当前集群不存在，", err)
		return nil, errors.New("当前集群不存在")
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
		global.TPLogger.Error("获取deployment失败", err)
		return nil, errors.New("获取deployment失败")
	}
	// 为 deployResp 分配内存并初始化
	deployResp = &DeployResp{}
	startIndex := (page - 1) * limit
	endIndex := page * limit
	if endIndex > len(deploymentList.Items) {
		endIndex = len(deploymentList.Items)
	}
	//暂存代码 暂时用不到
	//deployResp = deploymentList.Items[startIndex:endIndex]
	//var deployList []k8s.Deployment
	//for _, item := range deploymentList.Items[startIndex:endIndex] {
	//	fmt.Println(item.Name, item.Namespace, item.Spec.MinReadySeconds, item.Status.Replicas, item.Status.ReadyReplicas, item.CreationTimestamp)
	//
	//	deploy := &k8s.Deployment{
	//		Name:              item.Name,
	//		Namespace:         item.Namespace,
	//		MinReadySeconds:   item.Spec.MinReadySeconds,
	//		Replicas:          item.Status.Replicas,
	//		ReadyReplicas:     item.Status.ReadyReplicas,
	//		CreationTimestamp: k8s.Time(item.CreationTimestamp),
	//	}
	//	deployList = append(deployList, deploy)
	//}

	deployResp.Item = deploymentList.Items[startIndex:endIndex]
	deployResp.Total = len(deploymentList.Items)
	return deployResp, nil
}

func (k *k8sCluster) GetDeploymentDetails(cid, namespace, name string) (deploymentDetails *appsv1.Deployment, err error) {
	clientSetAny, fount := global.ClientSetCache.Get(cid)
	if !fount {
		global.TPLogger.Error("当前集群不存在", err)
		return nil, errors.New("当前集群不存在")
	}
	clientset := clientSetAny.(*kubernetes.Clientset)
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	ctx, cancel = context.WithTimeout(context.TODO(), time.Second*2)
	defer cancel()
	deploymentDetails, err = clientset.AppsV1().Deployments(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		global.TPLogger.Error("获取deployment失败", err)
		return nil, errors.New("获取deployment失败")
	}
	return deploymentDetails, err
}

// get deployment yaml
func (k *k8sCluster) GetDeploymentYaml(cid, namespace, name string) (str string, err error) {
	clientSetAny, fount := global.ClientSetCache.Get(cid)
	if !fount {
		global.TPLogger.Error("当前集群不存在", err)
		return "", errors.New("当前集群不存在")
	}
	clientset := clientSetAny.(*kubernetes.Clientset)
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	ctx, cancel = context.WithTimeout(context.TODO(), time.Second*2)
	defer cancel()
	deployment, err := clientset.AppsV1().Deployments(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		global.TPLogger.Error("获取deployment失败", err)
		return "", errors.New("获取deployment失败")
	}
	// 将 Deployment 对象序列化为 YAML
	schemeNew := runtime.NewScheme()
	codecs := json.NewSerializerWithOptions(json.DefaultMetaFactory, schemeNew, schemeNew, json.SerializerOptions{Yaml: true, Pretty: true})
	var yamlData []byte
	yamlData, err = runtime.Encode(codecs, deployment)
	if err != nil {
		return "", err
	}
	dest := utils.EncodeBase64(yamlData)
	return dest, nil
}

// delete Deployment
func (k *k8sCluster) DeleteDeployment(cid, namespace, name string) error {
	clientSetAny, fount := global.ClientSetCache.Get(cid)
	if !fount {
		global.TPLogger.Error("当前集群不存在")
		return errors.New("当前集群不存在")
	}
	clientset := clientSetAny.(*kubernetes.Clientset)
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	ctx, cancel = context.WithTimeout(context.TODO(), time.Second*2)
	defer cancel()
	err := clientset.AppsV1().Deployments(namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		global.TPLogger.Error("Deployment删除失败", err)
		return errors.New("Deployment删除失败")
	}
	return nil
}

// create or update Deployment2yaml
func (k *k8sCluster) CreateOrUpdateDeployment2Yaml(cid, text string) (string, error) {
	clientSetAny, fount := global.ClientSetCache.Get(cid)
	if !fount {
		global.TPLogger.Error("当前集群不存在")
		return "", errors.New("当前集群不存在")
	}
	clientset := clientSetAny.(*kubernetes.Clientset)
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	ctx, cancel = context.WithTimeout(context.TODO(), time.Second*2)
	defer cancel()
	dest, err := utils.DecodeBase64(text)
	if err != nil {
		global.TPLogger.Error("Base64解析失败")
		return "", errors.New("Base64解析失败")
	}
	// 创建包含 apps/v1 版本的 Deployment 的 Scheme
	err = scheme.AddToScheme(runtime.NewScheme())
	if err != nil {
		global.TPLogger.Error("数据转换失败")
		return "", errors.New("数据转换失败")
	}
	obj, _, err := scheme.Codecs.UniversalDeserializer().Decode([]byte(dest), nil, nil)
	if err != nil {
		global.TPLogger.Error("数据转换失败")
		return "", errors.New("数据转换失败")
	}
	deployment, ok := obj.(*appsv1.Deployment)
	if !ok {
		return "", fmt.Errorf("invalid object type: %T", obj)
	}
	// 获取 Deployment 的名称
	deploymentName := deployment.Name
	namespace := deployment.Namespace
	if deployment.Namespace == "" {
		namespace = "default"
	} else {
		namespace = deployment.Namespace
	}

	_, err = clientset.AppsV1().Deployments(namespace).Get(ctx, deploymentName, metav1.GetOptions{})
	if err != nil {
		//如果不存在，则创建
		_, err := clientset.AppsV1().Deployments(namespace).Create(ctx, deployment, metav1.CreateOptions{})
		if err != nil {
			global.TPLogger.Error("deployment创建失败")
			return "", errors.New("deployment创建失败")
		}
		fmt.Printf("Deployment %s created successfully \n", deploymentName)
		successMsg := fmt.Sprintf("Deployment %s created successfully \n", deploymentName)
		return successMsg, nil
	} else {
		// 如果deployment 已存在 ，则更新
		err := retry.RetryOnConflict(retry.DefaultBackoff, func() error {
			_, updateErr := clientset.AppsV1().Deployments(namespace).Update(ctx, deployment, metav1.UpdateOptions{})
			return updateErr
		})
		if err != nil {
			return "", err
		}
		fmt.Printf("Deployment %s updated successfully\n", deploymentName)
		successMsg := fmt.Sprintf("Deployment %s updated successfully\n", deploymentName)
		return successMsg, nil
	}
}

// create or update Deployment
func (k *k8sCluster) CreateOrUpdateDeployment2Arg(cid, namespace, deploymentName, containersName, image, lableKey, lableValue string, replicas int32) (string, error) {
	clientSetAny, fount := global.ClientSetCache.Get(cid)
	if !fount {
		global.TPLogger.Error("当前集群不存在")
		return "", errors.New("当前集群不存在")
	}
	clientset := clientSetAny.(*kubernetes.Clientset)
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	ctx, cancel = context.WithTimeout(context.TODO(), time.Second*2)
	defer cancel()
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: deploymentName,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{lableKey: lableValue},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{lableKey: lableValue},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  containersName,
							Image: image,
						},
					},
				},
			},
		},
	}
	// 获取已经存在的deployment
	_, err := clientset.AppsV1().Deployments(namespace).Get(ctx, deploymentName, metav1.GetOptions{})
	if err != nil {
		//如果不存在，则创建
		_, err := clientset.AppsV1().Deployments(namespace).Create(ctx, deployment, metav1.CreateOptions{})
		if err != nil {
			global.TPLogger.Error("deployment创建失败")
			return "", errors.New("deployment创建失败")
		}
		fmt.Printf("Deployment %s created successfully \n", deploymentName)
		successMsg := fmt.Sprintf("Deployment %s created successfully \n", deploymentName)
		return successMsg, nil
	} else {
		// 如果deployment 已存在 ，则更新
		err := retry.RetryOnConflict(retry.DefaultBackoff, func() error {
			_, updateErr := clientset.AppsV1().Deployments(namespace).Update(ctx, deployment, metav1.UpdateOptions{})
			return updateErr
		})
		if err != nil {
			return "", err
		}
		fmt.Printf("Deployment %s updated successfully\n", deploymentName)
		successMsg := fmt.Sprintf("Deployment %s updated successfully\n", deploymentName)
		return successMsg, nil
	}
}
