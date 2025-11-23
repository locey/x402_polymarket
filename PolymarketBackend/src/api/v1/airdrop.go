package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/locey/x402_polymarket/PolymarketBackend/src/service/svc"
	"github.com/locey/x402_polymarket/PolymarketBackend/src/service/v1"
	"github.com/locey/x402_polymarket/PolymarketBackend/src/types/v1"
	"net/http"
	"strconv"
	"time"
)

// CreateCampaignHandler 创建空投活动处理器
func CreateCampaignHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.CreateCampaignReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "请求参数错误: " + err.Error(),
				"data":    nil,
			})
			return
		}

		campaign, err := service.CreateCampaign(c.Request.Context(), svcCtx.DB, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "创建空投活动失败: " + err.Error(),
				"data":    nil,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "空投活动创建成功",
			"data":    campaign,
		})
	}
}

// GetActiveCampaignsHandler 获取活跃空投活动列表处理器
func GetActiveCampaignsHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		campaigns, err := service.GetActiveCampaigns(c.Request.Context(), svcCtx.DB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "获取活动列表失败: " + err.Error(),
				"data":    nil,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "获取成功",
			"data":    campaigns,
		})
	}
}

// GetCampaignDetailHandler 获取空投活动详情处理器
func GetCampaignDetailHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		campaignID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "无效的活动ID",
				"data":    nil,
			})
			return
		}

		// 这里调用服务层获取活动详情
		campaign, err := service.GetCampaignDetail(c.Request.Context(), svcCtx.DB, uint(campaignID))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "获取活动详情失败: " + err.Error(),
				"data":    nil,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "获取活动详情成功",
			"data":    campaign,
		})
	}
}

// UpdateCampaignStatusHandler 更新空投活动状态处理器
func UpdateCampaignStatusHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		campaignID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "无效的活动ID",
				"data":    nil,
			})
			return
		}

		var req struct {
			IsActive bool `json:"isActive" binding:"required"`
		}

		err = service.UpdateCampaignStatus(c.Request.Context(), svcCtx.DB, uint(campaignID), req.IsActive)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "更新活动状态失败: " + err.Error(),
				"data":    nil,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "更新活动状态成功",
			"data":    nil,
		})
	}
}

// GetCampaignStatsHandler 获取空投活动统计信息处理器
func GetCampaignStatsHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		campaignID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "无效的活动ID",
				"data":    nil,
			})
			return
		}

		stats, err := service.GetCampaignStats(c.Request.Context(), svcCtx.DB, uint(campaignID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "获取统计信息失败: " + err.Error(),
				"data":    nil,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "获取统计信息成功",
			"data":    stats,
		})
	}
}

// ClaimAirdropHandler 领取空投处理器
func ClaimAirdropHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ClaimAirdropReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "请求参数错误: " + err.Error(),
				"data":    nil,
			})
			return
		}

		claim, err := service.ClaimAirdrop(c.Request.Context(), svcCtx.DB, req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": err.Error(),
				"data":    nil,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "空投领取成功",
			"data":    claim,
		})
	}
}

// GetUserClaimsHandler 获取用户所有领取记录处理器
func GetUserClaimsHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		walletAddress := c.Param("address")

		claims, err := service.GetUserClaims(c.Request.Context(), svcCtx.DB, walletAddress)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "获取领取记录失败: " + err.Error(),
				"data":    nil,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "获取用户领取记录成功",
			"data": gin.H{
				"wallet_address": walletAddress,
				"claims":         claims,
				"count":          len(claims),
			},
		})
	}
}

// CheckEligibilityHandler 检查用户空投资格处理器
func CheckEligibilityHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		campaignID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "无效的活动ID",
				"data":    nil,
			})
			return
		}

		walletAddress := c.Param("address")

		// 这里调用服务层检查用户资格
		eligible, err := service.CheckEligibility(c.Request.Context(), svcCtx.DB, uint(campaignID), walletAddress)

		if err != nil {
			if err.Error() == "空投活动不存在" {
				c.JSON(http.StatusNotFound, gin.H{
					"code":    404,
					"message": err.Error(),
					"data":    nil,
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    500,
					"message": "检查资格失败: " + err.Error(),
					"data":    nil,
				})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "检查资格成功",
			"data":    eligible,
		})
	}
}

// DeleteCampaignHandler 删除空投活动处理器
func DeleteCampaignHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		campaignID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "无效的活动ID",
				"data":    nil,
			})
			return
		}

		var req types.DeleteCampaignReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "请求参数错误: " + err.Error(),
				"data":    nil,
			})
			return
		}

		// 调用服务层删除空投活动
		err = service.DeleteCampaign(c.Request.Context(), svcCtx.DB, uint(campaignID), req.ForceDelete)
		if err != nil {
			// 根据错误类型返回不同的HTTP状态码
			if err.Error() == "该活动已有用户领取记录，无法删除。如需强制删除，请设置force_delete=true" {
				c.JSON(http.StatusConflict, gin.H{
					"code":    409,
					"message": err.Error(),
					"data":    nil,
				})
			} else if err.Error() == "空投活动不存在" {
				c.JSON(http.StatusNotFound, gin.H{
					"code":    404,
					"message": err.Error(),
					"data":    nil,
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    500,
					"message": "删除空投活动失败: " + err.Error(),
					"data":    nil,
				})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "空投活动删除成功",
			"data": gin.H{
				"campaign_id":  campaignID,
				"force_delete": req.ForceDelete,
				"deleted_at":   time.Now(),
			},
		})
	}
}

