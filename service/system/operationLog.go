/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package system

import (
	"errors"
	"genbu/common/global"
	dao "genbu/dao/system"
	mod "genbu/models/system"
)

type OperationLogService interface {
	GetOperationLogList(limit, page int) (*mod.OperationLogList, error)
}
type operationLogService struct{}

func NewOperationLogService() OperationLogService {
	return &operationLogService{}
}

func (s *operationLogService) GetOperationLogList(limit, page int) (*mod.OperationLogList, error) {
	data, err := dao.NewOperationLogService().GetOperationLogList(limit, page)
	if err != nil {
		global.TPLogger.Error("获取操作日志列表失败：", err)
		return nil, errors.New("获取操作日志列表")
	}
	return data, nil
}
