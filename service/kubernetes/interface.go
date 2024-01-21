/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package kubernetes

// /api/k8s/cluster/{cid}/node/listNode
import (
	"genbu/models/kubernetes"

	v1 "k8s.io/api/core/v1"
)

type InterfaceK8s interface {
	AddK8sCluster(cluster *kubernetes.Configs) (err error)
	ListK8sCluster(name string, limit, page int) (clusters *kubernetes.ClusterK8sList, err error)
	GetK8sClusterNodeList(cid string) (err error)
	DeleteK8sCluster(id []string) error
	UpdateK8sCluster(cluster *kubernetes.Configs) error
	RefreshK8sCluster() error
	ListK8sConfig(cid, NameSpace string) ([]map[string]interface{}, error)
	GetK8sConfigInfo(cid, NameSpace, Name string) (*v1.ConfigMap, error)
	DeleteConfig(cid, NameSpace string, ConfigMapName []map[string]string) (string, error)
	DeleteConfigs(cid, NameSpace string) (string, error)
	CreateConfigMap(cid, NameSpace, ConfigMapName, Text string) (*v1.ConfigMap, error)
	UpdateConfigMap(cid, NameSpace, ConfigMapName, Text string) (*v1.ConfigMap, error)
	ListK8sSecret(cid, NameSpace string) (*v1.SecretList, error)
	CreateK8sSecret(cid, NameSpace, SecretName, Text string) (*v1.Secret, error)
}

type k8sCluster struct{}

func NewK8sInterface() InterfaceK8s {
	return &k8sCluster{}
}
