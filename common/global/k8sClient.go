/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package global

import (
	"github.com/patrickmn/go-cache"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// kubernetes client

type ClientInterface interface {
	NewClientSet(configText string) (*kubernetes.Clientset, error)
}
type clientResource struct {
}

func NewClientInterface() ClientInterface {
	return &clientResource{}
}

func (c *clientResource) NewClientSet(configText string) (*kubernetes.Clientset, error) {
	cfg, err := clientcmd.NewClientConfigFromBytes([]byte(configText))
	if err != nil {
		return nil, err
	}
	restConfig, err := cfg.ClientConfig()
	if err != nil {
		return nil, err
	}

	// 初始化客户端
	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	version, err := clientSet.Discovery().ServerVersion()
	if err != nil {
		return nil, err
	}
	TPLogger.Info("集群连接成功！！！", version.String())

	return clientSet, nil
}

// 初始化缓存设置永不过期

var ClientSetCache *cache.Cache

func InitK8sClientCache() {
	ClientSetCache = cache.New(cache.NoExpiration, 0)
}
