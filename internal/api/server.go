package api

import (
	"context"
	"encoding/json"

	"github.com/bybit-mcp/internal/model"
	"github.com/bybit-mcp/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// BybitMCPServer 实现了BybitMCPServiceServer接口
type BybitMCPServer struct {
	UnimplementedBybitMCPServiceServer
	service service.BybitService
}

// NewBybitMCPServer 创建一个新的Bybit MCP服务器
func NewBybitMCPServer(service service.BybitService) *BybitMCPServer {
	return &BybitMCPServer{
		service: service,
	}
}

// 将响应转换为gRPC响应格式
func (s *BybitMCPServer) toMCPResponse(requestID string, resp *model.Response, err error) (*MCPResponse, error) {
	if err != nil {
		return &MCPResponse{
			RequestId: requestID,
			Code:      int32(codes.Internal),
			Message:   err.Error(),
		}, nil
	}

	// 序列化响应数据
	data, err := json.Marshal(resp)
	if err != nil {
		return &MCPResponse{
			RequestId: requestID,
			Code:      int32(codes.Internal),
			Message:   "序列化响应失败: " + err.Error(),
		}, nil
	}

	return &MCPResponse{
		RequestId: requestID,
		Code:      int32(resp.RetCode),
		Message:   resp.RetMsg,
		Data:      data,
	}, nil
}

// ==================== 市场数据API实现 ====================

// GetKline 获取K线数据
func (s *BybitMCPServer) GetKline(ctx context.Context, req *KlineRequest) (*MCPResponse, error) {
	resp, err := s.service.GetKline(ctx, req.Category, req.Symbol, req.Interval, int(req.Limit))
	return s.toMCPResponse(req.RequestId, resp, err)
}

// GetOrderbook 获取订单簿
func (s *BybitMCPServer) GetOrderbook(ctx context.Context, req *OrderbookRequest) (*MCPResponse, error) {
	resp, err := s.service.GetOrderbook(ctx, req.Category, req.Symbol, int(req.Limit))
	return s.toMCPResponse(req.RequestId, resp, err)
}

// GetTickers 获取行情数据
func (s *BybitMCPServer) GetTickers(ctx context.Context, req *TickersRequest) (*MCPResponse, error) {
	resp, err := s.service.GetTickers(ctx, req.Category, req.Symbol)
	return s.toMCPResponse(req.RequestId, resp, err)
}

// GetRecentTrades 获取最近成交
func (s *BybitMCPServer) GetRecentTrades(ctx context.Context, req *RecentTradesRequest) (*MCPResponse, error) {
	resp, err := s.service.GetRecentTrades(ctx, req.Category, req.Symbol, int(req.Limit))
	return s.toMCPResponse(req.RequestId, resp, err)
}

// ==================== 订单管理API实现 ====================

// CreateOrder 创建订单
func (s *BybitMCPServer) CreateOrder(ctx context.Context, req *CreateOrderRequest) (*MCPResponse, error) {
	// 构建订单请求
	orderReq := &model.OrderRequest{
		Category:       req.Category,
		Symbol:         req.Symbol,
		Side:           req.Side,
		OrderType:      req.OrderType,
		Qty:            req.Qty,
		Price:          req.Price,
		TimeInForce:    req.TimeInForce,
		OrderLinkId:    req.OrderLinkId,
		TakeProfit:     req.TakeProfit,
		StopLoss:       req.StopLoss,
		ReduceOnly:     req.ReduceOnly,
		CloseOnTrigger: req.CloseOnTrigger,
	}

	resp, err := s.service.CreateOrder(ctx, orderReq)
	return s.toMCPResponse(req.RequestId, resp, err)
}

// CancelOrder 取消订单
func (s *BybitMCPServer) CancelOrder(ctx context.Context, req *CancelOrderRequest) (*MCPResponse, error) {
	resp, err := s.service.CancelOrder(ctx, req.Category, req.Symbol, req.OrderId)
	return s.toMCPResponse(req.RequestId, resp, err)
}

// GetOrders 获取订单列表
func (s *BybitMCPServer) GetOrders(ctx context.Context, req *GetOrdersRequest) (*MCPResponse, error) {
	resp, err := s.service.GetOrders(ctx, req.Category, req.Symbol, int(req.Limit))
	return s.toMCPResponse(req.RequestId, resp, err)
}

// GetOrderHistory 获取历史订单
func (s *BybitMCPServer) GetOrderHistory(ctx context.Context, req *GetOrderHistoryRequest) (*MCPResponse, error) {
	resp, err := s.service.GetOrderHistory(ctx, req.Category, req.Symbol, int(req.Limit))
	return s.toMCPResponse(req.RequestId, resp, err)
}

// CancelAllOrders 取消所有订单
func (s *BybitMCPServer) CancelAllOrders(ctx context.Context, req *CancelAllOrdersRequest) (*MCPResponse, error) {
	resp, err := s.service.CancelAllOrders(ctx, req.Category, req.Symbol, req.SettleCoin)
	return s.toMCPResponse(req.RequestId, resp, err)
}

// ==================== 仓位管理API实现 ====================

// GetPositions 获取仓位
func (s *BybitMCPServer) GetPositions(ctx context.Context, req *GetPositionsRequest) (*MCPResponse, error) {
	resp, err := s.service.GetPositions(ctx, req.Category, req.Symbol)
	return s.toMCPResponse(req.RequestId, resp, err)
}

// SetLeverage 设置杠杆
func (s *BybitMCPServer) SetLeverage(ctx context.Context, req *SetLeverageRequest) (*MCPResponse, error) {
	resp, err := s.service.SetLeverage(ctx, req.Category, req.Symbol, req.Leverage)
	return s.toMCPResponse(req.RequestId, resp, err)
}

// SetTpSlMode 设置止盈止损模式
func (s *BybitMCPServer) SetTpSlMode(ctx context.Context, req *SetTpSlModeRequest) (*MCPResponse, error) {
	resp, err := s.service.SetTpSlMode(ctx, req.Category, req.Symbol, req.TpSlMode)
	return s.toMCPResponse(req.RequestId, resp, err)
}

// SetRiskLimit 设置风险限额
func (s *BybitMCPServer) SetRiskLimit(ctx context.Context, req *SetRiskLimitRequest) (*MCPResponse, error) {
	resp, err := s.service.SetRiskLimit(ctx, req.Category, req.Symbol, int(req.RiskId))
	return s.toMCPResponse(req.RequestId, resp, err)
}

// ==================== 账户管理API实现 ====================

// GetWalletBalance 获取钱包余额
func (s *BybitMCPServer) GetWalletBalance(ctx context.Context, req *GetWalletBalanceRequest) (*MCPResponse, error) {
	resp, err := s.service.GetWalletBalance(ctx, req.AccountType, req.Coin)
	return s.toMCPResponse(req.RequestId, resp, err)
}

// GetAccountInfo 获取账户信息
func (s *BybitMCPServer) GetAccountInfo(ctx context.Context, req *GetAccountInfoRequest) (*MCPResponse, error) {
	resp, err := s.service.GetAccountInfo(ctx)
	return s.toMCPResponse(req.RequestId, resp, err)
}

// GetFeeRate 获取手续费率
func (s *BybitMCPServer) GetFeeRate(ctx context.Context, req *GetFeeRateRequest) (*MCPResponse, error) {
	resp, err := s.service.GetFeeRate(ctx, req.Category, req.Symbol)
	return s.toMCPResponse(req.RequestId, resp, err)
}

// GetAccountMode 获取账户模式
func (s *BybitMCPServer) GetAccountMode(ctx context.Context, req *GetAccountModeRequest) (*MCPResponse, error) {
	resp, err := s.service.GetAccountMode(ctx)
	return s.toMCPResponse(req.RequestId, resp, err)
}

// SetAccountMode 设置账户模式
func (s *BybitMCPServer) SetAccountMode(ctx context.Context, req *SetAccountModeRequest) (*MCPResponse, error) {
	resp, err := s.service.SetAccountMode(ctx, req.AccountMode)
	return s.toMCPResponse(req.RequestId, resp, err)
}

// ==================== 资产管理API实现 ====================

// GetAssetInfo 获取资产信息
func (s *BybitMCPServer) GetAssetInfo(ctx context.Context, req *GetAssetInfoRequest) (*MCPResponse, error) {
	resp, err := s.service.GetAssetInfo(ctx, req.AccountType)
	return s.toMCPResponse(req.RequestId, resp, err)
}

// AssetTransfer 资产划转
func (s *BybitMCPServer) AssetTransfer(ctx context.Context, req *AssetTransferRequest) (*MCPResponse, error) {
	resp, err := s.service.AssetTransfer(ctx, req.FromAccountType, req.ToAccountType, req.Coin, req.Amount)
	return s.toMCPResponse(req.RequestId, resp, err)
}

// GetTransferHistory 获取划转历史
func (s *BybitMCPServer) GetTransferHistory(ctx context.Context, req *GetTransferHistoryRequest) (*MCPResponse, error) {
	resp, err := s.service.GetTransferHistory(ctx, req.Coin, int(req.Limit))
	return s.toMCPResponse(req.RequestId, resp, err)
}

// GetDepositHistory 获取充值历史
func (s *BybitMCPServer) GetDepositHistory(ctx context.Context, req *GetDepositHistoryRequest) (*MCPResponse, error) {
	resp, err := s.service.GetDepositHistory(ctx, req.Coin, int(req.Limit))
	return s.toMCPResponse(req.RequestId, resp, err)
}

// GetWithdrawalHistory 获取提现历史
func (s *BybitMCPServer) GetWithdrawalHistory(ctx context.Context, req *GetWithdrawalHistoryRequest) (*MCPResponse, error) {
	resp, err := s.service.GetWithdrawalHistory(ctx, req.Coin, int(req.Limit))
	return s.toMCPResponse(req.RequestId, resp, err)
}