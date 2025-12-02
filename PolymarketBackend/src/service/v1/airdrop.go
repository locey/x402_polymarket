package service

import (
	"context"
	"errors"
	"github.com/locey/x402_polymarket/PolymarketBackend/src/types/v1"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

// CreateCampaign 创建空投活动
// ctx: 上下文，用于超时控制和链路追踪
// db: 数据库连接实例
// req: 创建空投活动的请求参数
// 返回值: 创建成功的空投活动信息和错误信息
func CreateCampaign(ctx context.Context, db *gorm.DB, req types.CreateCampaignReq) (*types.AirdropCampaign, error) {
	// 参数验证
	if req.Name == "" {
		return nil, errors.New("活动名称不能为空")
	}
	if req.TokenSymbol == "" {
		return nil, errors.New("代币符号不能为空")
	}
	if req.TokenAmount.LessThanOrEqual(decimal.Zero) {
		return nil, errors.New("代币数量必须大于0")
	}
	if req.StartTime.IsZero() || req.EndTime.IsZero() {
		return nil, errors.New("开始时间和结束时间不能为空")
	}
	if req.EndTime.Before(req.StartTime) {
		return nil, errors.New("结束时间不能早于开始时间")
	}

	// 构建空投活动对象
	campaign := &types.AirdropCampaign{
		Name:        req.Name,        // 活动名称
		Description: req.Description, // 活动描述
		TokenSymbol: req.TokenSymbol, // 代币符号
		TokenAmount: req.TokenAmount, // 每个用户获得的代币数量
		MaxClaims:   req.MaxClaims,   // 最大领取人数，0表示无限制
		StartTime:   req.StartTime,   // 活动开始时间
		EndTime:     req.EndTime,     // 活动结束时间
		IsActive:    true,            // 默认激活状态
		TotalClaims: 0,               // 初始领取人数为0
		CreatedAt:   time.Now(),      // 创建时间
		UpdatedAt:   time.Now(),      // 更新时间
	}

	// 使用上下文保存到数据库，支持超时控制
	if err := db.WithContext(ctx).Create(campaign).Error; err != nil {
		return nil, err
	}

	return campaign, nil
}

// ClaimAirdrop 用户领取空投
// ctx: 上下文，用于超时控制和链路追踪
// db: 数据库连接实例
// req: 领取空投的请求参数
// 返回值: 领取记录信息和错误信息
func ClaimAirdrop(ctx context.Context, db *gorm.DB, req types.ClaimAirdropReq) (*types.AirdropClaim, error) {
	// 开始数据库事务，确保数据一致性
	tx := db.WithContext(ctx).Begin()
	// 确保在发生panic时回滚事务
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1. 检查空投活动是否存在
	var campaign types.AirdropCampaign
	if err := tx.First(&campaign, req.CampaignID).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("空投活动不存在")
	}

	// 2. 验证活动状态
	now := time.Now()
	if !campaign.IsActive {
		tx.Rollback()
		return nil, errors.New("空投活动未激活")
	}
	if now.Before(campaign.StartTime) {
		tx.Rollback()
		return nil, errors.New("空投活动尚未开始")
	}
	if now.After(campaign.EndTime) {
		tx.Rollback()
		return nil, errors.New("空投活动已结束")
	}
	if campaign.MaxClaims > 0 && campaign.TotalClaims >= campaign.MaxClaims {
		tx.Rollback()
		return nil, errors.New("空投已领完")
	}

	// 3. 检查用户是否已经领取过（防止重复领取）
	var existingClaim types.AirdropClaim
	if err := tx.Where("airdrop_campaign_id = ? AND wallet_address = ?", req.CampaignID, req.WalletAddress).
		First(&existingClaim).Error; err == nil {
		tx.Rollback()
		return nil, errors.New("已经领取过空投")
	}

	// 4. 创建领取记录
	claim := &types.AirdropClaim{
		AirdropCampaignID: req.CampaignID,       // 关联的空投活动ID
		WalletAddress:     req.WalletAddress,    // 用户钱包地址
		ClaimAmount:       campaign.TokenAmount, // 领取的代币数量（使用活动中设置的固定数量）
		ClaimStatus:       "success",            // 领取状态：成功（简化设计，直接标记成功）
		ClaimedAt:         &now,                 // 领取时间
		CreatedAt:         time.Now(),           // 记录创建时间
		UpdatedAt:         time.Now(),           // 记录更新时间
	}

	// 保存领取记录到数据库
	if err := tx.Create(claim).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// 5. 更新活动的总领取人数统计
	if err := tx.Model(&campaign).Update("total_claims", campaign.TotalClaims+1).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// 提交事务，确保所有操作都成功执行
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return claim, nil
}

