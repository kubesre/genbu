/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package system

import (
	"encoding/json"
	"fmt"
	"genbu/common/global"
	mod "genbu/models/system"
	"genbu/utils"
	"github.com/spf13/viper"
	"time"
)

type OperationLogService interface {
	SaveOperationLogChannel(olc <-chan *mod.OperationLog)
	GetOperationLogList(limit, page int) (*mod.OperationLogList, error)
}
type operationLogService struct{}

func NewOperationLogService() OperationLogService {
	return &operationLogService{}
}

//处理OperationLogChan将日志记录到数据库

func (s *operationLogService) SaveOperationLogChannel(olc <-chan *mod.OperationLog) {
	var url = viper.GetString("ipLocation.siteURL")
	var header = map[string]string{}
	Logs := make([]mod.OperationLog, 0)
	//5s 自动同步一次
	duration := 5 * time.Second
	timer := time.NewTimer(duration)
	defer timer.Stop()
	for {
		select {
		case log := <-olc:
			var ipLocation mod.IPLocation
			fullURL := fmt.Sprintf("%s?ip=%s", url, log.Ip)
			data, err := utils.DoRequest("GET", fullURL, header, nil)
			if err != nil {
				log.IpLocation = "查找失败"
			}
			_ = json.Unmarshal([]byte(data), &ipLocation)
			if ipLocation.Data.City == "" {
				log.IpLocation = ipLocation.Data.Continent
			} else {
				log.IpLocation = ipLocation.Data.City
			}
			Logs = append(Logs, *log)
			// 每5条记录到数据库
			if len(Logs) > 5 {
				global.GORM.Create(&Logs)
				Logs = make([]mod.OperationLog, 0)
				timer.Reset(duration) // 入库重置定时器
			}
		case <-timer.C: //5s 自动同步一次
			if len(Logs) > 0 {
				global.GORM.Create(&Logs)
				Logs = make([]mod.OperationLog, 0)
			}
			timer.Reset(duration) // 入库重置定时器
		}
	}
}

// 获取操作日志列表

func (s *operationLogService) GetOperationLogList(limit, page int) (*mod.OperationLogList, error) {
	// 定义分页起始位置
	startSet := (page - 1) * limit
	var (
		operationLogList []mod.OperationLog
		total            int64
	)
	if err := global.GORM.Model(&mod.OperationLog{}).Count(&total).Limit(limit).Offset(startSet).Order("start_time desc").
		Find(&operationLogList).Error; err != nil {
		return nil, err
	}
	return &mod.OperationLogList{
		Items: operationLogList,
		Total: total,
	}, nil
}
