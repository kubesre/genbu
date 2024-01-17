/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package k8s

import (
	"genbu/models/k8s"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
)

type InterfaceK8s interface {
	AddK8sCluster(cluster *k8s.Configs) (err error)
	ListK8sCluster(name string, limit, page int) (clusters *k8s.ClusterK8sList, err error)
	InitClientSet(cid string) (clientSet *kubernetes.Clientset, err error)
	GetStatefulSetList(cid, namespace string) (statefulSetList *v1.StatefulSetList, err error)
	GetK8sClusterNodeList(cid string) (err error)
	DeleteK8sCluster(id []string) error
	UpdateK8sCluster(cluster *k8s.Configs) error
	RefreshK8sCluster() error
}

type k8sCluster struct{}

func NewK8sInterface() InterfaceK8s {
	return &k8sCluster{}
}
