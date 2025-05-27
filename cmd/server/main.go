package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/bybit-mcp/internal/api"
	"github.com/bybit-mcp/internal/config"
	"github.com/bybit-mcp/internal/service"
	"google.golang.org/grpc"
)

// 命令行参数
type CommandLineArgs struct {
	ConfigFile string
	Port       int
}

func main() {
	// 解析命令行参数
	var args CommandLineArgs
	flag.StringVar(&args.ConfigFile, "config", "config.json", "配置文件路径")
	flag.IntVar(&args.Port, "port", 50051, "服务监听端口（覆盖配置文件）")
	flag.Parse()

	// 加载配置
	var cfg *config.Config
	var err error

	// 尝试加载配置文件
	cfg, err = config.LoadConfig(args.ConfigFile)
	if err != nil {
		log.Printf("无法加载配置文件: %v, 使用默认配置", err)
		cfg = config.DefaultConfig()
	}

	// 命令行参数覆盖配置文件
	if args.Port != 0 {
		cfg.Server.Port = args.Port
	}

	// 创建Bybit服务
	bybitService := service.NewBybitService(cfg.Bybit.APIKey, cfg.Bybit.APISecret)
	// 设置调试模式
	if cfg.Bybit.Debug {
		// 使用bybitapi客户端的调试模式
		log.Println("启用调试模式")
	}

	// 创建MCP服务器
	mcpServer := api.NewBybitMCPServer(bybitService)

	// 创建gRPC服务器
	server := grpc.NewServer()

	// 注册服务
	api.RegisterBybitMCPServiceServer(server, mcpServer)

	// 启动服务器
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		log.Fatalf("无法监听端口: %v", err)
	}

	log.Printf("Bybit MCP服务启动，监听端口: %d", cfg.Server.Port)

	// 在后台启动服务
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("服务启动失败: %v", err)
		}
	}()

	// 等待中断信号优雅地关闭服务器
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	log.Println("正在关闭服务...")
	// 优雅停止
	server.GracefulStop()
	log.Println("服务已关闭")
}