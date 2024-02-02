package kubernetes

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"genbu/common/global"
	"io"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
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
		ctx     context.Context
		cancel  context.CancelFunc
		allPods []corev1.Pod
	)
	ctx, cancel = context.WithTimeout(context.TODO(), time.Second*2)
	defer cancel()

	listOptions := metav1.ListOptions{
		Limit: int64(pageSize),
	}

	for {
		pods, err := clientSet.CoreV1().Pods(namespace).List(ctx, listOptions)
		if err != nil {
			global.TPLogger.Error("获取pod失败：", err)
			return nil, errors.New("获取pod失败")
		}

		allPods = append(allPods, pods.Items...)

		if len(pods.Continue) == 0 || len(allPods) >= page*pageSize {
			break
		}

		listOptions.Continue = pods.Continue
	}

	startIndex := (page - 1) * pageSize
	endIndex := page * pageSize

	if startIndex >= len(allPods) {
		return []corev1.Pod{}, nil
	}

	if endIndex > len(allPods) {
		endIndex = len(allPods)
	}

	return allPods[startIndex:endIndex], nil
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
	return pod, nil

}

func (k *k8sCluster) CreatePod(cid string, pod interface{}) (ret interface{}, err error) {
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

	ret, err = clientSet.CoreV1().Pods(podPtr.Namespace).Create(ctx, podPtr, metav1.CreateOptions{})
	if err != nil {
		global.TPLogger.Error("创建pod失败：", err)
		return nil, errors.New("创建pod失败")
	}
	return ret, nil
}

func (k *k8sCluster) UpdatePod(cid string, patch interface{}) (ret interface{}, err error) {
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

	patchPtr, ok := patch.(*corev1.Pod)

	targetPodName := patchPtr.Name
	targetPodNamespace := patchPtr.Namespace

	if !ok {
		// 类型断言失败，处理错误情况
	}

	patchBytes, err := json.Marshal(patchPtr)
	if err != nil {
		global.TPLogger.Error("updatePod.JsonMarshal.error:", err)
		return nil, errors.New("更新pod失败")
	}
	ret, err = clientSet.CoreV1().Pods(targetPodNamespace).Patch(ctx, targetPodName, types.StrategicMergePatchType, patchBytes, metav1.PatchOptions{})
	if err != nil {
		global.TPLogger.Errorf("更新pod %s 失败：%v", targetPodName, err)
		return nil, errors.New("更新pod失败")
	}
	return ret, nil
}

func (k *k8sCluster) DeletePod(cid, namespace string, pods []string) (ret interface{}, err error) {
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
	var success_list []string
	for _, pod := range pods {
		err := clientSet.CoreV1().Pods(namespace).Delete(ctx, pod, metav1.DeleteOptions{})
		if err != nil {
			global.TPLogger.Errorf("删除pod %s 失败：%v", pod, err)
			return nil, errors.New("删除pod " + pod + " 失败,原因：" + err.Error())
		}
		global.TPLogger.Infof("删除Pod %s 成功。", pod)
		success_list = append(success_list, pod)

	}
	return fmt.Sprintf("删除Pod %s 完成", success_list), nil
}

func (k *k8sCluster) GetPodLogs(cid, namespace, name string, follow bool) (ret interface{}, err error) {
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

	// 创建日志流Watcher
	req := clientSet.CoreV1().Pods(namespace).GetLogs(name, &corev1.PodLogOptions{Follow: follow})
	podLogs, err := req.Stream(ctx)
	if err != nil {
		// 错误处理
	}
	defer podLogs.Close()

	// 读取日志内容
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		// 错误处理
	}
	// TODO logs -f 效果实现
	return buf.String(), nil
}

// TODO watch 效果实现
func (k *k8sCluster) WatchPod(cid string, namespace string) (pw watch.Interface, err error) {
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

	podWatch, err := clientSet.CoreV1().Pods(namespace).Watch(ctx, metav1.ListOptions{Watch: true})
	if err != nil {
		global.TPLogger.Error("获取pod失败：", err)
		return nil, errors.New("获取pod失败")
	}
	for {

		select {
		case event, _ := <-podWatch.ResultChan():
			pod, ok := event.Object.(*corev1.Pod)
			if !ok {
				continue
			}

			message := fmt.Sprintf("Pod %s [%s]: %s", pod.Name, event.Type, pod.Status.Phase)
			fmt.Println(message)
		}

	}
	return podWatch, nil

}

// TODO exec 效果实现
func (k *k8sCluster) ExecPod(cid, namespace, name, container, command string) (ret interface{}, err error) {
	clientSetAny, found := global.ClientSetCache.Get(cid)
	if !found {
		global.TPLogger.Error("当前集群不存在：", err)
		return nil, errors.New("当前集群不存在")
	}
	clientSet := clientSetAny.(*kubernetes.Clientset)
	// 获取config
	// var (
	// 	ctx    context.Context
	// 	cancel context.CancelFunc
	// )
	// ctx, cancel = context.WithTimeout(context.TODO(), time.Second*2)
	// defer cancel()
	// 创建 pod的执行请求
	request := clientSet.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(name).
		Namespace(namespace).
		SubResource("exec")

	// 声明执行请求的选项
	options := &corev1.PodExecOptions{
		Container: container,
		Command:   []string{"/bin/sh", "-c", command},
		Stdin:     false,
		Stdout:    true,
		Stderr:    true,
		TTY:       false,
	}

	// 在请求中设置选项
	request.VersionedParams(options, scheme.ParameterCodec)
	// 执行请求并获取执行结果
	exec, err := remotecommand.NewSPDYExecutor(getConfig(), "POST", request.URL())
	if err != nil {
		global.TPLogger.Error(err)
		return
	}

	// 创建输出流
	output := &bytes.Buffer{}

	// 执行命令
	err = exec.Stream(remotecommand.StreamOptions{
		Stdout: output,
		Stderr: output,
		Tty:    false,
	})
	if err != nil {
		global.TPLogger.Error(err)
		return
	}

	// 返回执行结果
	return output.String(), nil

}

func getConfig() *rest.Config {
	config, _ := rest.InClusterConfig()
	return config
}