// GetActiveCampaigns 获取所有活跃的空投活动
// ctx: 上下文，用于超时控制和链路追踪
// db: 数据库连接实例
// 返回值: 活跃的空投活动列表和错误信息
func GetActiveCampaigns(ctx context.Context, db *gorm.DB) ([]types.AirdropCampaign, error) {
	var campaigns []types.AirdropCampaign
	now := time.Now()

	// 查询条件：活动激活、当前时间在活动时间范围内
	err := db.WithContext(ctx).
		Where("is_active = ? AND start_time <= ? AND end_time >= ?", true, now, now).
		Order("created_at DESC"). // 按创建时间倒序排列，最新的活动显示在前面
		Find(&campaigns).Error

	return campaigns, err
}

func GetCampaignDetail(ctx context.Context, db *gorm.DB, campaignID uint) ([]types.AirdropCampaign, error) {
	var campaign []types.AirdropCampaign

	// 查询指定钱包地址的所有领取记录，并预加载关联的空投活动信息
	err := db.WithContext(ctx).
		Where("id = ?", campaignID).
		Find(&campaign).Error

	return campaign, err
}

// GetUserClaims 获取指定用户的所有空投领取记录
// ctx: 上下文，用于超时控制和链路追踪
// db: 数据库连接实例
// walletAddress: 用户钱包地址
// 返回值: 用户的领取记录列表和错误信息
func GetUserClaims(ctx context.Context, db *gorm.DB, walletAddress string) ([]types.AirdropClaim, error) {
	var claims []types.AirdropClaim

	// 查询指定钱包地址的所有领取记录，并预加载关联的空投活动信息
	err := db.WithContext(ctx).
		Preload("Campaign"). // 预加载关联的空投活动信息，避免N+1查询
		Where("wallet_address = ?", walletAddress).
		Order("created_at DESC"). // 按领取时间倒序排列，最新的记录显示在前面
		Find(&claims).Error

	return claims, err
}

// GetCampaignStats 获取空投活动的统计信息
// ctx: 上下文，用于超时控制和链路追踪
// db: 数据库连接实例
// campaignID: 空投活动ID
// 返回值: 活动统计信息和错误信息
func GetCampaignStats(ctx context.Context, db *gorm.DB, campaignID uint) (*types.CampaignStats, error) {
	// 查询空投活动信息
	var campaign types.AirdropCampaign
	if err := db.WithContext(ctx).First(&campaign, campaignID).Error; err != nil {
		return nil, err
	}

	// 计算剩余可领取数量
	remaining := 0
	if campaign.MaxClaims > 0 {
		remaining = campaign.MaxClaims - campaign.TotalClaims
		if remaining < 0 {
			remaining = 0 // 确保剩余数量不为负数
		}
	}

	// 计算领取率（如果设置了最大领取人数）
	claimRate := 0.0
	if campaign.MaxClaims > 0 {
		claimRate = float64(campaign.TotalClaims) / float64(campaign.MaxClaims)
	}

	// 构建统计信息对象
	stats := &types.CampaignStats{
		CampaignID:   campaign.ID,          // 活动ID
		CampaignName: campaign.Name,        // 活动名称
		TotalClaims:  campaign.TotalClaims, // 总领取人数
		MaxClaims:    campaign.MaxClaims,   // 最大可领取人数
		Remaining:    remaining,            // 剩余可领取数量
		TokenSymbol:  campaign.TokenSymbol, // 代币符号
		TokenAmount:  campaign.TokenAmount, // 每个用户获得的代币数量
		IsActive:     campaign.IsActive,    // 活动是否激活
		ClaimRate:    claimRate,            // 领取率（0-1之间的小数）
		StartTime:    campaign.StartTime,   // 活动开始时间
		EndTime:      campaign.EndTime,     // 活动结束时间
	}

	return stats, nil
}

