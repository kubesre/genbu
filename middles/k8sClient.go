/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/10
*/

package middles

import (
	"genbu/common/global"
	"genbu/dao/k8s"
	"github.com/gin-gonic/gin"
	"strings"
)

// k8s client 中间件

func K8sClientCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求路径
		path := strings.Split(c.Request.RequestURI, "?")[0]
		if !strings.Contains(path, "cluster") {
			global.TPLogger.Error("没有指定请求集群ID")
			c.Abort()
			return
		}
		clusterID := c.Keys["cluster"]
		clusterIDStr, _ := clusterID.(string)
		if clusterID != nil && strings.Contains(path, clusterIDStr) {
			c.Next()
		}
		// 获取集群列表
		data, err := k8s.NewK8sInterface().ActiveK8sClusterList()
		if err != nil {
			global.TPLogger.Error("获取集群信息失败：", err)
			c.Abort()
			return
		}
		for _, item := range data {
			if strings.Contains(path, item.CID) {
				c.Set("cluster", item.CID)
				break
			} else {
				global.TPLogger.Error("没有指定请求集群ID")
				c.Abort()
				return
			}
		}
		global.TPLogger.Info("放行k8s中间件")
		c.Next()

	}
}
