/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2023/12/4
*/

package routers

import (
	"fmt"
	"genbu/common/global"
	"genbu/middles"
	v1 "genbu/routers/v1"
	"github.com/gin-gonic/gin"
	"time"
)

func BaseRouters() *gin.Engine {
	r := gin.New()
	// 自定义日志格式
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("[%s | method: %s | path: %s | host: %s | proto: %s | code: %d | %s | %s ]\n",
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.ClientIP,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
		)
	}))
	// 开启全部跨域允许
	r.Use(middles.Cors())
	authMiddleware, err := middles.InitAuth()
	if err != nil {
		global.TPLogger.Error("初始化JWT中间件失败: ", err)
		panic(fmt.Sprintf("初始化JWT中间件失败：%v", err))
	}
	// 健康检查
	r.GET("/health", func(ctx *gin.Context) {
		global.ReturnContext(ctx).Successful("success", "success")
		return
	})
	// 不需要做鉴权的接口 PublicGroup
	PublicGroup := r.Group("/api/v1")
	{
		v1.InitBaseRouters(PublicGroup, authMiddleware)
	}
	// 需要做鉴权的接口
	PrivateGroup := r.Group("/api/v1")
	// 鉴权
	//PrivateGroup.Use(gin.Recovery()).
	//	Use(middles.OperationLog()).Use(middles.CasbinMiddle())
	{
		v1.InitUserRouters(PrivateGroup, authMiddleware)
		v1.InitRolesRouters(PrivateGroup, authMiddleware)
		v1.InitDeptRouters(PrivateGroup, authMiddleware)
		v1.InitMenusRouters(PrivateGroup, authMiddleware)
		v1.InitPolicyRouters(PrivateGroup, authMiddleware)
		v1.InitLogRouters(PrivateGroup, authMiddleware)
		v1.InitK8sRouters(PrivateGroup, authMiddleware)
	}
	r.NoRoute(func(ctx *gin.Context) {
		global.ReturnContext(ctx).Failed("fail", "该接口未开放")
		return
	})
	return r
}
