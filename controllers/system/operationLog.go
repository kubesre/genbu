/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package system

import (
	"genbu/common/global"
	service "genbu/service/system"
	"github.com/gin-gonic/gin"
)

func GetOperationLogList(ctx *gin.Context) {
	params := new(struct {
		Limit int `form:"limit"`
		Page  int `form:"page"`
	})
	if err := ctx.ShouldBindQuery(&params); err != nil {
		global.TPLogger.Error("查询操作日志列表数据绑定失败：", err)
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	data, err := service.NewOperationLogService().GetOperationLogList(params.Limit, params.Page)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	global.ReturnContext(ctx).Successful("success", data)
}
