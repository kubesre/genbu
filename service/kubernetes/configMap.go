package kubernetes

import "fmt"

func (k *k8sCluster) ListK8sConfig(cid string) {
	// 这里加载集群中的configmap数据
	fmt.Printf("cid: %v\n", cid)
}
