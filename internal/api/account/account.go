package account

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/bybit-mcp/internal/model"
	"github.com/bybit-mcp/pkg/bybitapi"
	"github.com/bybit-mcp/pkg/errors"
	"github.com/bybit-mcp/pkg/logger"
)

// AccountService 提供账户管理相关的API服务
type AccountService struct {
	client *bybitapi.Client
	logger *logger.Logger
}

// NewAccountService 创建一个新的账户管理服务
func NewAccountService(client *bybitapi.Client, logLevel, logOutput string) *AccountService {
	return &AccountService{
		client: client,
		logger: logger.New(logLevel, logOutput),
	}
}

// GetWalletBalance 获取钱包余额
func (s *AccountService) GetWalletBalance(ctx context.Context, accountType, coin string) (*model.Response, error) {
	s.logger.Debug("获取钱包余额: accountType=%s, coin=%s", accountType, coin)

	// 构建请求参数
	params := map[string]string{}

	// 添加可选参数
	if accountType != "" {
		params["accountType"] = accountType
	}
	if coin != "" {
		params["coin"] = coin
	}

	// 发送请求
	response, err := s.client.Get("account/wallet-balance", params, true)
	if err != nil {
		s.logger.Error("获取钱包余额失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIRequestFailed,
			Message: "获取钱包余额失败",
			Cause:   err,
		}
	}

	// 解析响应
	var resp model.Response
	if err := json.Unmarshal(response, &resp); err != nil {
		s.logger.Error("解析钱包余额响应失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIResponseInvalid,
			Message: "解析钱包余额响应失败",
			Cause:   err,
		}
	}

	return &resp, nil
}

// GetFeeRate 获取手续费率
func (s *AccountService) GetFeeRate(ctx context.Context, category, symbol string) (*model.Response, error) {
	s.logger.Debug("获取手续费率: category=%s, symbol=%s", category, symbol)

	// 构建请求参数
	params := map[string]string{
		"category": category,
	}

	// 添加可选参数
	if symbol != "" {
		params["symbol"] = symbol
	}

	// 发送请求
	response, err := s.client.Get("account/fee-rate", params, true)
	if err != nil {
		s.logger.Error("获取手续费率失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIRequestFailed,
			Message: "获取手续费率失败",
			Cause:   err,
		}
	}

	// 解析响应
	var resp model.Response
	if err := json.Unmarshal(response, &resp); err != nil {
		s.logger.Error("解析手续费率响应失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIResponseInvalid,
			Message: "解析手续费率响应失败",
			Cause:   err,
		}
	}

	return &resp, nil
}

// GetAccountInfo 获取账户信息
func (s *AccountService) GetAccountInfo(ctx context.Context) (*model.Response, error) {
	s.logger.Debug("获取账户信息")

	// 发送请求
	response, err := s.client.Get("account/info", nil, true)
	if err != nil {
		s.logger.Error("获取账户信息失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIRequestFailed,
			Message: "获取账户信息失败",
			Cause:   err,
		}
	}

	// 解析响应
	var resp model.Response
	if err := json.Unmarshal(response, &resp); err != nil {
		s.logger.Error("解析账户信息响应失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIResponseInvalid,
			Message: "解析账户信息响应失败",
			Cause:   err,
		}
	}

	return &resp, nil
}

// SetMarginMode 设置保证金模式
func (s *AccountService) SetMarginMode(ctx context.Context, marginMode string) (*model.Response, error) {
	s.logger.Debug("设置保证金模式: marginMode=%s", marginMode)

	// 构建请求参数
	params := map[string]string{
		"setMarginMode": marginMode,
	}

	// 发送请求
	response, err := s.client.Post("account/set-margin-mode", params, true)
	if err != nil {
		s.logger.Error("设置保证金模式失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIRequestFailed,
			Message: "设置保证金模式失败",
			Cause:   err,
		}
	}

	// 解析响应
	var resp model.Response
	if err := json.Unmarshal(response, &resp); err != nil {
		s.logger.Error("解析设置保证金模式响应失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIResponseInvalid,
			Message: "解析设置保证金模式响应失败",
			Cause:   err,
		}
	}

	return &resp, nil
}