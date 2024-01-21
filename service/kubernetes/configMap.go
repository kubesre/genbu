package kubernetes

/*
@auth: 啷个办
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/17
*/
import (
	"context"
	"errors"
	"genbu/common/global"
	"genbu/utils"
	"strings"
	"time"

	core_v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *k8sCluster) ListK8sConfig(cid, NameSpace string) (restful []map[string]interface{}, err error) {
	// 这里加载集群中的configmap数据
	c, err := utils.GetCache(cid)
	if err != nil {
		global.TPLogger.Error("当前集群不存在：", errors.New(""))
		return nil, errors.New("当前集群不存在")
	} else {
		cml, err := c.CoreV1().ConfigMaps(NameSpace).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			global.TPLogger.Error("获取ConfigMaps失败：", err)
			return nil, errors.New("获取ConfigMaps失败")
		}
		restful = []map[string]interface{}{}
		for _, item := range cml.Items {
			restful = append(restful, map[string]interface{}{
				"name":        item.Name,
				"api_version": item.APIVersion,
				"data":        item.Data,
				"kind":        item.Kind,
				"namespace":   item.Namespace,
				"create_time": item.CreationTimestamp.Format(time.DateTime), // 格式化成：2006-01-02 15:04:05
				"labels":      item.Labels,
			})
		}
		return restful, nil
	}
}

// 获取某个ConfigMap详情内容
func (k *k8sCluster) GetK8sConfigInfo(cid, NameSpace, Name string) (*core_v1.ConfigMap, error) {
	c, err := utils.GetCache(cid)
	if err != nil {
		global.TPLogger.Error("当前集群不存在：", errors.New(""))
		// return nil, errors.New("当前集群不存在")
		return nil, err
	}
	cm, err := c.CoreV1().ConfigMaps(NameSpace).Get(context.Background(), Name, v1.GetOptions{})
	if err != nil {
		global.TPLogger.Error("获取ConfigMaps失败：", err)
		return nil, err
	}
	return cm, nil
}

// 删除指定的ConfigMap
func (k *k8sCluster) DeleteConfig(cid, NameSpace string, ConfigMapName []map[string]string) (string, error) {
	c, err := utils.GetCache(cid)
	if err != nil {
		global.TPLogger.Error("当前集群不存在：", errors.New(""))
		return "", errors.New("集群不存在")
	}
	for _, item := range ConfigMapName {
		err = c.CoreV1().ConfigMaps(NameSpace).Delete(context.Background(), item["name"], v1.DeleteOptions{})
		if err != nil {
			global.TPLogger.Error("删除ConfigMap失败：", err)
			return "", errors.New("ConfigMap删除失败")
			break
		}
	}
	return "删除成功", nil
}

// 删除多个ConfigMap
func (k *k8sCluster) DeleteConfigs(cid, NameSpace string) (string, error) {
	c, err := utils.GetCache(cid)
	if err != nil {
		global.TPLogger.Error("当前集群不存在：", errors.New(""))
		return "", errors.New("集群不存在")
	}
	// err = c.CoreV1().ConfigMaps(NameSpace).DeleteCollection(context.Background(), metav1.DeleteOptions{}, metav1.ListOptions{})
	// if err != nil {
	// 	global.TPLogger.Error("删除ConfigMaps失败：", err)
	// 	return "", errors.New("ConfigMap删除失败")
	// } else {
	// 	return "删除成功", nil
	// }
	cml, err := c.CoreV1().ConfigMaps(NameSpace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		global.TPLogger.Error("删除ConfigMaps失败：", err)
		return "", errors.New("ConfigMaps删除失败")
	} else {
		for _, item := range cml.Items {
			err = c.CoreV1().ConfigMaps(NameSpace).Delete(context.Background(), item.Name, v1.DeleteOptions{})
			if err != nil {
				global.TPLogger.Error("删除ConfigMaps失败：", err)
				return "", errors.New("ConfigMaps删除失败")
			}
		}
		return "删除成功", nil
	}
}

func (k *k8sCluster) CreateConfigMap(cid, NameSpace, ConfigMapName, Text string) (*core_v1.ConfigMap, error) {
	c, err := utils.GetCache(cid)
	if err != nil {
		global.TPLogger.Error("当前集群不存在：", errors.New(""))
		return nil, errors.New("集群不存在")
	} else {
		configmap_data_name := strings.Replace(ConfigMapName, "-", ".", 1) // 更具用户提供的配置文件名称，生成ConfigMap的Data字段名称
		//fmt.Println("configmap_data_name:", configmap_data_name)
		config_map := core_v1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name: ConfigMapName,
			},
			Data: map[string]string{
				configmap_data_name: Text,
			},
		}

		cm, err := c.CoreV1().ConfigMaps(NameSpace).Create(context.Background(), &config_map, metav1.CreateOptions{})
		if err != nil {
			global.TPLogger.Error("创建ConfigMap失败：", err)
			return nil, err
		} else {

			return cm, nil
		}
	}
}

func (k *k8sCluster) UpdateConfigMap(cid, NameSpace, ConfigMapName, NewText string) (*core_v1.ConfigMap, error) {
	c, err := utils.GetCache(cid)
	if err != nil {
		global.TPLogger.Error("当前集群不存在：", errors.New(""))
		return nil, errors.New("集群不存在")
	} else {
		configmap_data_name := strings.Replace(ConfigMapName, "-", ".", 1) // 更具用户提供的配置文件名称，生成ConfigMap的Data字段名称
		configmap := core_v1.ConfigMap{
			ObjectMeta: v1.ObjectMeta{
				Name: ConfigMapName,
			},
			Data: map[string]string{
				configmap_data_name: NewText,
			},
		}
		update, err := c.CoreV1().ConfigMaps(NameSpace).Update(context.Background(), &configmap, metav1.UpdateOptions{})
		if err != nil {
			global.TPLogger.Error("ConfigMap更新失败: ", err)
			return nil, errors.New("ConfigMap更新失败")
		} else {
			return update, nil
		}

	}

}
