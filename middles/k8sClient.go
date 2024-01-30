/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/10
*/

package middles

import (
	"fmt"
	"genbu/common/global"
	"genbu/dao/kubernetes"
	"genbu/utils"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"strings"
)

// kubernetes client 中间件

func K8sClientCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求路径
		fmt.Println("缓存中间件路径：", c.Request.RequestURI)

		path := strings.Split(c.Request.RequestURI, "?")[0]
		if !strings.Contains(path, "cluster") {
			global.TPLogger.Error("请求失败")
			global.ReturnContext(c).Successful("failed", "请求路径有误")
			c.Abort()
			return
		}
		cidStr := c.Param("cid")
		// 查找缓存 key = cidStr
		_, found := global.ClientSetCache.Get(cidStr)
		if found {
			global.TPLogger.Info("该请求ID在缓存中")
			c.Next()
		}
		// 是否存在该cid
		cluster, err := kubernetes.NewK8sInterface().ActiveK8sCluster(cidStr)
		if err != nil {
			global.TPLogger.Error("集群获取失败：", err)
			global.ReturnContext(c).Successful("failed", "集群获取失败")
			c.Abort()
			return
		}
		// 将该集群加入到缓存中
		configStr := cluster.Text
		decodeConfig, _ := utils.DecodeBase64(configStr)
		clientSet, err := global.NewClientInterface().NewClientSet(decodeConfig)
		if err != nil {
			global.TPLogger.Error("集群初始化失败：", err)
			global.ReturnContext(c).Successful("failed", "集群初始化失败")
			c.Abort()
			return
		}
		global.ClientSetCache.Set(cidStr, clientSet, cache.NoExpiration)
		global.TPLogger.Info("放行k8s中间件")
		c.Next()

	}
}
