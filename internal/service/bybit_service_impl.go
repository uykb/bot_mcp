package service

import (
	"context"

	"github.com/bybit-mcp/internal/api/account"
	"github.com/bybit-mcp/internal/api/asset"
	"github.com/bybit-mcp/internal/api/market"
	"github.com/bybit-mcp/internal/api/order"
	"github.com/bybit-mcp/internal/api/position"
	"github.com/bybit-mcp/internal/model"
	"github.com/bybit-mcp/pkg/bybitapi"
	"github.com/bybit-mcp/pkg/logger"
)

// BybitServiceImpl 实现了BybitService接口
type BybitServiceImpl struct {
	client          *bybitapi.Client
	marketService   *market.MarketService
	orderService    *order.OrderService
	positionService *position.PositionService
	accountService  *account.AccountService
	assetService    *asset.AssetService
	logger          *logger.Logger
}

// NewBybitService 创建一个新的Bybit服务实现
func NewBybitService(apiKey, apiSecret, logLevel, logOutput string) BybitService {
	// 创建API客户端
	client := bybitapi.NewClient(apiKey, apiSecret)

	// 创建日志记录器
	logger := logger.New(logLevel, logOutput)
	logger.Info("初始化Bybit MCP服务")

	// 创建各个API服务
	marketService := market.NewMarketService(client, logLevel, logOutput)
	orderService := order.NewOrderService(client, logLevel, logOutput)
	positionService := position.NewPositionService(client, logLevel, logOutput)
	accountService := account.NewAccountService(client, logLevel, logOutput)
	assetService := asset.NewAssetService(client, logLevel, logOutput)

	return &BybitServiceImpl{
		client:          client,
		marketService:   marketService,
		orderService:    orderService,
		positionService: positionService,
		accountService:  accountService,
		assetService:    assetService,
		logger:          logger,
	}
}

// 市场数据API

// GetKline 获取K线数据
func (s *BybitServiceImpl) GetKline(ctx context.Context, category, symbol, interval string, limit int, start, end int64) (*model.Response, error) {
	s.logger.Debug("调用GetKline服务: category=%s, symbol=%s", category, symbol)
	return s.marketService.GetKline(ctx, category, symbol, interval, limit, start, end)
}

// GetOrderbook 获取订单簿数据
func (s *BybitServiceImpl) GetOrderbook(ctx context.Context, category, symbol string, limit int) (*model.Response, error) {
	s.logger.Debug("调用GetOrderbook服务: category=%s, symbol=%s", category, symbol)
	return s.marketService.GetOrderbook(ctx, category, symbol, limit)
}

// GetTickers 获取行情数据
func (s *BybitServiceImpl) GetTickers(ctx context.Context, category, symbol string) (*model.Response, error) {
	s.logger.Debug("调用GetTickers服务: category=%s, symbol=%s", category, symbol)
	return s.marketService.GetTickers(ctx, category, symbol)
}

// GetInstruments 获取交易对信息
func (s *BybitServiceImpl) GetInstruments(ctx context.Context, category, symbol, status string) (*model.Response, error) {
	s.logger.Debug("调用GetInstruments服务: category=%s, symbol=%s", category, symbol)
	return s.marketService.GetInstruments(ctx, category, symbol, status)
}

// 订单管理API

// CreateOrder 创建订单
func (s *BybitServiceImpl) CreateOrder(ctx context.Context, category, symbol, side, orderType string, qty float64, price float64, options map[string]string) (*model.Response, error) {
	s.logger.Debug("调用CreateOrder服务: category=%s, symbol=%s, side=%s", category, symbol, side)
	return s.orderService.CreateOrder(ctx, category, symbol, side, orderType, qty, price, options)
}

// CancelOrder 取消订单
func (s *BybitServiceImpl) CancelOrder(ctx context.Context, category, symbol, orderId, orderLinkId string) (*model.Response, error) {
	s.logger.Debug("调用CancelOrder服务: category=%s, symbol=%s, orderId=%s", category, symbol, orderId)
	return s.orderService.CancelOrder(ctx, category, symbol, orderId, orderLinkId)
}

// GetOrders 获取订单列表
func (s *BybitServiceImpl) GetOrders(ctx context.Context, category, symbol, orderId, orderLinkId, orderStatus string, limit int) (*model.Response, error) {
	s.logger.Debug("调用GetOrders服务: category=%s, symbol=%s, status=%s", category, symbol, orderStatus)
	return s.orderService.GetOrders(ctx, category, symbol, orderId, orderLinkId, orderStatus, limit)
}