// UpdateCampaignStatus 更新空投活动的激活状态
// ctx: 上下文，用于超时控制和链路追踪
// db: 数据库连接实例
// campaignID: 空投活动ID
// isActive: 新的激活状态
// 返回值: 错误信息
func UpdateCampaignStatus(ctx context.Context, db *gorm.DB, campaignID uint, isActive bool) error {
	// 更新指定活动的激活状态和更新时间
	return db.WithContext(ctx).Model(&types.AirdropCampaign{}).
		Where("id = ?", campaignID).
		Updates(map[string]interface{}{
			"is_active":  isActive,
			"updated_at": time.Now(),
		}).Error
}

// DeleteCampaign 删除空投活动
// ctx: 上下文，用于超时控制和链路追踪
// db: 数据库连接实例
// campaignID: 要删除的空投活动ID
// forceDelete: 是否强制删除（如果为false，只进行软删除）
// 返回值: 错误信息
func DeleteCampaign(ctx context.Context, db *gorm.DB, campaignID uint, forceDelete bool) error {
	// 开始数据库事务，确保数据一致性
	tx := db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1. 检查空投活动是否存在
	var campaign types.AirdropCampaign
	if err := tx.First(&campaign, campaignID).Error; err != nil {
		tx.Rollback()
		return errors.New("空投活动不存在")
	}

	// 2. 检查活动是否有领取记录
	var claimCount int64
	if err := tx.Model(&types.AirdropClaim{}).
		Where("airdrop_campaign_id = ?", campaignID).
		Count(&claimCount).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 3. 如果有领取记录且不是强制删除，则不允许删除
	if claimCount > 0 && !forceDelete {
		tx.Rollback()
		return errors.New("该活动已有用户领取记录，无法删除。如需强制删除，请设置force_delete=true")
	}

	if forceDelete {
		// 强制删除模式：删除所有相关数据

		// 3.1 删除该活动的所有领取记录
		if err := tx.Where("airdrop_campaign_id = ?", campaignID).
			Delete(&types.AirdropClaim{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		// 3.2 删除空投活动本身
		if err := tx.Delete(&campaign).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		// 软删除模式：只标记为删除状态（如果AirdropCampaign结构体有DeletedAt字段）
		// 或者直接删除（如果没有领取记录）
		if err := tx.Delete(&campaign).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

// UpdateCampaignPartial 部分更新空投活动信息
// ctx: 上下文，用于超时控制和链路追踪
// db: 数据库连接实例
// campaignID: 要更新的空投活动ID
// req: 部分更新空投活动的请求参数
// 返回值: 更新后的空投活动信息和错误信息
func UpdateCampaignPartial(ctx context.Context, db *gorm.DB, campaignID uint, req types.UpdateCampaignPartialReq) (*types.AirdropCampaign, error) {
	// 检查空投活动是否存在
	var campaign types.AirdropCampaign
	if err := db.WithContext(ctx).First(&campaign, campaignID).Error; err != nil {
		return nil, errors.New("空投活动不存在")
	}

	// 构建更新数据
	updates := map[string]interface{}{
		"updated_at": time.Now(),
	}

	// 根据提供的字段更新
	if req.Name != nil {
		if *req.Name == "" {
			return nil, errors.New("活动名称不能为空")
		}
		updates["name"] = *req.Name
	}

	if req.Description != nil {
		updates["description"] = *req.Description
	}

	if req.TokenSymbol != nil {
		if *req.TokenSymbol == "" {
			return nil, errors.New("代币符号不能为空")
		}
		updates["token_symbol"] = *req.TokenSymbol
	}

	if req.TokenAmount != nil {
		if req.TokenAmount.LessThanOrEqual(decimal.Zero) {
			return nil, errors.New("代币数量必须大于0")
		}

		// 检查是否修改了代币数量且活动已有领取记录
		if !req.TokenAmount.Equal(campaign.TokenAmount) {
			var claimCount int64
			if err := db.WithContext(ctx).Model(&types.AirdropClaim{}).
				Where("airdrop_campaign_id = ?", campaignID).
				Count(&claimCount).Error; err != nil {
				return nil, err
			}

			if claimCount > 0 {
				return nil, errors.New("该活动已有用户领取记录，无法修改代币数量")
			}
		}

		updates["token_amount"] = req.TokenAmount
	}

	if req.MaxClaims != nil {
		if *req.MaxClaims < 0 {
			return nil, errors.New("最大领取人数不能为负数")
		}
		updates["max_claims"] = *req.MaxClaims
	}

	if req.StartTime != nil {
		if req.StartTime.IsZero() {
			return nil, errors.New("开始时间不能为空")
		}
		if campaign.EndTime.Before(*req.StartTime) {
			return nil, errors.New("开始时间不能晚于结束时间")
		}
		updates["start_time"] = *req.StartTime
	}

	if req.EndTime != nil {
		if req.EndTime.IsZero() {
			return nil, errors.New("结束时间不能为空")
		}
		if req.EndTime.Before(campaign.StartTime) {
			return nil, errors.New("结束时间不能早于开始时间")
		}
		updates["end_time"] = *req.EndTime
	}

	if req.IsActive != nil {
		updates["is_active"] = *req.IsActive
	}

	// 如果没有提供任何可更新的字段
	if len(updates) == 1 { // 只有 updated_at
		return nil, errors.New("没有提供可更新的字段")
	}

	// 更新数据库
	if err := db.WithContext(ctx).Model(&campaign).Updates(updates).Error; err != nil {
		return nil, err
	}

	// 重新查询获取更新后的完整数据
	var updatedCampaign types.AirdropCampaign
	if err := db.WithContext(ctx).First(&updatedCampaign, campaignID).Error; err != nil {
		return nil, err
	}

	return &updatedCampaign, nil
}

// GetCampaignClaims 获取活动的所有领取记录
// ctx: 上下文，用于超时控制和链路追踪
// db: 数据库连接实例
// campaignID: 空投活动ID
// page: 页码
// limit: 每页数量
// status: 筛选状态（可选）
// 返回值: 领取记录列表、分页信息和错误信息
func GetCampaignClaims(ctx context.Context, db *gorm.DB, campaignID uint, page, limit int, status string) ([]types.AirdropClaim, *types.Pagination, error) {
	// 检查活动是否存在
	var campaign types.AirdropCampaign
	if err := db.WithContext(ctx).First(&campaign, campaignID).Error; err != nil {
		return nil, nil, errors.New("空投活动不存在")
	}

	var claims []types.AirdropClaim
	var total int64

	// 构建查询
	query := db.WithContext(ctx).Model(&types.AirdropClaim{}).
		Where("airdrop_campaign_id = ?", campaignID)

	// 按状态筛选
	if status != "" {
		query = query.Where("claim_status = ?", status)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, nil, err
	}

	// 计算偏移量
	offset := (page - 1) * limit

	// 查询数据
	err := query.Preload("Campaign").
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&claims).Error

	if err != nil {
		return nil, nil, err
	}

	// 构建分页信息
	pagination := &types.Pagination{
		Page:      page,
		Limit:     limit,
		Total:     total,
		TotalPage: int((total + int64(limit) - 1) / int64(limit)),
	}

	return claims, pagination, nil
}

// GetUserCampaignClaim 获取用户在特定活动的领取记录
// ctx: 上下文，用于超时控制和链路追踪
// db: 数据库连接实例
// walletAddress: 用户钱包地址
// campaignID: 空投活动ID
// 返回值: 领取记录和错误信息
func GetUserCampaignClaim(ctx context.Context, db *gorm.DB, walletAddress string, campaignID uint) (*types.AirdropClaim, error) {
	// 检查活动是否存在
	var campaign types.AirdropCampaign
	if err := db.WithContext(ctx).First(&campaign, campaignID).Error; err != nil {
		return nil, errors.New("空投活动不存在")
	}

	var claim types.AirdropClaim

	// 查询用户在特定活动的领取记录
	err := db.WithContext(ctx).
		Preload("Campaign").
		Where("airdrop_campaign_id = ? AND wallet_address = ?", campaignID, walletAddress).
		First(&claim).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("该用户在此活动中没有领取记录")
		}
		return nil, err
	}

	return &claim, nil
}

// GetCampaignClaimsSummary 获取活动领取记录统计摘要
func GetCampaignClaimsSummary(ctx context.Context, db *gorm.DB, campaignID uint) (*types.ClaimsSummary, error) {
	summary := &types.ClaimsSummary{}

	// 获取各种状态的领取数量
	var statusCounts []struct {
		ClaimStatus string
		Count       int64
	}

	err := db.WithContext(ctx).Model(&types.AirdropClaim{}).
		Select("claim_status, COUNT(*) as count").
		Where("airdrop_campaign_id = ?", campaignID).
		Group("claim_status").
		Scan(&statusCounts).Error

	if err != nil {
		return nil, err
	}

	// 统计各种状态的数量
	for _, sc := range statusCounts {
		switch sc.ClaimStatus {
		case "success":
			summary.SuccessCount = int(sc.Count)
		case "pending":
			summary.PendingCount = int(sc.Count)
		case "failed":
			summary.FailedCount = int(sc.Count)
		}
		summary.TotalCount += int(sc.Count)
	}

	// 获取总领取金额（只计算成功的）
	var totalAmount struct {
		Total decimal.Decimal
	}

	err = db.WithContext(ctx).Model(&types.AirdropClaim{}).
		Select("COALESCE(SUM(claim_amount), 0) as total").
		Where("airdrop_campaign_id = ? AND claim_status = 'success'", campaignID).
		Scan(&totalAmount).Error

	if err != nil {
		return nil, err
	}

	summary.TotalAmount = totalAmount.Total

	return summary, nil
}

// CheckEligibility 检查用户空投资格
func CheckEligibility(ctx context.Context, db *gorm.DB, campaignID uint, walletAddress string) (*types.EligibilityResponse, error) {
	// 检查空投活动是否存在且有效
	var campaign types.AirdropCampaign
	if err := db.WithContext(ctx).First(&campaign, campaignID).Error; err != nil {
		return nil, errors.New("空投活动不存在")
	}

	// 检查活动是否激活
	if !campaign.IsActive {
		return &types.EligibilityResponse{
			Eligible:     false,
			Reason:       "空投活动未激活",
			CampaignName: campaign.Name,
			TokenSymbol:  campaign.TokenSymbol,
			TokenAmount:  campaign.TokenAmount,
		}, nil
	}

	// 检查活动时间
	now := time.Now()
	if now.Before(campaign.StartTime) {
		return &types.EligibilityResponse{
			Eligible:     false,
			Reason:       "空投活动尚未开始",
			CampaignName: campaign.Name,
			TokenSymbol:  campaign.TokenSymbol,
			TokenAmount:  campaign.TokenAmount,
			StartTime:    campaign.StartTime,
			EndTime:      campaign.EndTime,
		}, nil
	}

	if now.After(campaign.EndTime) {
		return &types.EligibilityResponse{
			Eligible:     false,
			Reason:       "空投活动已结束",
			CampaignName: campaign.Name,
			TokenSymbol:  campaign.TokenSymbol,
			TokenAmount:  campaign.TokenAmount,
			StartTime:    campaign.StartTime,
			EndTime:      campaign.EndTime,
		}, nil
	}

	// 检查是否达到最大领取人数
	if campaign.MaxClaims > 0 && campaign.TotalClaims >= campaign.MaxClaims {
		return &types.EligibilityResponse{
			Eligible:     false,
			Reason:       "空投已领完",
			CampaignName: campaign.Name,
			TokenSymbol:  campaign.TokenSymbol,
			TokenAmount:  campaign.TokenAmount,
			TotalClaims:  campaign.TotalClaims,
			MaxClaims:    campaign.MaxClaims,
		}, nil
	}

	// 检查用户是否已经领取过
	var existingClaim types.AirdropClaim
	if err := db.WithContext(ctx).
		Where("airdrop_campaign_id = ? AND wallet_address = ?", campaignID, walletAddress).
		First(&existingClaim).Error; err == nil {
		// 用户已经领取过
		return &types.EligibilityResponse{
			Eligible:     false,
			Reason:       "已经领取过空投",
			CampaignName: campaign.Name,
			TokenSymbol:  campaign.TokenSymbol,
			TokenAmount:  campaign.TokenAmount,
			ClaimedAt:    existingClaim.ClaimedAt,
			ClaimStatus:  existingClaim.ClaimStatus,
		}, nil
	}

	// 所有检查通过，用户有资格领取
	return &types.EligibilityResponse{
		Eligible:     true,
		Reason:       "可以领取空投",
		CampaignName: campaign.Name,
		TokenSymbol:  campaign.TokenSymbol,
		TokenAmount:  campaign.TokenAmount,
		StartTime:    campaign.StartTime,
		EndTime:      campaign.EndTime,
	}, nil
}
