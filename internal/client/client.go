package client

import (
	"github.com/bybit-mcp/internal/api/account"
	"github.com/bybit-mcp/internal/api/asset"
	"github.com/bybit-mcp/internal/api/market"
	"github.com/bybit-mcp/internal/api/order"
	"github.com/bybit-mcp/internal/api/position"
	"github.com/bybit-mcp/pkg/bybitapi"
	"github.com/bybit-mcp/pkg/logger"
)

// BybitClient 是Bybit API的客户端，整合了所有API模块
type BybitClient struct {
	// API客户端
	client *bybitapi.Client
	// 日志记录器
	logger *logger.Logger
	// API模块
	Market   *market.MarketAPI
	Order    *order.OrderAPI
	Position *position.PositionService
	Account  *account.AccountAPI
	Asset    *asset.AssetAPI
}

// ClientConfig 客户端配置
type ClientConfig struct {
	APIKey     string // API密钥
	APISecret  string // API密钥对应的密文
	Debug      bool   // 是否开启调试模式
	LogLevel   string // 日志级别
	LogOutput  string // 日志输出位置
}

// NewBybitClient 创建一个新的Bybit客户端
func NewBybitClient(config ClientConfig) *BybitClient {
	// 创建API客户端
	client := bybitapi.NewClient(config.APIKey, config.APISecret)
	client.SetDebug(config.Debug)

	// 创建日志记录器
	logLevel := config.LogLevel
	if logLevel == "" {
		logLevel = logger.InfoLevel
	}
	logOutput := config.LogOutput
	if logOutput == "" {
		logOutput = "stdout"
	}
	log := logger.New(logLevel, logOutput)

	// 创建API模块
	marketAPI := market.NewMarketAPI(client, log)
	orderAPI := order.NewOrderAPI(client, log)
	positionAPI := position.NewPositionService(client, logLevel, logOutput)
	accountAPI := account.NewAccountAPI(client, log)
	assetAPI := asset.NewAssetAPI(client, log)

	return &BybitClient{
		client:   client,
		logger:   log,
		Market:   marketAPI,
		Order:    orderAPI,
		Position: positionAPI,
		Account:  accountAPI,
		Asset:    assetAPI,
	}
}

// SetLogLevel 设置日志级别
func (c *BybitClient) SetLogLevel(level string) {
	c.logger = logger.New(level, c.logger.(*logger.Logger).Output())
}

// SetDebug 设置调试模式
func (c *BybitClient) SetDebug(debug bool) {
	c.client.SetDebug(debug)
}