/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/13
*/

package k8s

import (
	"errors"
	"genbu/common/global"
	"genbu/dao/k8s"
	"genbu/utils"
	"github.com/patrickmn/go-cache"
)

// 初始化所有集群clientSet，并将其加入缓存

func InitAllClient() error {
	// 清空缓存
	global.ClientSetCache.Flush()
	configList, err := k8s.NewK8sInterface().ActiveK8sClusterList()
	if err != nil {
		global.TPLogger.Error("获取k8s集群列表失败：", err)
		return errors.New("获取k8s集群列表失败")
	}
	for _, item := range configList {
		configStr := item.Text
		cid := item.CID
		decodeConfig, _ := utils.DecodeBase64(configStr)
		clientSet, err := global.NewClientInterface().NewClientSet(decodeConfig)
		if err != nil {
			global.TPLogger.Error("初始化clientSet失败：", err)
			return errors.New("初始化clientSet失败")
		}
		// 设置缓存
		global.ClientSetCache.Set(cid, clientSet, cache.NoExpiration)
	}
	return nil
}
