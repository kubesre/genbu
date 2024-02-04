/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/10
*/

package kubernetes

import (
	"gorm.io/gorm"
)

// kubernetes config文件

type Configs struct {
	gorm.Model
	Name     string `json:"name" form:"name" gorm:"type:varchar(500);not null;unique;comment:'名称'"`
	CID      string `json:"cid" form:"cid" gorm:"type:varchar(500);not null;unique;comment:'集群ID'"`
	Text     string `json:"text" form:"text" gorm:"type:text;not null;comment:'config文件内容'"`
	Version  string `json:"version" form:"version" gorm:"type:text;not null;comment:'集群版本'"`
	Status   uint   `gorm:"type:tinyint(1);default:1;comment:'状态(正常/禁用, 默认正常)'" json:"status"`
	CreateBy string `gorm:"column:create_by;comment:'创建来源'" json:"create_by" form:"create_by"`
}

func (*Configs) TableName() string {
	return "k8s_config"
}

type ClusterK8sList struct {
	Items []*Configs `json:"items"`
	Total int64      `json:"total"`
}

type ConfigMap struct {
	Name      string `json:"name" form:"name" binding:"required"`
	NameSpace string `json:"namespace" form:"namespace"`
}

type ConfigMapDelete struct {
	ConfigMapName []map[string]string `json:"configmap_name" form:"name" binding:"required"`
	NameSpace     string              `json:"namespace" form:"namespace"`
}

type DeleteConfigMaps struct {
	NameSpace string `json:"namespace" form:"namespace"`
}

type CreateConfigMap struct {
	ConfigMapName string `json:"configmap_name" binding:"required"` // 这里的ConfigMapName表示的是通过kubectl get cm 查看出来的名称
	Text          string `json:"text" binding:"required"`           // ConfigMap类容
	NameSpace     string `json:"namespace"`                         // 命名空间的名称
}

type NameSpace struct {
	NameSpace string `json:"name" binding:"required"`
}
type UpdateNameSpace struct {
	NameSpaceName string `json:"name" binding:"required"`
	//Status        string `json:"status" `
	Annotations string `json:"annotations"`
	Labels      string `json:"labels"`
}
