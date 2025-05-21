package asset

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/bybit-mcp/internal/model"
	"github.com/bybit-mcp/pkg/bybitapi"
	"github.com/bybit-mcp/pkg/errors"
	"github.com/bybit-mcp/pkg/logger"
)

// AssetService 提供资产管理相关的API服务
type AssetService struct {
	client *bybitapi.Client
	logger *logger.Logger
}

// NewAssetService 创建一个新的资产管理服务
func NewAssetService(client *bybitapi.Client, logLevel, logOutput string) *AssetService {
	return &AssetService{
		client: client,
		logger: logger.New(logLevel, logOutput),
	}
}

// GetCoinBalance 获取币种余额
func (s *AssetService) GetCoinBalance(ctx context.Context, coin, accountType string) (*model.Response, error) {
	s.logger.Debug("获取币种余额: coin=%s, accountType=%s", coin, accountType)

	// 构建请求参数
	params := map[string]string{}

	// 添加可选参数
	if coin != "" {
		params["coin"] = coin
	}
	if accountType != "" {
		params["accountType"] = accountType
	}

	// 发送请求
	response, err := s.client.Get("asset/transfer/query-asset-info", params, true)
	if err != nil {
		s.logger.Error("获取币种余额失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIRequestFailed,
			Message: "获取币种余额失败",
			Cause:   err,
		}
	}

	// 解析响应
	var resp model.Response
	if err := json.Unmarshal(response, &resp); err != nil {
		s.logger.Error("解析币种余额响应失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIResponseInvalid,
			Message: "解析币种余额响应失败",
			Cause:   err,
		}
	}

	return &resp, nil
}

// TransferAsset 资产划转
func (s *AssetService) TransferAsset(ctx context.Context, transferId, coin, amount, fromAccountType, toAccountType string) (*model.Response, error) {
	s.logger.Debug("资产划转: coin=%s, amount=%s, fromAccountType=%s, toAccountType=%s", coin, amount, fromAccountType, toAccountType)

	// 构建请求参数
	params := map[string]string{
		"transferId":      transferId,
		"coin":            coin,
		"amount":          amount,
		"fromAccountType": fromAccountType,
		"toAccountType":   toAccountType,
	}

	// 发送请求
	response, err := s.client.Post("asset/transfer/inter-transfer", params, true)
	if err != nil {
		s.logger.Error("资产划转失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIRequestFailed,
			Message: "资产划转失败",
			Cause:   err,
		}
	}

	// 解析响应
	var resp model.Response
	if err := json.Unmarshal(response, &resp); err != nil {
		s.logger.Error("解析资产划转响应失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIResponseInvalid,
			Message: "解析资产划转响应失败",
			Cause:   err,
		}
	}

	return &resp, nil
}

// GetTransferHistory 获取划转历史
func (s *AssetService) GetTransferHistory(ctx context.Context, transferId, coin, status string, startTime, endTime int64, limit int) (*model.Response, error) {
	s.logger.Debug("获取划转历史: coin=%s, status=%s", coin, status)

	// 构建请求参数
	params := map[string]string{}

	// 添加可选参数
	if transferId != "" {
		params["transferId"] = transferId
	}
	if coin != "" {
		params["coin"] = coin
	}
	if status != "" {
		params["status"] = status
	}
	if startTime > 0 {
		params["startTime"] = strconv.FormatInt(startTime, 10)
	}
	if endTime > 0 {
		params["endTime"] = strconv.FormatInt(endTime, 10)
	}
	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	// 发送请求
	response, err := s.client.Get("asset/transfer/query-transfer-list", params, true)
	if err != nil {
		s.logger.Error("获取划转历史失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIRequestFailed,
			Message: "获取划转历史失败",
			Cause:   err,
		}
	}

	// 解析响应
	var resp model.Response
	if err := json.Unmarshal(response, &resp); err != nil {
		s.logger.Error("解析划转历史响应失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIResponseInvalid,
			Message: "解析划转历史响应失败",
			Cause:   err,
		}
	}

	return &resp, nil
}

// Withdraw 提现
func (s *AssetService) Withdraw(ctx context.Context, coin, chain, address, tag, amount string, options map[string]string) (*model.Response, error) {
	s.logger.Debug("提现: coin=%s, chain=%s, address=%s, amount=%s", coin, chain, address, amount)

	// 构建请求参数
	params := map[string]string{
		"coin":    coin,
		"chain":   chain,
		"address": address,
		"amount":  amount,
	}

	// 添加可选参数
	if tag != "" {
		params["tag"] = tag
	}

	// 添加其他可选参数
	for k, v := range options {
		params[k] = v
	}

	// 发送请求
	response, err := s.client.Post("asset/withdraw/create", params, true)
	if err != nil {
		s.logger.Error("提现失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIRequestFailed,
			Message: "提现失败",
			Cause:   err,
		}
	}

	// 解析响应
	var resp model.Response
	if err := json.Unmarshal(response, &resp); err != nil {
		s.logger.Error("解析提现响应失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIResponseInvalid,
			Message: "解析提现响应失败",
			Cause:   err,
		}
	}

	return &resp, nil
}