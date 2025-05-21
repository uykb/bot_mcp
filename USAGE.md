# Bybit V5 API MCP服务使用指南

## 简介

本文档提供了Bybit V5 API MCP服务的安装、配置和使用说明。该服务基于Bybit V5 API，提供了统一的接口来访问Bybit交易所的各种功能，包括现货、衍生品和期权交易。

## 安装

### 前提条件

- Go 1.18或更高版本
- 有效的Bybit API密钥和密钥

### 获取代码

```bash
# 克隆仓库
git clone https://github.com/your-username/bybit-mcp.git
cd bybit-mcp

# 安装依赖
go mod download
```

## 配置

在运行服务前，需要配置API密钥和其他参数。配置文件为`config.json`，示例如下：

```json
{
  "server": {
    "host": "0.0.0.0",
    "port": 50051
  },
  "bybit": {
    "baseUrl": "https://api.bybit.com",
    "apiKey": "您的API密钥",
    "apiSecret": "您的API密钥",
    "debug": false
  },
  "logger": {
    "level": "info",
    "output": "stdout"
  }
}
```

请将`apiKey`和`apiSecret`替换为您的Bybit API密钥和密钥。

## 运行服务

```bash
# 编译服务
go build -o bybit-mcp ./cmd/server

# 运行服务（使用默认配置文件config.json）
./bybit-mcp

# 或者指定配置文件
./bybit-mcp --config=my-config.json

# 或者覆盖端口
./bybit-mcp --port=8080
```

服务启动后，将在指定端口（默认50051）监听gRPC请求。

## 使用示例

### 客户端示例

在`examples`目录中提供了一个简单的客户端示例，演示了如何连接到MCP服务并调用API：

```bash
# 运行示例客户端
go run ./examples/client.go
```

### API使用

本服务实现了Bybit V5 API的主要模块：

1. **市场数据 (Market)**：K线图、订单簿、行情、平台交易数据等
2. **订单管理 (Order)**：创建、修改、取消订单等
3. **仓位管理 (Position)**：查询、修改仓位等
4. **账户管理 (Account)**：统一资金账户、费率等
5. **资产管理 (Asset)**：资产管理、资金管理等

#### 市场数据示例

```go
// 获取K线数据
klineReq := &api.KlineRequest{
    RequestId: "req-1",
    Category:  "spot",
    Symbol:    "BTCUSDT",
    Interval:  "1h",
    Limit:     100,
}

klineResp, err := client.GetKline(ctx, klineReq)
```

#### 订单管理示例

```go
// 创建订单
orderReq := &api.CreateOrderRequest{
    RequestId:   "req-2",
    Category:    "spot",
    Symbol:      "BTCUSDT",
    Side:        "Buy",
    OrderType:   "Limit",
    Qty:         0.001,
    Price:       30000,
    TimeInForce: "GTC",
}

orderResp, err := client.CreateOrder(ctx, orderReq)
```

## 产品类别

Bybit V5 API支持以下产品类别：

- `spot`: 现货
- `linear`: USDT永续
- `inverse`: 反向合约
- `option`: 期权

在调用API时，需要指定正确的产品类别。

## 错误处理

API响应中包含了状态码和错误消息，可以通过检查这些信息来处理错误：

```go
if resp.Code != 0 {
    log.Printf("API错误: %s, 状态码: %d", resp.Message, resp.Code)
    // 处理错误
}
```

## 高级功能

### 统一账户模式

Bybit V5 API支持统一账户模式，允许在现货、USDT永续、USDC永续和期权合约之间共享和交叉使用资金。可以通过以下API设置账户模式：

```go
// 获取账户模式
modeReq := &api.GetAccountModeRequest{
    RequestId: "req-3",
}

modeResp, err := client.GetAccountMode(ctx, modeReq)

// 设置账户模式
setModeReq := &api.SetAccountModeRequest{
    RequestId:   "req-4",
    AccountMode: "UNIFIED",
}

setModeResp, err := client.SetAccountMode(ctx, setModeReq)
```

### 投资组合保证金模式

统一账户支持逆向永续、逆向期货、USDT永续、USDC永续、USDC期货和期权之间的组合保证金。

## 注意事项

1. 请确保API密钥具有足够的权限来执行所需的操作。
2. 在生产环境中使用时，建议启用TLS加密。
3. 请遵循Bybit的API使用限制和规则。

## 故障排除

如果遇到问题，可以尝试以下方法：

1. 启用调试模式（在配置文件中设置`debug: true`）
2. 检查日志输出
3. 确认API密钥和密钥是否正确
4. 确认网络连接是否正常

## 更多资源

- [Bybit V5 API文档](https://bybit-exchange.github.io/docs/v5/intro)
- [Go gRPC文档](https://grpc.io/docs/languages/go/)