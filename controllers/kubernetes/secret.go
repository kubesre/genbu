/*
@auth: 啷个办
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/17
*/

package kubernetes

import (
	"fmt"
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
		global.ReturnContext(ctx).Failed("failed", err)
	} else {
		global.ReturnContext(ctx).Successful("success", secret)
	}
}

/*获取集群上所有的Secret*/
func ListK8sSecret(ctx *gin.Context) {
	cid := ctx.Param("cid")

	namespace := ctx.DefaultQuery("namespace", "default")

	secret, err := kubernetes.NewK8sInterface().ListK8sSecret(cid, namespace)
	if err != nil {
		global.TPLogger.Error("Secret获取失败：", err)
		global.ReturnContext(ctx).Failed("failed", err)
	} else {
		global.ReturnContext(ctx).Successful("success", secret)
	}
}

/*指定获取某个Secret*/
func GetK8sSecret(ctx *gin.Context) {
	cid := ctx.Param("cid")
	params := new(mod.Secret)
	var name_space string
	if err := ctx.ShouldBind(&params); err != nil {
		global.ReturnContext(ctx).Failed("failed", "请传入必传参数")
		return
	}
	if params.NameSpace == "" {
		name_space = "default"
	} else {
		name_space = params.NameSpace
	}
	secret, err := kubernetes.NewK8sInterface().GetK8sSecret(cid, name_space, params.Name)
	if err != nil {
		global.TPLogger.Error("Secret获取失败：", err)
		global.ReturnContext(ctx).Failed("failed", err)
	} else {
		global.ReturnContext(ctx).Successful("success", secret)
	}
}

/*删除指定secret*/
func DeleteSecret(ctx *gin.Context) {
	cid := ctx.Param("cid")
	params := new(mod.DeleteSecret)
	var name_space string
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.ReturnContext(ctx).Failed("failed", "请传入必传参数")
		return
	}
	if params.NameSpace == "" {
		name_space = "default"
	} else {
		name_space = params.NameSpace
	}
	err := kubernetes.NewK8sInterface().DeleteK8sSecret(cid, name_space, params.SecretName)
	if err != nil {
		global.TPLogger.Error("Secret删除失败：", err)
		global.ReturnContext(ctx).Failed("failed", err)
	} else {
		global.ReturnContext(ctx).Successful("success", "删除成功")
	}
}

// 删除指定命名空间下的所有Secret
func DeleteSecrets(ctx *gin.Context) {
	params := new(mod.DeleteSecrets)
	cid := ctx.Param("cid")
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.ReturnContext(ctx).Failed("failed", "请传入必传参数")
		return
	}
	var name_space string

	if params.NameSpace == "" {
		name_space = "default"
	} else {
		name_space = params.NameSpace
	}
	err := kubernetes.NewK8sInterface().DeleteK8sSecrets(cid, name_space)
	if err != nil {
		global.TPLogger.Error(fmt.Sprintf("%s命令空间下的Secret删除失败: %s\n", name_space, err.Error()))
		global.ReturnContext(ctx).Failed("failed", err)
	} else {
		global.ReturnContext(ctx).Successful("success", err)
	}
}

// 更新Secret
func UpdataSecret(ctx *gin.Context) {
	cid := ctx.Param("cid")

	params := new(mod.Update)
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.ReturnContext(ctx).Failed("failed", "请传入必传参数")
		return
	}
	var name_space string

	if params.NameSpace == "" {
		name_space = "default"
	} else {
		name_space = params.NameSpace
	}
	secret, err := kubernetes.NewK8sInterface().UpdateK8sSecret(cid, name_space, params.SecretName, params.Text)
	if err != nil {
		global.TPLogger.Error("Secret更新失败：", err)
		global.ReturnContext(ctx).Failed("failed", err)
	} else {
		global.ReturnContext(ctx).Successful("success", secret)
	}

}
