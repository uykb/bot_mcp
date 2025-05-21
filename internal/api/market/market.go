package market

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/bybit-mcp/internal/model"
	"github.com/bybit-mcp/pkg/bybitapi"
	"github.com/bybit-mcp/pkg/errors"
	"github.com/bybit-mcp/pkg/logger"
)

// MarketService 提供市场数据相关的API服务
type MarketService struct {
	client *bybitapi.Client
	logger *logger.Logger
}

// NewMarketService 创建一个新的市场数据服务
func NewMarketService(client *bybitapi.Client, logLevel, logOutput string) *MarketService {
	return &MarketService{
		client: client,
		logger: logger.New(logLevel, logOutput),
	}
}

// GetKline 获取K线数据
func (s *MarketService) GetKline(ctx context.Context, category, symbol, interval string, limit int, start, end int64) (*model.Response, error) {
	s.logger.Debug("获取K线数据: category=%s, symbol=%s, interval=%s", category, symbol, interval)

	// 构建请求参数
	params := map[string]string{
		"category": category,
		"symbol":   symbol,
		"interval": interval,
	}

	// 添加可选参数
	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}
	if start > 0 {
		params["start"] = strconv.FormatInt(start, 10)
	}
	if end > 0 {
		params["end"] = strconv.FormatInt(end, 10)
	}

	// 发送请求
	response, err := s.client.Get("market/kline", params, true)
	if err != nil {
		s.logger.Error("获取K线数据失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIRequestFailed,
			Message: "获取K线数据失败",
			Cause:   err,
		}
	}

	// 解析响应
	var resp model.Response
	if err := json.Unmarshal(response, &resp); err != nil {
		s.logger.Error("解析K线数据失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIResponseInvalid,
			Message: "解析K线数据失败",
			Cause:   err,
		}
	}

	return &resp, nil
}

// GetOrderbook 获取订单簿数据
func (s *MarketService) GetOrderbook(ctx context.Context, category, symbol string, limit int) (*model.Response, error) {
	s.logger.Debug("获取订单簿数据: category=%s, symbol=%s, limit=%d", category, symbol, limit)

	// 构建请求参数
	params := map[string]string{
		"category": category,
		"symbol":   symbol,
	}

	// 添加可选参数
	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	// 发送请求
	response, err := s.client.Get("market/orderbook", params, true)
	if err != nil {
		s.logger.Error("获取订单簿数据失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIRequestFailed,
			Message: "获取订单簿数据失败",
			Cause:   err,
		}
	}

	// 解析响应
	var resp model.Response
	if err := json.Unmarshal(response, &resp); err != nil {
		s.logger.Error("解析订单簿数据失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIResponseInvalid,
			Message: "解析订单簿数据失败",
			Cause:   err,
		}
	}

	return &resp, nil
}

// GetTickers 获取行情数据
func (s *MarketService) GetTickers(ctx context.Context, category, symbol string) (*model.Response, error) {
	s.logger.Debug("获取行情数据: category=%s, symbol=%s", category, symbol)

	// 构建请求参数
	params := map[string]string{
		"category": category,
	}

	// 添加可选参数
	if symbol != "" {
		params["symbol"] = symbol
	}

	// 发送请求
	response, err := s.client.Get("market/tickers", params, true)
	if err != nil {
		s.logger.Error("获取行情数据失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIRequestFailed,
			Message: "获取行情数据失败",
			Cause:   err,
		}
	}

	// 解析响应
	var resp model.Response
	if err := json.Unmarshal(response, &resp); err != nil {
		s.logger.Error("解析行情数据失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIResponseInvalid,
			Message: "解析行情数据失败",
			Cause:   err,
		}
	}

	return &resp, nil
}

// GetInstruments 获取交易对信息
func (s *MarketService) GetInstruments(ctx context.Context, category, symbol, status string) (*model.Response, error) {
	s.logger.Debug("获取交易对信息: category=%s, symbol=%s, status=%s", category, symbol, status)

	// 构建请求参数
	params := map[string]string{
		"category": category,
	}

	// 添加可选参数
	if symbol != "" {
		params["symbol"] = symbol
	}
	if status != "" {
		params["status"] = status
	}

	// 发送请求
	response, err := s.client.Get("market/instruments-info", params, true)
	if err != nil {
		s.logger.Error("获取交易对信息失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIRequestFailed,
			Message: "获取交易对信息失败",
			Cause:   err,
		}
	}

	// 解析响应
	var resp model.Response
	if err := json.Unmarshal(response, &resp); err != nil {
		s.logger.Error("解析交易对信息失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIResponseInvalid,
			Message: "解析交易对信息失败",
			Cause:   err,
		}
	}

	return &resp, nil
}