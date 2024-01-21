package kubernetes

import (
	"context"
	"errors"
	"genbu/common/global"
	"genbu/utils"
	core_v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *k8sCluster) ListK8sSecret(cid, NameSpace string) (*core_v1.SecretList, error) {
	c, err := utils.GetCache(cid)
	if err != nil {
		global.TPLogger.Error("当前集群不存在：", errors.New(""))
		return nil, errors.New("当前集群不存在")
	} else {
		list, err := c.CoreV1().Secrets(NameSpace).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			return nil, err
		} else {
			return list, nil
		}
	}
}

func (k *k8sCluster) CreateK8sSecret(cid, NameSpace, SecretName, Text string) (*core_v1.Secret, error) {
	c, err := utils.GetCache(cid)
	//var types any
	//if Type == "" {
	//	types = core_v1.SecretTypeOpaque
	//} else {
	//	types = Type
	//}
	if err != nil {
		global.TPLogger.Error("当前集群不存在：", errors.New(""))
		return nil, errors.New("当前集群不存在")
	} else {
		dest, err := utils.DecodeBase64(Text)
		if err != nil {
			global.TPLogger.Error("Base64解析失败: ", err)
			return nil, err
		}
		secret_config := core_v1.Secret{

			ObjectMeta: metav1.ObjectMeta{
				Name: SecretName,
			},

			Data: map[string][]byte{

				"token": []byte(dest),
			},
			//Type: types.(core_v1.SecretType),
		}
		list, err := c.CoreV1().Secrets(NameSpace).Create(context.Background(), &secret_config, metav1.CreateOptions{})
		if err != nil {
			return nil, err
		} else {
			return list, nil
		}
	}
}
