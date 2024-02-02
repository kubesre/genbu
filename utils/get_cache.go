package utils

import (
	"errors"
	"genbu/common/global"
	"k8s.io/client-go/kubernetes"
)

func GetCache(cid string) (*kubernetes.Clientset, error) {
	clientSetAny, found := global.ClientSetCache.Get(cid)
	if !found {
		// global.TPLogger.Error("当前集群不存在：", errors.New(""))
		return nil, errors.New("当前集群不存在")
	}
	c := clientSetAny.(*kubernetes.Clientset)
	return c, nil
}
