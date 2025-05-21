# 多阶段构建以减小最终镜像大小
FROM golang:1.20-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum文件
COPY go.mod ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 修复main.go中的错误（config.Port应该是cfg.Server.Port）
RUN sed -i 's/config.Port/cfg.Server.Port/g' cmd/server/main.go

# 编译应用
RUN CGO_ENABLED=0 GOOS=linux go build -o bybit-mcp ./cmd/server

# 使用轻量级的alpine镜像作为运行环境
FROM alpine:latest

# 安装CA证书，用于HTTPS请求
RUN apk --no-cache add ca-certificates tzdata

# 设置时区为亚洲/上海
ENV TZ=Asia/Shanghai

# 创建非root用户
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# 创建工作目录并设置权限
WORKDIR /app
RUN chown -R appuser:appgroup /app

# 从构建阶段复制编译好的二进制文件
COPY --from=builder /app/bybit-mcp /app/bybit-mcp

# 复制配置文件模板
COPY config.json /app/config.json

# 切换到非root用户
USER appuser

# 暴露gRPC端口
EXPOSE 50051

# 设置健康检查
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD nc -z localhost 50051 || exit 1

# 启动应用
CMD ["/app/bybit-mcp", "--config=/app/config.json"]