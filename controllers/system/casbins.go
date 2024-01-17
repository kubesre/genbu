/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package system

import (
	"genbu/common/global"
	mod "genbu/models/system"
	service "genbu/service/system"
	"github.com/gin-gonic/gin"
)

func AddCasbin(ctx *gin.Context) {
	params := new(struct {
		Policy []*mod.CasbinPolicy `json:"policy"`
	})
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.TPLogger.Error("添加授权参数校验失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	err := service.NewCasbinInterface().AddPolicy(params.Policy)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	if err = global.CasbinEnforcer.LoadPolicy(); err != nil {
		global.TPLogger.Error("加载权限失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", "添加权限成功")

}

// 删除授权

func DelPolicy(ctx *gin.Context) {
	params := new(struct {
		Policy []*mod.CasbinPolicy `json:"policy"`
	})
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.TPLogger.Error("删除授权参数校验失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	if err := service.NewCasbinInterface().DelPolicy(params.Policy); err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	if err := global.CasbinEnforcer.LoadPolicy(); err != nil {
		global.TPLogger.Error("加载权限失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", "删除授权成功")

}

// 查看授权

func ListPolicy(ctx *gin.Context) {
	params := new(struct {
		keyWord string `form:"keyword"`
		Limit   int    `form:"limit"`
		Page    int    `form:"page"`
	})
	if err := ctx.ShouldBindQuery(&params); err != nil {
		global.TPLogger.Error("用户查询数据绑定失败：", err)
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	data := service.NewCasbinInterface().ListPolicy(params.keyWord, params.Limit, params.Page)
	global.ReturnContext(ctx).Successful("success", data)
}
