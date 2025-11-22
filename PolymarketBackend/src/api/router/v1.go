package router

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/locey/x402_polymarket/PolymarketBackend/src/api/v1"
	"github.com/locey/x402_polymarket/PolymarketBackend/src/service/svc"
)

func loadV1(r *gin.Engine, svcCtx *svc.ServerCtx) {
	apiV1 := r.Group("/api/v1")

	user := apiV1.Group("/user")
	{
		user.GET("/:address/login-message", v1.GetLoginMessageHandler(svcCtx)) // 生成login签名信息
		user.POST("/login", v1.UserLoginHandler(svcCtx))                       // 登陆
		user.GET("/:address/sig-status", v1.GetSigStatusHandler(svcCtx))       // 获取用户签名状态
	}

}
