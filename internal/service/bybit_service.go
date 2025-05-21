package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bybit-mcp/internal/model"
	"github.com/bybit-mcp/pkg/bybitapi"
)

// bybitService 实现了BybitService接口
type bybitService struct {
	client *bybitapi.Client
}

// NewBybitService 创建一个新的Bybit服务实例
func NewBybitService(apiKey, apiSecret string) BybitService {
	client := bybitapi.NewClient(apiKey, apiSecret)
	return &bybitService{
		client: client,
	}
}

// 解析API响应
func (s *bybitService) parseResponse(data []byte) (*model.Response, error) {
	var response model.Response
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}
	return &response, nil
}

// ==================== 市场数据服务实现 ====================

// GetKline 获取K线数据
func (s *bybitService) GetKline(ctx context.Context, category, symbol, interval string, limit int) (*model.Response, error) {
	data, err := s.client.GetKline(category, symbol, interval, limit)
	if err != nil {
		return nil, err
	}
	return s.parseResponse(data)
}

// GetOrderbook 获取订单簿
func (s *bybitService) GetOrderbook(ctx context.Context, category, symbol string, limit int) (*model.Response, error) {
	data, err := s.client.GetOrderbook(category, symbol, limit)
	if err != nil {
		return nil, err
	}
	return s.parseResponse(data)
}

// GetTickers 获取行情数据
func (s *bybitService) GetTickers(ctx context.Context, category, symbol string) (*model.Response, error) {
	data, err := s.client.GetTickers(category, symbol)
	if err != nil {
		return nil, err
	}
	return s.parseResponse(data)
}

// GetRecentTrades 获取最近成交
func (s *bybitService) GetRecentTrades(ctx context.Context, category, symbol string, limit int) (*model.Response, error) {
	// 调用API获取最近成交
	params := map[string]string{
		"category": category,
		"symbol":   symbol,
	}

	if limit > 0 {
		params["limit"] = fmt.Sprintf("%d", limit)
	}

	data, err := s.client.sendRequest("GET", "market/recent-trade", params, false)
	if err != nil {
		return nil, err
	}
	return s.parseResponse(data)
}

// ==================== 订单管理服务实现 ====================

// CreateOrder 创建订单
func (s *bybitService) CreateOrder(ctx context.Context, req *model.OrderRequest) (*model.Response, error) {
	// 构建参数
	params := map[string]string{
		"category":  req.Category,
		"symbol":    req.Symbol,
		"side":      req.Side,
		"orderType": req.OrderType,
		"qty":       fmt.Sprintf("%f", req.Qty),
	}

	if req.Price > 0 {
		params["price"] = fmt.Sprintf("%f", req.Price)
	}

	if req.TimeInForce != "" {
		params["timeInForce"] = req.TimeInForce
	}

	if req.OrderLinkId != "" {
		params["orderLinkId"] = req.OrderLinkId
	}

	if req.TakeProfit > 0 {
		params["takeProfit"] = fmt.Sprintf("%f", req.TakeProfit)
	}

	if req.StopLoss > 0 {
		params["stopLoss"] = fmt.Sprintf("%f", req.StopLoss)
	}

	if req.ReduceOnly {
		params["reduceOnly"] = "true"
	}

	if req.CloseOnTrigger {
		params["closeOnTrigger"] = "true"
	}

	data, err := s.client.sendRequest("POST", "order/create", params, true)
	if err != nil {
		return nil, err
	}
	return s.parseResponse(data)
}

// CancelOrder 取消订单
func (s *bybitService) CancelOrder(ctx context.Context, category, symbol, orderId string) (*model.Response, error) {
	data, err := s.client.CancelOrder(category, symbol, orderId)
	if err != nil {
		return nil, err
	}
	return s.parseResponse(data)
}

// GetOrders 获取订单列表
func (s *bybitService) GetOrders(ctx context.Context, category, symbol string, limit int) (*model.Response, error) {
	data, err := s.client.GetOrders(category, symbol, limit)
	if err != nil {
		return nil, err
	}
	return s.parseResponse(data)
}

// GetOrderHistory 获取历史订单
func (s *bybitService) GetOrderHistory(ctx context.Context, category, symbol string, limit int) (*model.Response, error) {
	// 调用API获取历史订单
	params := map[string]string{
		"category": category,
	}

	if symbol != "" {
		params["symbol"] = symbol
	}

	if limit > 0 {
		params["limit"] = fmt.Sprintf("%d", limit)
	}

	data, err := s.client.sendRequest("GET", "order/history", params, true)
	if err != nil {
		return nil, err
	}
	return s.parseResponse(data)
}

// CancelAllOrders 取消所有订单
func (s *bybitService) CancelAllOrders(ctx context.Context, category, symbol, settleCoin string) (*model.Response, error) {
	// 调用API取消所有订单
	params := map[string]string{
		"category": category,
	}

	if symbol != "" {
		params["symbol"] = symbol
	}

	if settleCoin != "" {
		params["settleCoin"] = settleCoin
	}

	data, err := s.client.sendRequest("POST", "order/cancel-all", params, true)
	if err != nil {
		return nil, err
	}
	return s.parseResponse(data)
}

// ==================== 仓位管理服务实现 ====================

// GetPositions 获取仓位
func (s *bybitService) GetPositions(ctx context.Context, category, symbol string) (*model.Response, error) {
	data, err := s.client.GetPositions(category, symbol)
	if err != nil {
		return nil, err
	}
	return s.parseResponse(data)
}

