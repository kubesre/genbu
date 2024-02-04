package kubernetes

import (
	"genbu/common/global"
	mod "genbu/models/kubernetes"
	service "genbu/service/kubernetes"
	"github.com/gin-gonic/gin"
	"time"
)

// 获取命名空间
func GetK8sNameSpaceList(ctx *gin.Context) {
	clusterID := ctx.Param("cid")
	err := service.NewK8sInterface().GetK8sNameSpaceList(clusterID)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	//global.ReturnContext(ctx).Failed("failed", err.Error())
	global.ReturnContext(ctx).Successful("success", "success")
}

// 创建命名空间
func CreateK8sNameSpace(ctx *gin.Context) {
	cid := ctx.Param("cid")
	params := new(mod.CreateNameSpace)
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	ns, err := service.NewK8sInterface().CreateK8sNameSpace(cid, params.NameSpace)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
	} else {
		global.ReturnContext(ctx).Successful("success", map[string]interface{}{
			"name":        ns.Name,
			"create_time": ns.CreationTimestamp.Format(time.DateTime),
		})

	}

}
