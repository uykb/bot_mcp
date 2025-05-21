package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/bybit-mcp/internal/api/account"
	"github.com/bybit-mcp/internal/api/market"
	"github.com/bybit-mcp/internal/api/order"
	"github.com/bybit-mcp/internal/client"
	"github.com/bybit-mcp/pkg/logger"
)

func main() {
	// 从环境变量获取API密钥
	apiKey := os.Getenv("BYBIT_API_KEY")
	apiSecret := os.Getenv("BYBIT_API_SECRET")

	if apiKey == "" || apiSecret == "" {
		fmt.Println("请设置BYBIT_API_KEY和BYBIT_API_SECRET环境变量")
		os.Exit(1)
	}

	// 创建客户端配置
	config := client.ClientConfig{
		APIKey:    apiKey,
		APISecret: apiSecret,
		Debug:     true,
		LogLevel:  logger.DebugLevel,
		LogOutput: "stdout",
	}

	// 创建Bybit客户端
	bybitClient := client.NewBybitClient(config)

	// 创建上下文
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 示例1: 获取BTC/USDT的行情数据
	fmt.Println("获取BTC/USDT的行情数据:")
	tickersParams := market.TickerParams{
		Category: "spot",
		Symbol:   "BTCUSDT",
	}
	tickersResp, err := bybitClient.Market.GetTickers(ctx, tickersParams)
	if err != nil {
		fmt.Printf("获取行情数据失败: %v\n", err)
	} else {
		// 打印数据
		prettyJSON, _ := json.MarshalIndent(tickersResp, "", "  ")
		fmt.Printf("行情数据: %s\n", string(prettyJSON))
	}

	// 示例2: 获取账户钱包余额
	fmt.Println("\n获取账户钱包余额:")
	walletParams := account.GetWalletBalanceParams{
		AccountType: "UNIFIED",
		Coins:       []string{"BTC", "USDT"},
	}
	walletResp, err := bybitClient.Account.GetWalletBalance(ctx, walletParams)
	if err != nil {
		fmt.Printf("获取钱包余额失败: %v\n", err)
	} else {
		// 打印数据
		prettyJSON, _ := json.MarshalIndent(walletResp, "", "  ")
		fmt.Printf("钱包余额: %s\n", string(prettyJSON))
	}

	// 示例3: 获取仓位列表
	fmt.Println("\n获取仓位列表:")
	posResp, err := bybitClient.Position.GetPositions(ctx, "linear", "BTCUSDT", "", "")
	if err != nil {
		fmt.Printf("获取仓位列表失败: %v\n", err)
	} else {
		// 打印数据
		prettyJSON, _ := json.MarshalIndent(posResp, "", "  ")
		fmt.Printf("仓位列表: %s\n", string(prettyJSON))
	}

	// 示例4: 创建限价单(注意：此操作会实际下单，谨慎执行)
	fmt.Println("\n创建限价单示例(未执行):")
	orderParams := order.CreateOrderParams{
		Category:    "spot",
		Symbol:      "BTCUSDT",
		Side:        "Buy",
		OrderType:   "Limit",
		Qty:         "0.001",
		Price:       "20000",
		TimeInForce: "GTC",
	}
	fmt.Printf("订单参数示例: %+v\n", orderParams)
	// 实际下单时取消下面注释
	// orderResp, err := bybitClient.Order.CreateOrder(ctx, orderParams)
	// if err != nil {
	// 	fmt.Printf("创建订单失败: %v\n", err)
	// } else {
	// 	fmt.Printf("订单创建结果: %+v\n", orderResp)
	// }

	fmt.Println("\n客户端示例执行完成")
}