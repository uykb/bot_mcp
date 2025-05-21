package service

import (
	"context"

	"github.com/bybit-mcp/internal/model"
)

// BybitService 是Bybit MCP服务的接口定义
// 它包含了所有Bybit V5 API的主要功能模块
type BybitService interface {
	// 市场数据API
	GetKline(ctx context.Context, category, symbol, interval string, limit int, start, end int64) (*model.Response, error)
	GetOrderbook(ctx context.Context, category, symbol string, limit int) (*model.Response, error)
	GetTickers(ctx context.Context, category, symbol string) (*model.Response, error)
	GetInstruments(ctx context.Context, category, symbol, status string) (*model.Response, error)

	// 订单管理API
	CreateOrder(ctx context.Context, category, symbol, side, orderType string, qty float64, price float64, options map[string]string) (*model.Response, error)
	CancelOrder(ctx context.Context, category, symbol, orderId, orderLinkId string) (*model.Response, error)
	GetOrders(ctx context.Context, category, symbol, orderId, orderLinkId, orderStatus string, limit int) (*model.Response, error)
	AmendOrder(ctx context.Context, category, symbol, orderId, orderLinkId string, qty float64, price float64, options map[string]string) (*model.Response, error)

	// 仓位管理API
	GetPositions(ctx context.Context, category, symbol, settleCoin, positionIdx string) (*model.Response, error)
	SetLeverage(ctx context.Context, category, symbol string, buyLeverage, sellLeverage float64) (*model.Response, error)
	SetTradingStop(ctx context.Context, category, symbol string, takeProfit, stopLoss float64, options map[string]string) (*model.Response, error)
	SwitchPositionMode(ctx context.Context, category, symbol, mode string) (*model.Response, error)

	// 账户管理API
	GetWalletBalance(ctx context.Context, accountType, coin string) (*model.Response, error)
	GetFeeRate(ctx context.Context, category, symbol string) (*model.Response, error)
	GetAccountInfo(ctx context.Context) (*model.Response, error)
	SetMarginMode(ctx context.Context, marginMode string) (*model.Response, error)

	// 资产管理API
	GetCoinBalance(ctx context.Context, coin, accountType string) (*model.Response, error)
	TransferAsset(ctx context.Context, transferId, coin, amount, fromAccountType, toAccountType string) (*model.Response, error)
	GetTransferHistory(ctx context.Context, transferId, coin, status string, startTime, endTime int64, limit int) (*model.Response, error)
	Withdraw(ctx context.Context, coin, chain, address, tag, amount string, options map[string]string) (*model.Response, error)
}

// 以下是旧版接口定义，已被整合到BybitService接口中
// 保留此注释以便于理解代码演进历史