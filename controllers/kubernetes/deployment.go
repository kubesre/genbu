/*
@auth: Ansu
@source: Ansu
@time: 2024/1/18
*/

package kubernetes

import (
	"genbu/common/global"
	mod "genbu/models/kubernetes"
	service "genbu/service/kubernetes"
	"github.com/gin-gonic/gin"
)

// get deployment list
func GetDeploymentList(ctx *gin.Context) {
	params := new(struct {
		Name  string `form:"name"`
		Page  int    `form:"page"`
		Limit int    `form:"limit"`
	})
	clusterID := ctx.Param("cid")
	if err := ctx.Bind(params); err != nil {
		global.ReturnContext(ctx).Failed("绑定参数失败", err.Error())
		return
	}
	data, err := service.NewK8sInterface().GetDeploymentList(clusterID, ctx.Query("namespace"), params.Page, params.Limit)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", data)

}

// get deployment details

func GetDeploymentDetails(ctx *gin.Context) {
	clusterID := ctx.Param("cid")
	data, err := service.NewK8sInterface().GetDeploymentDetails(clusterID, ctx.DefaultQuery("namespace", "default"), ctx.Query("name"))
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", data)

}

// 删除deployment
func DeleteDeployment(ctx *gin.Context) {
	clusterID := ctx.Param("cid")
	params := new(mod.DeleteDeployment)
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.TPLogger.Error("参数校验失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	err := service.NewK8sInterface().DeleteDeployment(clusterID, params.NameSpace, params.Name)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	global.ReturnContext(ctx).Successful("success", "删除成功")
}

// 获取deployment 的yaml
func GetDeploymentYaml(ctx *gin.Context) {
	clusterID := ctx.Param("cid")
	data, err := service.NewK8sInterface().GetDeploymentYaml(clusterID, ctx.DefaultQuery("namespace", "default"), ctx.Query("name"))
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", data)

}

// 根据yaml创建或更新deployment
func CreateOrUpdateDeployment2Yaml(ctx *gin.Context) {
	clusterID := ctx.Param("cid")
	params := new(mod.UpdateDeployment2Yaml)
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.TPLogger.Error("参数校验失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	data, err := service.NewK8sInterface().CreateOrUpdateDeployment2Yaml(clusterID, params.Text)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", data)

}

// 根据参数创建或更新deployment
func CreateOrUpdateDeployment2Arg(ctx *gin.Context) {
	clusterID := ctx.Param("cid")
	params := new(mod.UpdateDeployment2Arg)
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.TPLogger.Error("参数校验失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	data, err := service.NewK8sInterface().CreateOrUpdateDeployment2Arg(clusterID, params.NameSpace, params.ContainersName, params.Name, params.Image, params.LableKey, params.LableValue, params.Replicas)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", data)

}
