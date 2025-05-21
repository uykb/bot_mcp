package position

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/bybit-mcp/internal/model"
	"github.com/bybit-mcp/pkg/bybitapi"
	"github.com/bybit-mcp/pkg/errors"
	"github.com/bybit-mcp/pkg/logger"
)

// PositionService 提供仓位管理相关的API服务
type PositionService struct {
	client *bybitapi.Client
	logger *logger.Logger
}

// NewPositionService 创建一个新的仓位管理服务
func NewPositionService(client *bybitapi.Client, logLevel, logOutput string) *PositionService {
	return &PositionService{
		client: client,
		logger: logger.New(logLevel, logOutput),
	}
}

// GetPositions 获取仓位列表
func (s *PositionService) GetPositions(ctx context.Context, category, symbol, settleCoin, positionIdx string) (*model.Response, error) {
	s.logger.Debug("获取仓位列表: category=%s, symbol=%s", category, symbol)

	// 构建请求参数
	params := map[string]string{
		"category": category,
	}

	// 添加可选参数
	if symbol != "" {
		params["symbol"] = symbol
	}
	if settleCoin != "" {
		params["settleCoin"] = settleCoin
	}
	if positionIdx != "" {
		params["positionIdx"] = positionIdx
	}

	// 发送请求
	response, err := s.client.Get("position/list", params, true)
	if err != nil {
		s.logger.Error("获取仓位列表失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIRequestFailed,
			Message: "获取仓位列表失败",
			Cause:   err,
		}
	}

	// 解析响应
	var resp model.Response
	if err := json.Unmarshal(response, &resp); err != nil {
		s.logger.Error("解析仓位列表响应失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIResponseInvalid,
			Message: "解析仓位列表响应失败",
			Cause:   err,
		}
	}

	return &resp, nil
}

// SetLeverage 设置杠杆
func (s *PositionService) SetLeverage(ctx context.Context, category, symbol string, buyLeverage, sellLeverage float64) (*model.Response, error) {
	s.logger.Debug("设置杠杆: category=%s, symbol=%s, buyLeverage=%f, sellLeverage=%f", category, symbol, buyLeverage, sellLeverage)

	// 构建请求参数
	params := map[string]string{
		"category":     category,
		"symbol":       symbol,
		"buyLeverage":  strconv.FormatFloat(buyLeverage, 'f', -1, 64),
		"sellLeverage": strconv.FormatFloat(sellLeverage, 'f', -1, 64),
	}

	// 发送请求
	response, err := s.client.Post("position/set-leverage", params, true)
	if err != nil {
		s.logger.Error("设置杠杆失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIRequestFailed,
			Message: "设置杠杆失败",
			Cause:   err,
		}
	}

	// 解析响应
	var resp model.Response
	if err := json.Unmarshal(response, &resp); err != nil {
		s.logger.Error("解析设置杠杆响应失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIResponseInvalid,
			Message: "解析设置杠杆响应失败",
			Cause:   err,
		}
	}

	return &resp, nil
}

// SetTradingStop 设置止盈止损
func (s *PositionService) SetTradingStop(ctx context.Context, category, symbol string, takeProfit, stopLoss float64, options map[string]string) (*model.Response, error) {
	s.logger.Debug("设置止盈止损: category=%s, symbol=%s, takeProfit=%f, stopLoss=%f", category, symbol, takeProfit, stopLoss)

	// 构建请求参数
	params := map[string]string{
		"category": category,
		"symbol":   symbol,
	}

	// 添加止盈止损参数
	if takeProfit > 0 {
		params["takeProfit"] = strconv.FormatFloat(takeProfit, 'f', -1, 64)
	}
	if stopLoss > 0 {
		params["stopLoss"] = strconv.FormatFloat(stopLoss, 'f', -1, 64)
	}

	// 添加可选参数
	for k, v := range options {
		params[k] = v
	}

	// 发送请求
	response, err := s.client.Post("position/trading-stop", params, true)
	if err != nil {
		s.logger.Error("设置止盈止损失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIRequestFailed,
			Message: "设置止盈止损失败",
			Cause:   err,
		}
	}

	// 解析响应
	var resp model.Response
	if err := json.Unmarshal(response, &resp); err != nil {
		s.logger.Error("解析设置止盈止损响应失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIResponseInvalid,
			Message: "解析设置止盈止损响应失败",
			Cause:   err,
		}
	}

	return &resp, nil
}

// SwitchPositionMode 切换持仓模式
func (s *PositionService) SwitchPositionMode(ctx context.Context, category, symbol, mode string) (*model.Response, error) {
	s.logger.Debug("切换持仓模式: category=%s, symbol=%s, mode=%s", category, symbol, mode)

	// 构建请求参数
	params := map[string]string{
		"category": category,
	}

	// 添加可选参数
	if symbol != "" {
		params["symbol"] = symbol
	}
	if mode != "" {
		params["mode"] = mode
	}

	// 发送请求
	response, err := s.client.Post("position/switch-mode", params, true)
	if err != nil {
		s.logger.Error("切换持仓模式失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIRequestFailed,
			Message: "切换持仓模式失败",
			Cause:   err,
		}
	}

	// 解析响应
	var resp model.Response
	if err := json.Unmarshal(response, &resp); err != nil {
		s.logger.Error("解析切换持仓模式响应失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIResponseInvalid,
			Message: "解析切换持仓模式响应失败",
			Cause:   err,
		}
	}

	return &resp, nil
}