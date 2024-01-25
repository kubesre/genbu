package kubernetes

import (
	"context"
	"errors"
	"genbu/common/global"
	"genbu/utils"

	core_v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 获取所有Secret
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
		}
		list, err := c.CoreV1().Secrets(NameSpace).Create(context.Background(), &secret_config, metav1.CreateOptions{})
		if err != nil {
			return nil, err
		} else {
			return list, nil
		}
	}
}

// 获取指定Secret
func (k *k8sCluster) GetK8sSecret(cid, NameSpace, Name string) (*core_v1.Secret, error) {
	c, err := utils.GetCache(cid)
	if err != nil {
		global.TPLogger.Error("当前集群不存在：", errors.New(""))
		return nil, errors.New("当前集群不存在")
	} else {
		get, err := c.CoreV1().Secrets(NameSpace).Get(context.Background(), Name, metav1.GetOptions{})
		if err != nil {
			return nil, err
		} else {
			return get, nil
		}
	}
}

// 删除Secret
func (k *k8sCluster) DeleteK8sSecret(cid, NameSpace string, ConfigMapName []map[string]string) error {
	c, err := utils.GetCache(cid)
	if err != nil {
		global.TPLogger.Error("当前集群不存在：", errors.New(""))
		return errors.New("当前集群不存在")
	} else {
		//
		for _, item := range ConfigMapName {
			err := c.CoreV1().Secrets(NameSpace).Delete(context.Background(), item["name"], metav1.DeleteOptions{})
			if err != nil {
				global.TPLogger.Error("Secret删除失败：", err)
				return errors.New("Secret删除失败")
				break
			}
		}
		return nil
	}
}

// 删除指定命名空间下的所有Secret
func (k *k8sCluster) DeleteK8sSecrets(cid, NameSpace string) error {
	c, err := utils.GetCache(cid)
	if err != nil {
		return err
	} else {

		err := c.CoreV1().Secrets(NameSpace).DeleteCollection(context.Background(), metav1.DeleteOptions{}, metav1.ListOptions{})
		if err != nil {
			global.TPLogger.Error(err)
			return errors.New("该namespace下没有Secret")
		} else {
			return nil
		}
	}
}

// 更新
func (k *k8sCluster) UpdateK8sSecret(cid, NameSpace, SecretName, Text string) (*core_v1.Secret, error) {
	c, err := utils.GetCache(cid)
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
		}
		list, err := c.CoreV1().Secrets(NameSpace).Update(context.Background(), &secret_config, metav1.UpdateOptions{})
		if err != nil {
			return nil, err
		} else {
			return list, nil
		}
	}
}
