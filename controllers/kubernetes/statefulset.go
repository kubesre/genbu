package kubernetes

/*
@auth: Meersburg
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/2/02
*/
import (
	"genbu/common/global"
	sk "genbu/service/kubernetes"
	"github.com/gin-gonic/gin"
)

// yaml方式创建statefulSet
func CreateStatefulSetYaml(ctx *gin.Context) {
	cid := ctx.Param("cid")
	parms := new(struct {
		Content string `json:"content" form:"content"`
	})
	err := ctx.ShouldBind(&parms)
	if err != nil {
		global.TPLogger.Error("添加参数校验失败", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	err = sk.NewK8sInterface().CreateStatefulSetYaml(cid, parms.Content)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", "statefulSet服务创建成功")
}

// yaml方式更新statefulSet
func UpdateStatefulSetYaml(ctx *gin.Context) {
	cid := ctx.Param("cid")
	parms := new(struct {
		Content string `json:"content" form:"content"`
	})
	err := ctx.ShouldBind(&parms)
	if err != nil {
		global.TPLogger.Error("添加参数校验失败", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	err = sk.NewK8sInterface().UpdateStatefulSetYaml(cid, parms.Content)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", "statefulSet服务更新成功")
}

// statefulset列表
func ListstatefulSet(ctx *gin.Context) {
	cid := ctx.Param("cid")
	parms := new(struct {
		Namespace  string `form:"namespace"`
		FilterName string `form:"name"`
		Limit      int    `form:"limit"`
		Page       int    `form:"page"`
	})
	err := ctx.ShouldBindQuery(&parms)
	if err != nil {
		global.TPLogger.Error("添加参数校验失败", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	statefulsetList, err := sk.NewK8sInterface().GetStatefulSetList(cid, parms.Namespace, parms.FilterName, parms.Limit, parms.Page)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", statefulsetList)
}

// statefulSet详情
func DetailStatefulSet(ctx *gin.Context) {
	cid := ctx.Param("cid")
	parms := new(struct {
		Name      string `form:"name"`
		Namespace string `form:"namespace"`
	})
	err := ctx.ShouldBindQuery(&parms)
	if err != nil {
		global.TPLogger.Error("添加参数校验失败", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	statefulSet, err := sk.NewK8sInterface().GetStatefulSetDetail(cid, parms.Name, parms.Namespace)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", statefulSet)

}

// 删除statefulSet
func DelectStatefulSet(ctx *gin.Context) {
	cid := ctx.Param("cid")
	parms := new(struct {
		Name      string `form:"name"`
		Namespace string `form:"namespace"`
	})
	err := ctx.ShouldBindQuery(&parms)
	if err != nil {
		global.TPLogger.Error("添加参数校验失败", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	err = sk.NewK8sInterface().DeleteStatefulSet(cid, parms.Name, parms.Namespace)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", "删除statefulSet成功")
}
