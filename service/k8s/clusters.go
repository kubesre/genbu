/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/10
*/

package k8s

import (
	"errors"
	"genbu/common/global"
	dk "genbu/dao/k8s"
	"genbu/models/k8s"
	"genbu/utils"
)

// k8s 集群添加

func (k *k8sCluster) AddK8sCluster(cluster *k8s.Configs) (err error) {
	err = dk.NewK8sInterface().AddK8sCluster(cluster)
	if err != nil {
		global.TPLogger.Error("集群添加失败：", err)
		return errors.New("集群添加失败")

	}
	return nil
}

// 集群列表

func (k *k8sCluster) ListK8sCluster(name string, limit, page int) (clusters *k8s.ClusterK8sList, err error) {
	clusters, err = dk.NewK8sInterface().ListK8sCluster(name, limit, page)
	if err != nil {
		global.TPLogger.Error("获取集群列表失败：", err)
		return nil, errors.New("获取集群列表失败")
	}
	for _, item := range clusters.Items {
		configStr := item.Text
		decodeConfig, _ := utils.DecodeBase64(configStr)
		_, err = global.NewClientInterface().NewClientSet(decodeConfig)
		if err != nil {
			item.Active = "down"
		} else {
			item.Active = "running"
		}
		item.Text = ""
	}
	return clusters, nil
}

// TODO 查看config文件
