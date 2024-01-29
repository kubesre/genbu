/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package kubernetes

import (
	"genbu/models/kubernetes"
	corev1 "k8s.io/api/core/v1"
)

type InterfaceK8s interface {
	AddK8sCluster(cluster *kubernetes.Configs) (err error)
	ListK8sCluster(name string, limit, page int) (clusters *kubernetes.ClusterK8sList, err error)
	GetK8sClusterNodeList(cid string, name string, page, limit int) (nodeResp *NodesResp, err error)
	GetK8sClusterNodeDetail(cid string, name string) (node *corev1.Node, err error)
	DeleteK8sCluster(id []string) error
	UpdateK8sCluster(cluster *kubernetes.Configs) error
	RefreshK8sCluster() error
}

type k8sCluster struct{}

func NewK8sInterface() InterfaceK8s {
	return &k8sCluster{}
}
