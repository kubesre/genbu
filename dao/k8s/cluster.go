/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/10
*/

package k8s

import (
	"genbu/common/global"
	"genbu/models/k8s"
)

type InterfaceK8s interface {
	AddK8sCluster(cluster *k8s.Configs) (err error)
	ListK8sCluster(name string, limit, page int) (clusters *k8s.ClusterK8sList, err error)
	ActiveK8sCluster(cid string) (cluster *k8s.Configs, err error)
	ActiveK8sClusterList() (clusters []*k8s.Configs, err error)
	DeleteK8sCluster(cid []string) error
	UpdateK8sCluster(cluster *k8s.Configs) error
}

type k8sCluster struct{}

func NewK8sInterface() InterfaceK8s {
	return &k8sCluster{}
}

// k8s 集群添加

func (k *k8sCluster) AddK8sCluster(cluster *k8s.Configs) (err error) {
	if err = global.GORM.Model(&k8s.Configs{}).Create(&cluster).Error; err != nil {
		return err
	}
	return nil
}

// 集群列表

func (k *k8sCluster) ListK8sCluster(name string, limit, page int) (clusters *k8s.ClusterK8sList, err error) {
	// 定义分页起始位置
	startSet := (page - 1) * limit
	var (
		K8sClusterList []*k8s.Configs
		total          int64
	)
	if err = global.GORM.Model(&k8s.Configs{}).Where("name LIKE ?", "%"+name+"%").Count(&total).
		Limit(limit).Offset(startSet).Order("id desc").Find(&K8sClusterList).Error; err != nil {
		return nil, err
	}
	return &k8s.ClusterK8sList{
		Items: K8sClusterList,
		Total: total,
	}, nil
}

// 集群可用配置信息详情

func (k *k8sCluster) ActiveK8sCluster(cid string) (cluster *k8s.Configs, err error) {
	if err = global.GORM.Model(&k8s.Configs{}).Where("c_id = ? AND status = ?", cid, 1).First(&cluster).Error; err != nil {
		return nil, err
	}
	return cluster, nil
}

// 集群可用配置信息列表

func (k *k8sCluster) ActiveK8sClusterList() (clusters []*k8s.Configs, err error) {
	if err = global.GORM.Model(&k8s.Configs{}).Where("status = ?", 1).Find(&clusters).Error; err != nil {
		return nil, err
	}
	return clusters, nil
}

// 集群删除

func (k *k8sCluster) DeleteK8sCluster(cid []string) error {
	if err := global.GORM.Delete(&k8s.Configs{}, "c_id IN (?)", cid).Error; err != nil {
		return err
	}
	return nil
}

// 集群更新

func (k *k8sCluster) UpdateK8sCluster(cluster *k8s.Configs) error {
	if err := global.GORM.Model(&k8s.Configs{}).Where("id = ?", cluster.ID).Updates(&cluster).Error; err != nil {
		return err
	}
	return nil
}
