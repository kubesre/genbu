package kubernetes

/*
@auth: 啷个办
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/17
*/
import (
	"genbu/common/global"
	mod "genbu/models/kubernetes"
	"genbu/service/kubernetes"
	"genbu/utils"
	"github.com/gin-gonic/gin"
	"time"
)

// 创建ConfigMap
func CreateConfigMap(ctx *gin.Context) {
	cid := ctx.Param("cid")
	params := new(mod.CreateConfigMap)
	var name_space string
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}

	if params.NameSpace == "" {
		name_space = "default"
	} else {
		name_space = params.NameSpace
	}

	// 解析前段传送的ConfigMap.yaml数据，该数据前段通过base64加密
	dest, err := utils.DecodeBase64(params.Text)

	if err != nil {
		global.TPLogger.Error("base64解析错误：", err)
		global.ReturnContext(ctx).Failed("failed", "数据解析失败")

	} else {

		cm, err := kubernetes.NewK8sInterface().CreateConfigMap(cid, name_space, params.ConfigMapName, dest)
		if err != nil {
			global.ReturnContext(ctx).Failed("failed", err.Error())
		} else {
			global.ReturnContext(ctx).Successful("success", map[string]interface{}{
				"name":        cm.Name,
				"namespace":   cm.Namespace,
				"data":        cm.Data,
				"create_time": cm.CreationTimestamp.Format(time.DateTime),
			})

		}
	}

}

// ConfigMap列表
func ListK8sConfigMap(ctx *gin.Context) {
	cid := ctx.Param("cid")
	namespace := ctx.DefaultQuery("namespace", "default")
	res, err := kubernetes.NewK8sInterface().ListK8sConfig(cid, namespace)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", res)
}

// 获取指定ConfigMap信息
func Get8sConfigMapInfo(ctx *gin.Context) {
	params := new(mod.ConfigMap)
	var name_space string
	cid := ctx.Param("cid")
	if err := ctx.ShouldBind(&params); err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	if params.NameSpace == "" {
		name_space = "default"
	} else {
		name_space = params.NameSpace
	}
	// 这里的命名空间不能为空
	res, err := kubernetes.NewK8sInterface().GetK8sConfigInfo(cid, name_space, params.Name)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", res)
}

// 删除指定ConfigMap，传入的是一个切片类型
func DeleteConfig(ctx *gin.Context) {
	params := new(mod.ConfigMapDelete)
	var name_space string
	cid := ctx.Param("cid")
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.ReturnContext(ctx).Failed("failed", "请传入必传参数")
		return
	}
	if params.NameSpace == "" {
		name_space = "default"
	} else {
		name_space = params.NameSpace
	}
	s, err := kubernetes.NewK8sInterface().DeleteConfig(cid, name_space, params.ConfigMapName)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", s)
}

// 删除多个ConfigMap
func DeleteConfigs(ctx *gin.Context) {
	var name_space string
	params := new(mod.DeleteConfigMaps)

	cid := ctx.Param("cid")
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.ReturnContext(ctx).Failed("failed", "请传入必传参数")
		return
	}

	// 判断是否传入namespace
	if params.NameSpace == "" {
		name_space = "default"
	} else {
		name_space = params.NameSpace
	}

	s, err := kubernetes.NewK8sInterface().DeleteConfigs(cid, name_space)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", s)
}

// 更新ConfigMap
func UpdateConfigMap(ctx *gin.Context) {
	cid := ctx.Param("cid")
	params := new(mod.CreateConfigMap)
	var name_space string
	if err := ctx.ShouldBindJSON(&params); err != nil {
		global.ReturnContext(ctx).Failed("failed", "请传入必传参数")
		return
	}
	// 判断是否传入namespace
	if params.NameSpace == "" {
		name_space = "default"
	} else {
		name_space = params.NameSpace
	}
	dest, err := utils.DecodeBase64(params.Text)
	if err != nil {
		global.TPLogger.Error("Base64解析失败: ", err)
		global.ReturnContext(ctx).Failed("failed", "配置文件解析失败")
		return
	}
	configMap, err := kubernetes.NewK8sInterface().UpdateConfigMap(cid, name_space, params.ConfigMapName, dest)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", configMap)
}
