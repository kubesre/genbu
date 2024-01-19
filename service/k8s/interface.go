/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package k8s

import "genbu/models/k8s"

type InterfaceK8s interface {
	AddK8sCluster(cluster *k8s.Configs) (err error)
	ListK8sCluster(name string, limit, page int) (clusters *k8s.ClusterK8sList, err error)
	GetK8sClusterNodeList(cid interface{}) (err error)
	GetK8sClusterPodList(cid interface{}, namespace, name string, page, pagesize int) (err error)
}

type k8sCluster struct{}

func NewK8sInterface() InterfaceK8s {
	return &k8sCluster{}
}
