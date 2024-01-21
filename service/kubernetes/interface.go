/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package kubernetes

import "genbu/models/kubernetes"

type InterfaceK8s interface {
	AddK8sCluster(cluster *kubernetes.Configs) (err error)
	ListK8sCluster(name string, limit, page int) (clusters *kubernetes.ClusterK8sList, err error)
	GetK8sClusterNodeList(cid string) (err error)
	GetPodList(cid string, namespace string, page, pageSize int) (ret interface{}, err error)
	GetPod(cid string, namespace, name string) (ret interface{}, err error)
	CreatePod(cid string, pod interface{}) (ret interface{}, err error)
	UpdatePod(cid string, patch interface{}) (ret interface{}, err error)
	DeletePod(cid, namespace string, pods []string) (ret interface{}, err error)
	DeleteK8sCluster(id []string) error
	UpdateK8sCluster(cluster *kubernetes.Configs) error
	RefreshK8sCluster() error
}

type k8sCluster struct{}

func NewK8sInterface() InterfaceK8s {
	return &k8sCluster{}
}
