# Bybit V5 API MCP服务

## 项目概述

本项目是基于Bybit V5 API的MCP（Microservice Communication Protocol）服务，旨在提供统一的接口来访问Bybit交易所的各种功能，包括现货、衍生品和期权交易。

## 功能特点

- **产品线统一**：统一了现货、衍生品和期权合约的API，通过单一接口即可交易不同产品
- **提高资金效率**：支持统一账户模式，允许在现货、USDT永续、USDC永续和期权合约之间共享和交叉使用资金
- **统一账户借贷**：支持统一账户模式下的跨产品借贷功能
- **投资组合保证金模式**：支持逆向永续、逆向期货、USDT永续、USDC永续、USDC期货和期权之间的组合保证金

## 项目结构

```
├── cmd/                  # 应用程序入口点
│   └── server/           # 服务器启动代码
├── internal/             # 内部包
│   ├── api/              # API处理层
│   │   ├── market/       # 市场数据API
│   │   ├── order/        # 订单管理API
│   │   ├── position/     # 仓位管理API
│   │   ├── account/      # 账户管理API
│   │   └── asset/        # 资产管理API
│   ├── client/           # Bybit API客户端
│   ├── config/           # 配置管理
│   ├── model/            # 数据模型
│   ├── service/          # 业务逻辑层
│   └── util/             # 工具函数
├── pkg/                  # 可重用的公共包
│   ├── bybitapi/         # Bybit API封装
│   ├── errors/           # 错误处理
│   └── logger/           # 日志处理
├── test/                 # 测试代码
├── go.mod                # Go模块定义
├── go.sum                # Go依赖校验
└── README.md             # 项目说明
```

## API模块

本服务实现了Bybit V5 API的主要模块：

1. **市场数据 (Market)**：K线图、订单簿、行情、平台交易数据等
2. **订单管理 (Order)**：创建、修改、取消订单等
3. **仓位管理 (Position)**：查询、修改仓位等
4. **账户管理 (Account)**：统一资金账户、费率等
5. **资产管理 (Asset)**：资产管理、资金管理等

## 技术栈

- Go语言
- gRPC/HTTP用于MCP服务通信
- 配置管理
- 日志系统
- 错误处理

## 使用方法

### 传统部署

详细的安装、配置和使用说明请参考[使用指南](USAGE.md)。

### Docker部署（推荐）

我们提供了完整的Docker部署解决方案，使您可以快速部署和运行服务，无需担心环境依赖问题。

```bash
# 直接在docker-compose.yml文件中配置环境变量
# 编辑docker-compose.yml文件，在environment部分填入您的API密钥

# 构建并启动服务
docker-compose build
docker-compose up -d
```

详细的Docker部署说明请参考[Docker部署指南](DOCKER.md)。

## 开发计划

- [x] 项目结构设计
- [x] 基础框架搭建
- [x] API客户端实现
- [x] 服务层实现
- [x] API接口实现
- [x] 测试与文档
- [x] Docker化部署