package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/bybit-mcp/internal/service"
	"github.com/bybit-mcp/pkg/logger"
)

func main() {
	// 设置日志
	logger.SetDefaultLogger("info", "stdout")
	logger.Info("启动Bybit API订单管理示例客户端")

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

	// 查询账户余额
	logger.Info("查询账户余额")
	resp, err := bybitService.GetWalletBalance(ctx, "UNIFIED", "USDT")
	if err != nil {
		logger.Error("查询账户余额失败: %v", err)
		return
	}

	// 打印响应
	printResponse("账户余额", resp)

	// 创建限价买单
	symbol := "BTCUSDT"
	side := "Buy"
	orderType := "Limit"
	qty := 0.001
	price := 30000.0

	logger.Info("创建限价买单: %s %s %f @ %f", symbol, side, qty, price)
	
	// 创建订单可选参数
	options := map[string]string{
		"timeInForce": "GTC",
		"orderLinkId": "test-order-" + strconv.FormatInt(int64(price), 10),
	}

	resp, err = bybitService.CreateOrder(ctx, "spot", symbol, side, orderType, qty, price, options)
	if err != nil {
		logger.Error("创建订单失败: %v", err)
		return
	}

	// 打印响应
	printResponse("创建订单", resp)

	// 从响应中获取订单ID
	var orderId string
	if respData, ok := resp.Result.(map[string]interface{}); ok {
		if id, ok := respData["orderId"].(string); ok {
			orderId = id
		}
	}

	if orderId == "" {
		logger.Error("无法获取订单ID")
		return
	}

	// 查询订单状态
	logger.Info("查询订单状态: %s", orderId)
	resp, err = bybitService.GetOrders(ctx, "spot", symbol, orderId, "", "", 0)
	if err != nil {
		logger.Error("查询订单状态失败: %v", err)
		return
	}

	// 打印响应
	printResponse("订单状态", resp)

	// 修改订单价格
	newPrice := price * 0.99
	logger.Info("修改订单价格: %s -> %f", orderId, newPrice)
	resp, err = bybitService.AmendOrder(ctx, "spot", symbol, orderId, "", 0, newPrice, nil)
	if err != nil {
		logger.Error("修改订单价格失败: %v", err)
		return
	}

	// 打印响应
	printResponse("修改订单", resp)

	// 取消订单
	logger.Info("取消订单: %s", orderId)
	resp, err = bybitService.CancelOrder(ctx, "spot", symbol, orderId, "")
	if err != nil {
		logger.Error("取消订单失败: %v", err)
		return
	}

	// 打印响应
	printResponse("取消订单", resp)

	logger.Info("订单管理示例客户端运行完成")
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