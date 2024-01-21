/*
@auth: 啷个办
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/17
*/

package kubernetes

import (
	"genbu/common/global"
	mod "genbu/models/kubernetes"
	"genbu/service/kubernetes"
	"github.com/gin-gonic/gin"
)

/*创建secret*/
func CreateK8sSecret(ctx *gin.Context) {
	cid := ctx.Param("cid")
	var name_space string

	params := new(mod.Create)
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.ReturnContext(ctx).Failed("failed", "请传入必传参数")
		return
	}
	if params.NameSpace == "" {
		name_space = "default"
	} else {
		name_space = params.NameSpace
	}
	secret, err := kubernetes.NewK8sInterface().CreateK8sSecret(cid, name_space, params.SecretName, params.Text)
	if err != nil {
		global.TPLogger.Error("Secret创建失败：", err)
		global.ReturnContext(ctx).Failed("failed", "Secret创建失败")
	} else {
		global.ReturnContext(ctx).Successful("success", secret)
	}
}

func ListK8sSecret(ctx *gin.Context) {
	cid := ctx.Param("cid")

	namespace := ctx.DefaultQuery("namespace", "default")

	secret, err := kubernetes.NewK8sInterface().ListK8sSecret(cid, namespace)
	if err != nil {
		global.TPLogger.Error("Secret获取失败：", err)
		global.ReturnContext(ctx).Failed("failed", "Secret获取失败")
	} else {
		global.ReturnContext(ctx).Successful("success", secret)
	}
}
