/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package system

import (
	"gorm.io/gorm"
	"time"
)

type OperationLog struct {
	gorm.Model
	Username   string    `gorm:"type:varchar(20);comment:'用户登录名'" json:"username"`
	Ip         string    `gorm:"type:varchar(20);comment:'Ip地址'" json:"ip"`
	IpLocation string    `gorm:"type:varchar(20);comment:'Ip所在地'" json:"ipLocation"`
	Method     string    `gorm:"type:varchar(20);comment:'请求方式'" json:"method"`
	Path       string    `gorm:"type:varchar(100);comment:'访问路径'" json:"path"`
	Remark     string    `gorm:"type:varchar(100);comment:'备注'" json:"remark"`
	Status     int       `gorm:"type:int(4);comment:'响应状态码'" json:"status"`
	StartTime  time.Time `gorm:"type:datetime(3);comment:'发起时间'" json:"startTime"`
	TimeCost   int64     `gorm:"type:int(6);comment:'请求耗时(ms)'" json:"timeCost"`
	UserAgent  string    `gorm:"type:varchar(20);comment:'浏览器标识'" json:"userAgent"`
}

func (*OperationLog) TableName() string {
	return "operation_log"
}

type OperationLogList struct {
	Items []OperationLog `json:"items"`
	Total int64          `json:"total"`
}

// IP归属地查询

type IPLocation struct {
	Data Data `json:"data"`
}
type Data struct {
	Continent string `json:"continent"`
	Owner     string `json:"owner"`
	Prov      string `json:"prov"`
	City      string `json:"city"`
	District  string `json:"district"`
}