// AmendOrder 修改订单
func (s *BybitServiceImpl) AmendOrder(ctx context.Context, category, symbol, orderId, orderLinkId string, qty float64, price float64, options map[string]string) (*model.Response, error) {
	s.logger.Debug("调用AmendOrder服务: category=%s, symbol=%s, orderId=%s", category, symbol, orderId)
	return s.orderService.AmendOrder(ctx, category, symbol, orderId, orderLinkId, qty, price, options)
}

// 仓位管理API

// GetPositions 获取仓位列表
func (s *BybitServiceImpl) GetPositions(ctx context.Context, category, symbol, settleCoin, positionIdx string) (*model.Response, error) {
	s.logger.Debug("调用GetPositions服务: category=%s, symbol=%s", category, symbol)
	return s.positionService.GetPositions(ctx, category, symbol, settleCoin, positionIdx)
}

// SetLeverage 设置杠杆
func (s *BybitServiceImpl) SetLeverage(ctx context.Context, category, symbol string, buyLeverage, sellLeverage float64) (*model.Response, error) {
	s.logger.Debug("调用SetLeverage服务: category=%s, symbol=%s", category, symbol)
	return s.positionService.SetLeverage(ctx, category, symbol, buyLeverage, sellLeverage)
}

// SetTradingStop 设置止盈止损
func (s *BybitServiceImpl) SetTradingStop(ctx context.Context, category, symbol string, takeProfit, stopLoss float64, options map[string]string) (*model.Response, error) {
	s.logger.Debug("调用SetTradingStop服务: category=%s, symbol=%s", category, symbol)
	return s.positionService.SetTradingStop(ctx, category, symbol, takeProfit, stopLoss, options)
}

// SwitchPositionMode 切换持仓模式
func (s *BybitServiceImpl) SwitchPositionMode(ctx context.Context, category, symbol, mode string) (*model.Response, error) {
	s.logger.Debug("调用SwitchPositionMode服务: category=%s, symbol=%s, mode=%s", category, symbol, mode)
	return s.positionService.SwitchPositionMode(ctx, category, symbol, mode)
}

// 账户管理API

// GetWalletBalance 获取钱包余额
func (s *BybitServiceImpl) GetWalletBalance(ctx context.Context, accountType, coin string) (*model.Response, error) {
	s.logger.Debug("调用GetWalletBalance服务: accountType=%s, coin=%s", accountType, coin)
	return s.accountService.GetWalletBalance(ctx, accountType, coin)
}

// GetFeeRate 获取手续费率
func (s *BybitServiceImpl) GetFeeRate(ctx context.Context, category, symbol string) (*model.Response, error) {
	s.logger.Debug("调用GetFeeRate服务: category=%s, symbol=%s", category, symbol)
	return s.accountService.GetFeeRate(ctx, category, symbol)
}

// GetAccountInfo 获取账户信息
func (s *BybitServiceImpl) GetAccountInfo(ctx context.Context) (*model.Response, error) {
	s.logger.Debug("调用GetAccountInfo服务")
	return s.accountService.GetAccountInfo(ctx)
}

// SetMarginMode 设置保证金模式
func (s *BybitServiceImpl) SetMarginMode(ctx context.Context, marginMode string) (*model.Response, error) {
	s.logger.Debug("调用SetMarginMode服务: marginMode=%s", marginMode)
	return s.accountService.SetMarginMode(ctx, marginMode)
}

// 资产管理API

// GetCoinBalance 获取币种余额
func (s *BybitServiceImpl) GetCoinBalance(ctx context.Context, coin, accountType string) (*model.Response, error) {
	s.logger.Debug("调用GetCoinBalance服务: coin=%s, accountType=%s", coin, accountType)
	return s.assetService.GetCoinBalance(ctx, coin, accountType)
}

// TransferAsset 资产划转
func (s *BybitServiceImpl) TransferAsset(ctx context.Context, transferId, coin, amount, fromAccountType, toAccountType string) (*model.Response, error) {
	s.logger.Debug("调用TransferAsset服务: coin=%s, amount=%s", coin, amount)
	return s.assetService.TransferAsset(ctx, transferId, coin, amount, fromAccountType, toAccountType)
}

// GetTransferHistory 获取划转历史
func (s *BybitServiceImpl) GetTransferHistory(ctx context.Context, transferId, coin, status string, startTime, endTime int64, limit int) (*model.Response, error) {
	s.logger.Debug("调用GetTransferHistory服务: coin=%s, status=%s", coin, status)
	return s.assetService.GetTransferHistory(ctx, transferId, coin, status, startTime, endTime, limit)
}

// Withdraw 提现
func (s *BybitServiceImpl) Withdraw(ctx context.Context, coin, chain, address, tag, amount string, options map[string]string) (*model.Response, error) {
	s.logger.Debug("调用Withdraw服务: coin=%s, chain=%s, amount=%s", coin, chain, amount)
	return s.assetService.Withdraw(ctx, coin, chain, address, tag, amount, options)
}