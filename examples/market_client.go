package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/bybit-mcp/internal/service"
	"github.com/bybit-mcp/pkg/logger"
)

func main() {
	// 设置日志
	logger.SetDefaultLogger("info", "stdout")
	logger.Info("启动Bybit API示例客户端")

	// 从环境变量获取API密钥
	apiKey := os.Getenv("BYBIT_API_KEY")
	apiSecret := os.Getenv("BYBIT_API_SECRET")

	if apiKey == "" || apiSecret == "" {
		logger.Fatal("请设置BYBIT_API_KEY和BYBIT_API_SECRET环境变量")
	}

	// 创建Bybit服务
	bybitService := service.NewBybitService(apiKey, apiSecret, "info", "stdout")

	// 创建上下文
	ctx := context.Background()

	// 获取BTC/USDT的K线数据
	logger.Info("获取BTC/USDT的K线数据")
	resp, err := bybitService.GetKline(ctx, "spot", "BTCUSDT", "1h", 10, 0, 0)
	if err != nil {
		logger.Error("获取K线数据失败: %v", err)
		return
	}

	// 打印响应
	printResponse("K线数据", resp)

	// 获取BTC/USDT的订单簿
	logger.Info("获取BTC/USDT的订单簿")
	resp, err = bybitService.GetOrderbook(ctx, "spot", "BTCUSDT", 5)
	if err != nil {
		logger.Error("获取订单簿失败: %v", err)
		return
	}

	// 打印响应
	printResponse("订单簿", resp)

	// 获取BTC/USDT的行情数据
	logger.Info("获取BTC/USDT的行情数据")
	resp, err = bybitService.GetTickers(ctx, "spot", "BTCUSDT")
	if err != nil {
		logger.Error("获取行情数据失败: %v", err)
		return
	}

	// 打印响应
	printResponse("行情数据", resp)

	// 获取交易对信息
	logger.Info("获取交易对信息")
	resp, err = bybitService.GetInstruments(ctx, "spot", "BTCUSDT", "")
	if err != nil {
		logger.Error("获取交易对信息失败: %v", err)
		return
	}

	// 打印响应
	printResponse("交易对信息", resp)

	logger.Info("示例客户端运行完成")
}

// 打印响应
func printResponse(title string, resp interface{}) {
	fmt.Printf("\n===== %s =====\n", title)
	jsonData, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		fmt.Printf("序列化响应失败: %v\n", err)
		return
	}
	fmt.Println(string(jsonData))
}