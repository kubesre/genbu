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
	Status   uint   `gorm:"type:tinyint(1);default:1;comment:'状态(正常/禁用, 默认正常)'" json:"status"`
	CreateBy string `gorm:"column:create_by;comment:'创建来源'" json:"create_by" form:"create_by"`
}

func (*Configs) TableName() string {
	return "k8s_config"
}

// kubernetes 集群列表

type ClusterK8s struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	CID      string `json:"cid" form:"cid"`
	Status   uint   `json:"status" form:"status"`
	Active   string `json:"active" form:"active"`
	CreateBy string `json:"create_by" form:"create_by"`
}

type ClusterK8sList struct {
	Items []*Configs `json:"items"`
	Total int64      `json:"total"`
}


type ConfigMap struct {}