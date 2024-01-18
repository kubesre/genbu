/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/10
*/

package kubernetes

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"genbu/common/global"
	dao "genbu/dao/kubernetes"
	"genbu/models/kubernetes"
	"genbu/utils"
	"github.com/patrickmn/go-cache"
	"time"
)

// kubernetes 集群添加

func (k *k8sCluster) AddK8sCluster(cluster *kubernetes.Configs) (err error) {
	localTime := time.Now().String()
	hash := md5.Sum([]byte(localTime))
	cluster.CID = hex.EncodeToString(hash[:])
	decodeConfig, _ := utils.DecodeBase64(cluster.Text)
	clientSet, err := global.NewClientInterface().NewClientSet(decodeConfig)
	if err != nil {
		global.TPLogger.Error("初始化clientSet失败：", err)
		return
	}
	// 获取集群版本信息
	version, _ := clientSet.ServerVersion()
	cluster.Version = version.String()
	err = dao.NewK8sInterface().AddK8sCluster(cluster)
	if err != nil {
		global.TPLogger.Error("集群添加失败：", err)
		return errors.New("集群添加失败")
	}
	// 添加缓存
	global.ClientSetCache.Set(cluster.CID, clientSet, cache.NoExpiration)
	return nil
}

// 集群列表

func (k *k8sCluster) ListK8sCluster(name string, limit, page int) (clusters *kubernetes.ClusterK8sList, err error) {
	clusters, err = dao.NewK8sInterface().ListK8sCluster(name, limit, page)
	if err != nil {
		global.TPLogger.Error("获取集群列表失败：", err)
		return nil, errors.New("获取集群列表失败")
	}
	return clusters, nil
}

// TODO 查看config文件

// 集群删除

func (k *k8sCluster) DeleteK8sCluster(cid []string) error {
	err := dao.NewK8sInterface().DeleteK8sCluster(cid)
	if err != nil {
		global.TPLogger.Error("集群删除失败：", err)
		return errors.New("集群删除失败")
	}
	for _, item := range cid {
		global.ClientSetCache.Delete(item)
	}
	return nil
}

// 集群更新

func (k *k8sCluster) UpdateK8sCluster(cluster *kubernetes.Configs) error {
	err := dao.NewK8sInterface().UpdateK8sCluster(cluster)
	if err != nil {
		global.TPLogger.Error("集群更新失败：", err)
		return errors.New("集群更新失败")
	}
	if cluster.Text != "" {
		global.ClientSetCache.Delete(cluster.CID)
		decodeConfig, _ := utils.DecodeBase64(cluster.Text)
		clientSet, err := global.NewClientInterface().NewClientSet(decodeConfig)
		if err != nil {
			global.TPLogger.Error("初始化clientSet失败：", err)
			return errors.New("初始化clientSet失败")
		}
		global.ClientSetCache.Set(cluster.CID, clientSet, cache.NoExpiration)
	}
	return nil
}

// 集群刷新

func (k *k8sCluster) RefreshK8sCluster() error {
	ClusterList, _ := dao.NewK8sInterface().ListK8sCluster("", 0, 0)
	for _, cluster := range ClusterList.Items {
		decodeConfig, _ := utils.DecodeBase64(cluster.Text)
		clientSet, err := global.NewClientInterface().NewClientSet(decodeConfig)
		if err != nil {
			global.TPLogger.Error("初始化clientSet失败：", err)
			cluster.Status = 0
			err = dao.NewK8sInterface().UpdateK8sCluster(cluster)
			return err
		}
		cluster.Status = 0
		version, _ := clientSet.ServerVersion()
		cluster.Version = version.String()
		cluster.Status = 0
		err = dao.NewK8sInterface().UpdateK8sCluster(cluster)
		return err
	}
	return nil
}
