package types

import (
	"github.com/shopspring/decimal"
	"time"
)

// CreateCampaignReq 创建空投活动的请求参数结构体
type CreateCampaignReq struct {
	Name        string          `json:"name" binding:"required"`         // 活动名称（必需）
	Description string          `json:"description"`                     // 活动描述（可选）
	TokenSymbol string          `json:"token_symbol" binding:"required"` // 代币符号（必需）
	TokenAmount decimal.Decimal `json:"token_amount" binding:"required"` // 每个用户获得的代币数量（必需）
	MaxClaims   int             `json:"max_claims"`                      // 最大领取人数，0表示无限制（可选）
	StartTime   time.Time       `json:"start_time" binding:"required"`   // 活动开始时间（必需）
	EndTime     time.Time       `json:"end_time" binding:"required"`     // 活动结束时间（必需）
}

// ClaimAirdropReq 领取空投的请求参数结构体
type ClaimAirdropReq struct {
	CampaignID    uint   `json:"campaignId" binding:"required"`    // 空投活动ID（必需）
	WalletAddress string `json:"walletAddress" binding:"required"` // 用户钱包地址（必需）
}

// CampaignStats 空投活动统计信息响应结构体
type CampaignStats struct {
	CampaignID   uint            `json:"campaign_id"`   // 活动ID
	CampaignName string          `json:"campaign_name"` // 活动名称
	TotalClaims  int             `json:"total_claims"`  // 总领取人数
	MaxClaims    int             `json:"max_claims"`    // 最大可领取人数
	Remaining    int             `json:"remaining"`     // 剩余可领取数量
	TokenSymbol  string          `json:"token_symbol"`  // 代币符号
	TokenAmount  decimal.Decimal `json:"token_amount"`  // 每个用户获得的代币数量
	IsActive     bool            `json:"is_active"`     // 活动是否激活
	ClaimRate    float64         `json:"claim_rate"`    // 领取率（0-1之间的小数）
	StartTime    time.Time       `json:"start_time"`    // 活动开始时间
	EndTime      time.Time       `json:"end_time"`      // 活动结束时间
}

// AirdropCampaign 空投活动数据库模型
// 对应数据库中的 airdrop_campaigns 表
type AirdropCampaign struct {
	ID          uint            `gorm:"primarykey" json:"id"`                             // 主键ID
	Name        string          `gorm:"size:200;not null" json:"name"`                    // 活动名称
	Description string          `gorm:"type:text" json:"description"`                     // 活动描述
	TokenSymbol string          `gorm:"size:20;not null" json:"token_symbol"`             // 代币符号
	TokenAmount decimal.Decimal `gorm:"type:decimal(36,18);not null" json:"token_amount"` // 每个用户获得的代币数量
	IsActive    bool            `gorm:"default:true" json:"is_active"`                    // 是否激活状态
	MaxClaims   int             `gorm:"default:0" json:"max_claims"`                      // 最大领取人数，0表示无限制
	StartTime   time.Time       `json:"start_time"`                                       // 活动开始时间
	EndTime     time.Time       `json:"end_time"`                                         // 活动结束时间
	TotalClaims int             `gorm:"default:0" json:"total_claims"`                    // 总领取次数
	CreatedAt   time.Time       `json:"created_at"`                                       // 创建时间
	UpdatedAt   time.Time       `json:"updated_at"`                                       // 更新时间
}

// AirdropClaim 空投领取记录数据库模型
// 对应数据库中的 airdrop_claims 表
type AirdropClaim struct {
	ID                uint            `gorm:"primarykey" json:"id"`                                   // 主键ID
	AirdropCampaignID uint            `gorm:"index;not null" json:"airdrop_campaign_id"`              // 关联的空投活动ID
	WalletAddress     string          `gorm:"size:42;not null" json:"wallet_address"`                 // 用户钱包地址
	ClaimAmount       decimal.Decimal `gorm:"type:decimal(36,18);not null" json:"claim_amount"`       // 实际领取数量
	ClaimStatus       string          `gorm:"size:20;default:pending" json:"claim_status"`            // 领取状态: pending, success, failed
	ClaimTxHash       string          `gorm:"size:66" json:"claim_tx_hash"`                           // 区块链交易哈希（可选）
	ClaimedAt         *time.Time      `json:"claimed_at"`                                             // 领取时间
	CreatedAt         time.Time       `json:"created_at"`                                             // 创建时间
	UpdatedAt         time.Time       `json:"updated_at"`                                             // 更新时间
	Campaign          AirdropCampaign `gorm:"foreignKey:AirdropCampaignID" json:"campaign,omitempty"` // 关联的空投活动信息
}