// SetLeverage 设置杠杆
func (s *bybitService) SetLeverage(ctx context.Context, category, symbol string, leverage float64) (*model.Response, error) {
	// 调用API设置杠杆
	params := map[string]string{
		"category": category,
		"symbol":   symbol,
		"leverage": fmt.Sprintf("%f", leverage),
	}

	data, err := s.client.sendRequest("POST", "position/set-leverage", params, true)
	if err != nil {
		return nil, err
	}
	return s.parseResponse(data)
}

// SetTpSlMode 设置止盈止损模式
func (s *bybitService) SetTpSlMode(ctx context.Context, category, symbol, tpSlMode string) (*model.Response, error) {
	// 调用API设置止盈止损模式
	params := map[string]string{
		"category": category,
		"symbol":   symbol,
		"tpSlMode": tpSlMode,
	}

	data, err := s.client.sendRequest("POST", "position/set-tpsl-mode", params, true)
	if err != nil {
		return nil, err
	}
	return s.parseResponse(data)
}

// SetRiskLimit 设置风险限额
func (s *bybitService) SetRiskLimit(ctx context.Context, category, symbol string, riskId int) (*model.Response, error) {
	// 调用API设置风险限额
	params := map[string]string{
		"category": category,
		"symbol":   symbol,
		"riskId":   fmt.Sprintf("%d", riskId),
	}

	data, err := s.client.sendRequest("POST", "position/set-risk-limit", params, true)
	if err != nil {
		return nil, err
	}
	return s.parseResponse(data)
}

// ==================== 账户管理服务实现 ====================

// GetWalletBalance 获取钱包余额
func (s *bybitService) GetWalletBalance(ctx context.Context, accountType, coin string) (*model.Response, error) {
	data, err := s.client.GetWalletBalance(accountType, coin)
	if err != nil {
		return nil, err
	}
	return s.parseResponse(data)
}

// GetAccountInfo 获取账户信息
func (s *bybitService) GetAccountInfo(ctx context.Context) (*model.Response, error) {
	// 调用API获取账户信息
	data, err := s.client.sendRequest("GET", "account/info", map[string]string{}, true)
	if err != nil {
		return nil, err
	}
	return s.parseResponse(data)
}

// GetFeeRate 获取手续费率
func (s *bybitService) GetFeeRate(ctx context.Context, category, symbol string) (*model.Response, error) {
	// 调用API获取手续费率
	params := map[string]string{
		"category": category,
	}

	if symbol != "" {
		params["symbol"] = symbol
	}

	data, err := s.client.sendRequest("GET", "account/fee-rate", params, true)
	if err != nil {
		return nil, err
	}
	return s.parseResponse(data)
}

// GetAccountMode 获取账户模式
func (s *bybitService) GetAccountMode(ctx context.Context) (*model.Response, error) {
	// 调用API获取账户模式
	data, err := s.client.sendRequest("GET", "account/mode", map[string]string{}, true)
	if err != nil {
		return nil, err
	}
	return s.parseResponse(data)
}

// SetAccountMode 设置账户模式
func (s *bybitService) SetAccountMode(ctx context.Context, accountMode string) (*model.Response, error) {
	// 调用API设置账户模式
	params := map[string]string{
		"accountMode": accountMode,
	}

	data, err := s.client.sendRequest("POST", "account/set-mode", params, true)
	if err != nil {
		return nil, err
	}
	return s.parseResponse(data)
}

// ==================== 资产管理服务实现 ====================

// GetAssetInfo 获取资产信息
func (s *bybitService) GetAssetInfo(ctx context.Context, accountType string) (*model.Response, error) {
	data, err := s.client.GetAssetInfo(accountType)
	if err != nil {
		return nil, err
	}
	return s.parseResponse(data)
}

// AssetTransfer 资产划转
func (s *bybitService) AssetTransfer(ctx context.Context, fromAccountType, toAccountType, coin string, amount float64) (*model.Response, error) {
	// 调用API进行资产划转
	params := map[string]string{
		"fromAccountType": fromAccountType,
		"toAccountType":   toAccountType,
		"coin":            coin,
		"amount":          fmt.Sprintf("%f", amount),
	}

	data, err := s.client.sendRequest("POST", "asset/transfer", params, true)
	if err != nil {
		return nil, err
	}
	return s.parseResponse(data)
}

// GetTransferHistory 获取划转历史
func (s *bybitService) GetTransferHistory(ctx context.Context, coin string, limit int) (*model.Response, error) {
	// 调用API获取划转历史
	params := map[string]string{}

	if coin != "" {
		params["coin"] = coin
	}

	if limit > 0 {
		params["limit"] = fmt.Sprintf("%d", limit)
	}

	data, err := s.client.sendRequest("GET", "asset/transfer/query-transfer-list", params, true)
	if err != nil {
		return nil, err
	}
	return s.parseResponse(data)
}

// GetDepositHistory 获取充值历史
func (s *bybitService) GetDepositHistory(ctx context.Context, coin string, limit int) (*model.Response, error) {
	// 调用API获取充值历史
	params := map[string]string{}

	if coin != "" {
		params["coin"] = coin
	}

	if limit > 0 {
		params["limit"] = fmt.Sprintf("%d", limit)
	}

	data, err := s.client.sendRequest("GET", "asset/deposit/query-record", params, true)
	if err != nil {
		return nil, err
	}
	return s.parseResponse(data)
}

// GetWithdrawalHistory 获取提现历史
func (s *bybitService) GetWithdrawalHistory(ctx context.Context, coin string, limit int) (*model.Response, error) {
	// 调用API获取提现历史
	params := map[string]string{}

	if coin != "" {
		params["coin"] = coin
	}

	if limit > 0 {
		params["limit"] = fmt.Sprintf("%d", limit)
	}

	data, err := s.client.sendRequest("GET", "asset/withdraw/query-record", params, true)
	if err != nil {
		return nil, err
	}
	return s.parseResponse(data)
}