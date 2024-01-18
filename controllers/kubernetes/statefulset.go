package k8s

import (
	"genbu/common/global"
	sk "genbu/service/k8s"
	"github.com/gin-gonic/gin"
)

//statefulset列表
func ListstatefulSet(ctx *gin.Context) {
	cid := ctx.Param("cid")
	parms := new(struct {
		namespace string
	})
	err := ctx.ShouldBindQuery(&parms)
	if err != nil {
		global.TPLogger.Error("获取statefulset列表失败", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	statefulsetList, err := sk.NewK8sInterface().GetStatefulSetList(cid, parms.namespace)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", statefulsetList)
}
