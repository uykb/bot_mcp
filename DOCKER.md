# Bybit V5 API MCP服务 Docker部署指南

## 简介

本文档提供了使用Docker部署Bybit V5 API MCP服务的详细说明。通过Docker化，您可以更轻松地在各种环境中部署和运行该服务，无需担心依赖问题。

## 前提条件

- Docker 19.03或更高版本
- Docker Compose 1.27.0或更高版本
- 有效的Bybit API密钥和密钥

## 快速开始

### 1. 配置API密钥

在部署前，您需要配置Bybit API密钥。有两种方式：

#### 方式一：直接修改config.json文件

编辑项目根目录下的`config.json`文件，将您的API密钥填入：

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

#### 方式二：直接在docker-compose.yml中配置环境变量（推荐）

直接在`docker-compose.yml`文件的`environment`部分配置所需的环境变量：

```yaml
services:
  bybit-mcp:
    # 其他配置...
    environment:
      - TZ=Asia/Shanghai
      - BYBIT_API_KEY=您的API密钥
      - BYBIT_API_SECRET=您的API密钥
      - BYBIT_BASE_URL=https://api.bybit.com
      - BYBIT_DEBUG=false
      - SERVER_PORT=50051
      - LOGGER_LEVEL=info
```

使用此方法，您可以直接在启动容器时修改这些值，无需创建额外的配置文件。

### 2. 构建和启动服务

在项目根目录下运行以下命令：

```bash
# 构建Docker镜像
docker-compose build

# 启动服务
docker-compose up -d
```

服务将在后台启动，并监听50051端口。

### 3. 验证服务是否正常运行

```bash
# 查看容器状态
docker-compose ps

# 查看服务日志
docker-compose logs -f
```

## 自定义配置

### 修改端口

如果您需要修改默认的50051端口，可以在`docker-compose.yml`文件中更改端口映射：

```yaml
ports:
  - "8080:50051"  # 将主机的8080端口映射到容器的50051端口
```

### 使用不同的配置文件

您可以创建自定义配置文件，并通过卷挂载使用它：

```yaml
volumes:
  - ./my-custom-config.json:/app/config.json:ro
```

## 生产环境部署建议

### 安全性

1. **不要在代码或配置文件中硬编码API密钥**，使用环境变量或Docker secrets。
2. 限制容器的资源使用，在`docker-compose.yml`中添加资源限制：

```yaml
services:
  bybit-mcp:
    # 其他配置...
    deploy:
      resources:
        limits:
          cpus: '1.0'
          memory: 512M
```

### 可靠性

1. 使用`restart: always`或`restart: unless-stopped`确保服务自动重启。
2. 配置健康检查（已在Dockerfile中设置）。
3. 考虑使用Docker Swarm或Kubernetes进行更复杂的部署和扩展。

### 日志管理

默认情况下，日志输出到标准输出，可以通过Docker日志系统查看。如果需要持久化日志，可以：

1. 修改配置文件，将日志输出到文件。
2. 挂载日志目录：

```yaml
volumes:
  - ./logs:/app/logs
```

## 故障排除

### 常见问题

1. **容器无法启动**
   - 检查配置文件格式是否正确
   - 确认端口是否被占用
   - 查看Docker日志：`docker-compose logs bybit-mcp`

2. **API调用失败**
   - 验证API密钥是否正确
   - 检查网络连接
   - 启用调试模式：设置`"debug": true`或环境变量`BYBIT_DEBUG=true`

3. **性能问题**
   - 增加容器资源限制
   - 检查主机资源使用情况

## 更新服务

当有新版本发布时，按照以下步骤更新服务：

```bash
# 拉取最新代码
git pull

# 重新构建镜像
docker-compose build

# 重启服务
docker-compose up -d
```

## 附录：完整的环境变量列表

| 环境变量 | 说明 | 默认值 |
|---------|------|-------|
| BYBIT_API_KEY | Bybit API密钥 | - |
| BYBIT_API_SECRET | Bybit API密钥 | - |
| BYBIT_BASE_URL | Bybit API基础URL | https://api.bybit.com |
| BYBIT_DEBUG | 是否启用调试模式 | false |
| SERVER_PORT | 服务监听端口 | 50051 |
| LOGGER_LEVEL | 日志级别 | info |