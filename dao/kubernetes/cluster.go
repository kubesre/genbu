/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/10
*/

package kubernetes

import (
	"genbu/common/global"
	mod "genbu/models/kubernetes"
)

type InterfaceK8s interface {
	AddK8sCluster(cluster *mod.Configs) (err error)
	ListK8sCluster(name string, limit, page int) (clusters *mod.ClusterK8sList, err error)
	ActiveK8sCluster(cid string) (cluster *mod.Configs, err error)
	ActiveK8sClusterList() (clusters []*mod.Configs, err error)
	DeleteK8sCluster(cid []string) error
	UpdateK8sCluster(cluster *mod.Configs) error
}

type k8sCluster struct{}

func NewK8sInterface() InterfaceK8s {
	return &k8sCluster{}
}

// kubernetes 集群添加

func (k *k8sCluster) AddK8sCluster(cluster *mod.Configs) (err error) {
	if err = global.GORM.Model(&mod.Configs{}).Create(&cluster).Error; err != nil {
		return err
	}
	return nil
}

// 集群列表

func (k *k8sCluster) ListK8sCluster(name string, limit, page int) (clusters *mod.ClusterK8sList, err error) {
	// 定义分页起始位置
	startSet := (page - 1) * limit
	var (
		K8sClusterList []*mod.Configs
		total          int64
	)
	if err = global.GORM.Model(&mod.Configs{}).Where("name LIKE ?", "%"+name+"%").Count(&total).
		Limit(limit).Offset(startSet).Order("id desc").Find(&K8sClusterList).Error; err != nil {
		return nil, err
	}
	return &mod.ClusterK8sList{
		Items: K8sClusterList,
		Total: total,
	}, nil
}

// 集群可用配置信息详情

func (k *k8sCluster) ActiveK8sCluster(cid string) (cluster *mod.Configs, err error) {
	if err = global.GORM.Model(&mod.Configs{}).Where("c_id = ? AND status = ?", cid, 1).First(&cluster).Error; err != nil {
		return nil, err
	}
	return cluster, nil
}

// 集群可用配置信息列表

func (k *k8sCluster) ActiveK8sClusterList() (clusters []*mod.Configs, err error) {
	if err = global.GORM.Model(&mod.Configs{}).Where("status = ?", 1).Find(&clusters).Error; err != nil {
		return nil, err
	}
	return clusters, nil
}

// 集群删除

func (k *k8sCluster) DeleteK8sCluster(cid []string) error {
	if err := global.GORM.Delete(&mod.Configs{}, "c_id IN (?)", cid).Error; err != nil {
		return err
	}
	return nil
}

// 集群更新

func (k *k8sCluster) UpdateK8sCluster(cluster *mod.Configs) error {
	if err := global.GORM.Model(&mod.Configs{}).Where("id = ?", cluster.ID).Updates(&cluster).Error; err != nil {
		return err
	}
	return nil
}