// DeleteCampaignReq 删除空投活动的请求参数结构体
type DeleteCampaignReq struct {
	ForceDelete bool `json:"force_delete"` // 是否强制删除（即使有领取记录也删除）
}

// UpdateCampaignPartialReq 部分更新空投活动的请求参数结构体
type UpdateCampaignPartialReq struct {
	Name        *string          `json:"name"`        // 活动名称（可选）
	Description *string          `json:"description"` // 活动描述（可选）
	TokenSymbol *string          `json:"tokenSymbol"` // 代币符号（可选）
	TokenAmount *decimal.Decimal `json:"tokenAmount"` // 每个用户获得的代币数量（可选）
	MaxClaims   *int             `json:"maxClaims"`   // 最大领取人数（可选）
	StartTime   *time.Time       `json:"startTime"`   // 活动开始时间（可选）
	EndTime     *time.Time       `json:"endTime"`     // 活动结束时间（可选）
	IsActive    *bool            `json:"isActive"`    // 是否激活（可选）
}

// Pagination 分页信息
type Pagination struct {
	Page      int   `json:"page"`      // 当前页码
	Limit     int   `json:"limit"`     // 每页数量
	Total     int64 `json:"total"`     // 总记录数
	TotalPage int   `json:"totalPage"` // 总页数
}

// ClaimsSummary 领取记录统计摘要
type ClaimsSummary struct {
	TotalCount   int             `json:"totalCount"`   // 总领取次数
	SuccessCount int             `json:"successCount"` // 成功次数
	PendingCount int             `json:"pendingCount"` // 待处理次数
	FailedCount  int             `json:"failedCount"`  // 失败次数
	TotalAmount  decimal.Decimal `json:"totalAmount"`  // 总领取金额
}

// CampaignClaimsResponse 活动领取记录响应
type CampaignClaimsResponse struct {
	CampaignID   uint           `json:"campaignId"`   // 活动ID
	CampaignName string         `json:"campaignName"` // 活动名称
	Claims       []AirdropClaim `json:"claims"`       // 领取记录列表
	Summary      ClaimsSummary  `json:"summary"`      // 统计摘要
	Pagination   Pagination     `json:"pagination"`   // 分页信息
}

// UserCampaignClaimResponse 用户活动领取记录响应
type UserCampaignClaimResponse struct {
	WalletAddress string        `json:"walletAddress"` // 钱包地址
	CampaignID    uint          `json:"campaignId"`    // 活动ID
	CampaignName  string        `json:"campaignName"`  // 活动名称
	Claim         *AirdropClaim `json:"claim"`         // 领取记录
}

// CampaignInfo 活动基本信息
type CampaignInfo struct {
	TokenSymbol string          `json:"tokenSymbol"` // 代币符号
	StartTime   time.Time       `json:"startTime"`   // 开始时间
	EndTime     time.Time       `json:"endTime"`     // 结束时间
	IsActive    bool            `json:"isActive"`    // 是否激活
	TokenAmount decimal.Decimal `json:"tokenAmount"` // 代币数量
}

// EligibilityResponse 资格检查响应
type EligibilityResponse struct {
	Eligible        bool            `json:"eligible"`                  // 是否有资格
	Reason          string          `json:"reason"`                    // 资格原因说明
	CampaignName    string          `json:"campaignName"`              // 活动名称
	TokenSymbol     string          `json:"tokenSymbol"`               // 代币符号
	TokenAmount     decimal.Decimal `json:"tokenAmount"`               // 可领取的代币数量
	StartTime       time.Time       `json:"startTime,omitempty"`       // 活动开始时间
	EndTime         time.Time       `json:"endTime,omitempty"`         // 活动结束时间
	TotalClaims     int             `json:"totalClaims,omitempty"`     // 总领取人数
	MaxClaims       int             `json:"maxClaims,omitempty"`       // 最大领取人数
	RemainingClaims int             `json:"remainingClaims,omitempty"` // 剩余可领取数量
	ClaimedAt       *time.Time      `json:"claimedAt,omitempty"`       // 已领取时间（如果已领取）
	ClaimStatus     string          `json:"claimStatus,omitempty"`     // 领取状态（如果已领取）
}
