package kubernetes

import (
	"genbu/service/kubernetes"

	"github.com/gin-gonic/gin"
)

func ListK8sConfigMap(ctx *gin.Context) {
	cid := ctx.Query("cid")
	kubernetes.NewK8sInterface().ListK8sConfig(cid)
}