// UpdateCampaignPartialHandler 部分更新空投活动处理器
func UpdateCampaignPartialHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		campaignID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "无效的活动ID",
				"data":    nil,
			})
			return
		}

		var req types.UpdateCampaignPartialReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "请求参数错误: " + err.Error(),
				"data":    nil,
			})
			return
		}

		// 调用服务层部分更新空投活动
		campaign, err := service.UpdateCampaignPartial(c.Request.Context(), svcCtx.DB, uint(campaignID), req)
		if err != nil {
			// 根据错误类型返回不同的HTTP状态码
			if err.Error() == "空投活动不存在" {
				c.JSON(http.StatusNotFound, gin.H{
					"code":    404,
					"message": err.Error(),
					"data":    nil,
				})
			} else if err.Error() == "该活动已有用户领取记录，无法修改代币数量" {
				c.JSON(http.StatusConflict, gin.H{
					"code":    409,
					"message": err.Error(),
					"data":    nil,
				})
			} else if err.Error() == "没有提供可更新的字段" {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    400,
					"message": err.Error(),
					"data":    nil,
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    500,
					"message": "更新空投活动失败: " + err.Error(),
					"data":    nil,
				})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "空投活动更新成功",
			"data":    campaign,
		})
	}
}

// GetCampaignClaimsHandler 获取活动所有领取记录处理器
func GetCampaignClaimsHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取路径参数
		campaignID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "无效的活动ID",
				"data":    nil,
			})
			return
		}

		// 获取查询参数
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
		status := c.Query("status") // 可选的状态筛选

		// 参数验证
		if page < 1 {
			page = 1
		}
		if limit < 1 || limit > 100 {
			limit = 20
		}

		// 调用服务层获取领取记录
		claims, pagination, err := service.GetCampaignClaims(
			c.Request.Context(),
			svcCtx.DB,
			uint(campaignID),
			page,
			limit,
			status,
		)

		if err != nil {
			if err.Error() == "空投活动不存在" {
				c.JSON(http.StatusNotFound, gin.H{
					"code":    404,
					"message": err.Error(),
					"data":    nil,
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    500,
					"message": "获取领取记录失败: " + err.Error(),
					"data":    nil,
				})
			}
			return
		}

		// 获取统计摘要
		summary, err := service.GetCampaignClaimsSummary(c.Request.Context(), svcCtx.DB, uint(campaignID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "获取统计信息失败: " + err.Error(),
				"data":    nil,
			})
			return
		}

		// 获取活动名称
		var campaignName string
		if len(claims) > 0 && claims[0].Campaign.Name != "" {
			campaignName = claims[0].Campaign.Name
		} else {
			// 如果没有记录或记录中没有活动名称，单独查询
			var campaign types.AirdropCampaign
			if err := svcCtx.DB.WithContext(c.Request.Context()).
				Select("name").First(&campaign, campaignID).Error; err == nil {
				campaignName = campaign.Name
			} else {
				campaignName = "未知活动"
			}
		}

		// 构建响应数据
		response := types.CampaignClaimsResponse{
			CampaignID:   uint(campaignID),
			CampaignName: campaignName,
			Claims:       claims,
			Summary:      *summary,
			Pagination:   *pagination,
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "获取活动领取记录成功",
			"data":    response,
		})
	}
}

// GetUserCampaignClaimHandler 获取用户在特定活动的领取记录处理器
func GetUserCampaignClaimHandler(svcCtx *svc.ServerCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取路径参数
		walletAddress := c.Param("address")
		campaignID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "无效的活动ID",
				"data":    nil,
			})
			return
		}

		// 参数验证
		if walletAddress == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "钱包地址不能为空",
				"data":    nil,
			})
			return
		}

		// 调用服务层获取用户活动领取记录
		claim, err := service.GetUserCampaignClaim(
			c.Request.Context(),
			svcCtx.DB,
			walletAddress,
			uint(campaignID),
		)

		if err != nil {
			if err.Error() == "空投活动不存在" {
				c.JSON(http.StatusNotFound, gin.H{
					"code":    404,
					"message": err.Error(),
					"data":    nil,
				})
			} else if err.Error() == "该用户在此活动中没有领取记录" {
				c.JSON(http.StatusNotFound, gin.H{
					"code":    404,
					"message": err.Error(),
					"data":    nil,
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    500,
					"message": "获取领取记录失败: " + err.Error(),
					"data":    nil,
				})
			}
			return
		}

		// 构建响应数据
		response := types.UserCampaignClaimResponse{
			WalletAddress: walletAddress,
			CampaignID:    uint(campaignID),
			CampaignName:  claim.Campaign.Name,
			Claim:         claim,
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "获取用户活动领取记录成功",
			"data":    response,
		})
	}
}
