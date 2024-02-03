/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/10
*/

package kubernetes

import (
	"context"
	"errors"
	"genbu/common/global"
	"genbu/utils"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

// 获取node节点信息

type NodesResp struct {
	Items []corev1.Node `json:"items"`
	Total int           `json:"total"`
}

func (k *k8sCluster) GetK8sClusterNodeList(cid string, name string, page, limit int) (nodeResp *NodesResp, err error) {
	clientSet, err := utils.GetCache(cid)
	if err != nil {
		global.TPLogger.Error("Node管理模块获取节点:", errors.New("当前集群不存在"))
		// return nil, errors.New("当前集群不存在")
		return nil, err
	}
	// 获取config
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	ctx, cancel = context.WithTimeout(context.TODO(), time.Second*2)
	defer cancel()
	nodeList, err := clientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		global.TPLogger.Error("获取node失败：", err)
		return nil, errors.New("获取node失败")
	}

	selectableData := &DataSelector{
		GenericDataList: k.toCells(nodeList.Items),
		DataSelect: &DataSelectQuery{
			Filter: &FilterQuery{Name: name},
			Paginatite: &PaginateQuery{
				Limit: limit,
				Page:  page,
			},
		},
	}

	filtered := selectableData.Filter()
	total := len(filtered.GenericDataList)
	data := filtered.Sort().Paginate()
	nodes := k.fromCells(data.GenericDataList)

	return &NodesResp{
		Items: nodes,
		Total: total,
	}, nil
}

func (k *k8sCluster) GetK8sClusterNodeDetail(cid string, name string) (node *corev1.Node, err error) {
	clientSet, err := utils.GetCache(cid)
	if err != nil {
		global.TPLogger.Error("Node管理模块节点详情:", errors.New("当前集群不存在"))
		// return nil, errors.New("当前集群不存在")
		return nil, err
	}
	// 获取config
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	ctx, cancel = context.WithTimeout(context.TODO(), time.Second*2)
	defer cancel()

	node, err = clientSet.CoreV1().Nodes().Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		global.TPLogger.Error("获取node详情失败：", err)
		return nil, errors.New("获取node详情失败")
	}

	return node, nil
}

// 把corev1 node转成DataCell
func (k *k8sCluster) toCells(std []corev1.Node) []DataCell {
	cells := make([]DataCell, len(std))
	for i := range std {
		cells[i] = nodeCell(std[i])
	}
	return cells
}

// 把nodeCell转成corev1 node
func (k *k8sCluster) fromCells(cells []DataCell) []corev1.Node {
	nodes := make([]corev1.Node, len(cells))
	for i := range cells {
		nodes[i] = corev1.Node(cells[i].(nodeCell))
	}
	return nodes
}
