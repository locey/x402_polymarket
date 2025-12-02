package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/locey/x402_polymarket/PolymarketBackend/src/api/v1"
	"github.com/locey/x402_polymarket/PolymarketBackend/src/service/svc"
	"net/http"
)

func loadV1(r *gin.Engine, svcCtx *svc.ServerCtx) {
	apiV1 := r.Group("/api/v1")

	user := apiV1.Group("/user")
	{
		user.GET("/:address/login-message", v1.GetLoginMessageHandler(svcCtx)) // 生成login签名信息
		user.POST("/login", v1.UserLoginHandler(svcCtx))                       // 登陆
		user.GET("/:address/sig-status", v1.GetSigStatusHandler(svcCtx))       // 获取用户签名状态
	}

	// 空投相关路由
	airdrop := apiV1.Group("/airdrop")
	{
		// 空投活动管理
		airdrop.POST("/campaigns", v1.CreateCampaignHandler(svcCtx))                 // 创建空投活动
		airdrop.GET("/campaigns", v1.GetActiveCampaignsHandler(svcCtx))              // 获取活跃空投活动列表
		airdrop.GET("/campaigns/:id", v1.GetCampaignDetailHandler(svcCtx))           // 获取空投活动详情
		airdrop.PUT("/campaigns/:id/status", v1.UpdateCampaignStatusHandler(svcCtx)) // 更新空投活动状态
		airdrop.GET("/campaigns/:id/stats", v1.GetCampaignStatsHandler(svcCtx))      // 获取空投活动统计信息
		airdrop.DELETE("/campaigns/:id", v1.DeleteCampaignHandler(svcCtx))           // 删除空投活动
		airdrop.PATCH("/campaigns/:id", v1.UpdateCampaignPartialHandler(svcCtx))     // 部分更新空投活动

		// 空投领取相关
		airdrop.POST("/campaigns/claim", v1.ClaimAirdropHandler(svcCtx))                     // 领取空投
		airdrop.GET("/claims/:address", v1.GetUserClaimsHandler(svcCtx))                     // 获取用户所有领取记录
		airdrop.GET("/campaigns/:id/claims", v1.GetCampaignClaimsHandler(svcCtx))            // 获取活动所有领取记录
		airdrop.GET("/claims/:address/campaign/:id", v1.GetUserCampaignClaimHandler(svcCtx)) // 获取用户在特定活动的领取记录

		// 资格验证相关
		airdrop.GET("/campaigns/:id/eligibility/:address", v1.CheckEligibilityHandler(svcCtx)) // 检查用户空投资格
	}

	// 健康检查端点
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "polymarket",
		})
	})
}
