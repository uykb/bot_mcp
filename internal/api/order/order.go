package order

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/bybit-mcp/internal/model"
	"github.com/bybit-mcp/pkg/bybitapi"
	"github.com/bybit-mcp/pkg/errors"
	"github.com/bybit-mcp/pkg/logger"
)

// OrderService 提供订单管理相关的API服务
type OrderService struct {
	client *bybitapi.Client
	logger *logger.Logger
}

// NewOrderService 创建一个新的订单管理服务
func NewOrderService(client *bybitapi.Client, logLevel, logOutput string) *OrderService {
	return &OrderService{
		client: client,
		logger: logger.New(logLevel, logOutput),
	}
}

// CreateOrder 创建订单
func (s *OrderService) CreateOrder(ctx context.Context, category, symbol, side, orderType string, qty float64, price float64, options map[string]string) (*model.Response, error) {
	s.logger.Debug("创建订单: category=%s, symbol=%s, side=%s, orderType=%s, qty=%f, price=%f", category, symbol, side, orderType, qty, price)

	// 构建请求参数
	params := map[string]string{
		"category":  category,
		"symbol":    symbol,
		"side":      side,
		"orderType": orderType,
		"qty":       strconv.FormatFloat(qty, 'f', -1, 64),
	}

	// 添加价格参数（如果不是市价单）
	if orderType != "Market" && price > 0 {
		params["price"] = strconv.FormatFloat(price, 'f', -1, 64)
	}

	// 添加可选参数
	for k, v := range options {
		params[k] = v
	}

	// 发送请求
	response, err := s.client.Post("order/create", params, true)
	if err != nil {
		s.logger.Error("创建订单失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrOrderFailed,
			Message: "创建订单失败",
			Cause:   err,
		}
	}

	// 解析响应
	var resp model.Response
	if err := json.Unmarshal(response, &resp); err != nil {
		s.logger.Error("解析订单响应失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIResponseInvalid,
			Message: "解析订单响应失败",
			Cause:   err,
		}
	}

	return &resp, nil
}

// CancelOrder 取消订单
func (s *OrderService) CancelOrder(ctx context.Context, category, symbol, orderId, orderLinkId string) (*model.Response, error) {
	s.logger.Debug("取消订单: category=%s, symbol=%s, orderId=%s, orderLinkId=%s", category, symbol, orderId, orderLinkId)

	// 构建请求参数
	params := map[string]string{
		"category": category,
		"symbol":   symbol,
	}

	// 添加订单ID参数（至少需要一个）
	if orderId != "" {
		params["orderId"] = orderId
	}
	if orderLinkId != "" {
		params["orderLinkId"] = orderLinkId
	}

	// 发送请求
	response, err := s.client.Post("order/cancel", params, true)
	if err != nil {
		s.logger.Error("取消订单失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrOrderCancelFailed,
			Message: "取消订单失败",
			Cause:   err,
		}
	}

	// 解析响应
	var resp model.Response
	if err := json.Unmarshal(response, &resp); err != nil {
		s.logger.Error("解析取消订单响应失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIResponseInvalid,
			Message: "解析取消订单响应失败",
			Cause:   err,
		}
	}

	return &resp, nil
}

// GetOrders 获取订单列表
func (s *OrderService) GetOrders(ctx context.Context, category, symbol, orderId, orderLinkId, orderStatus string, limit int) (*model.Response, error) {
	s.logger.Debug("获取订单列表: category=%s, symbol=%s, status=%s", category, symbol, orderStatus)

	// 构建请求参数
	params := map[string]string{
		"category": category,
	}

	// 添加可选参数
	if symbol != "" {
		params["symbol"] = symbol
	}
	if orderId != "" {
		params["orderId"] = orderId
	}
	if orderLinkId != "" {
		params["orderLinkId"] = orderLinkId
	}
	if orderStatus != "" {
		params["orderStatus"] = orderStatus
	}
	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	// 发送请求
	response, err := s.client.Get("order/history", params, true)
	if err != nil {
		s.logger.Error("获取订单列表失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIRequestFailed,
			Message: "获取订单列表失败",
			Cause:   err,
		}
	}

	// 解析响应
	var resp model.Response
	if err := json.Unmarshal(response, &resp); err != nil {
		s.logger.Error("解析订单列表响应失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIResponseInvalid,
			Message: "解析订单列表响应失败",
			Cause:   err,
		}
	}

	return &resp, nil
}

// AmendOrder 修改订单
func (s *OrderService) AmendOrder(ctx context.Context, category, symbol, orderId, orderLinkId string, qty float64, price float64, options map[string]string) (*model.Response, error) {
	s.logger.Debug("修改订单: category=%s, symbol=%s, orderId=%s, orderLinkId=%s", category, symbol, orderId, orderLinkId)

	// 构建请求参数
	params := map[string]string{
		"category": category,
		"symbol":   symbol,
	}

	// 添加订单ID参数（至少需要一个）
	if orderId != "" {
		params["orderId"] = orderId
	}
	if orderLinkId != "" {
		params["orderLinkId"] = orderLinkId
	}

	// 添加修改参数
	if qty > 0 {
		params["qty"] = strconv.FormatFloat(qty, 'f', -1, 64)
	}
	if price > 0 {
		params["price"] = strconv.FormatFloat(price, 'f', -1, 64)
	}

	// 添加可选参数
	for k, v := range options {
		params[k] = v
	}

	// 发送请求
	response, err := s.client.Post("order/amend", params, true)
	if err != nil {
		s.logger.Error("修改订单失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIRequestFailed,
			Message: "修改订单失败",
			Cause:   err,
		}
	}

	// 解析响应
	var resp model.Response
	if err := json.Unmarshal(response, &resp); err != nil {
		s.logger.Error("解析修改订单响应失败: %v", err)
		return nil, &errors.Error{
			Code:    errors.ErrAPIResponseInvalid,
			Message: "解析修改订单响应失败",
			Cause:   err,
		}
	}

	return &resp, nil
}