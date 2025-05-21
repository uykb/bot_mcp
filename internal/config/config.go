package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config 表示MCP服务的配置
type Config struct {
	// 服务配置
	Server ServerConfig `json:"server"`

	// Bybit API配置
	Bybit BybitConfig `json:"bybit"`

	// 日志配置
	Logger LoggerConfig `json:"logger"`
}

// ServerConfig 表示服务器配置
type ServerConfig struct {
	Host string `json:"host"` // 服务主机
	Port int    `json:"port"` // 服务端口
}

// BybitConfig 表示Bybit API配置
type BybitConfig struct {
	BaseURL   string `json:"baseUrl"`   // API基础URL
	APIKey    string `json:"apiKey"`    // API密钥
	APISecret string `json:"apiSecret"` // API密钥
	Debug     bool   `json:"debug"`     // 调试模式
}

// LoggerConfig 表示日志配置
type LoggerConfig struct {
	Level  string `json:"level"`  // 日志级别
	Output string `json:"output"` // 日志输出
}

// LoadConfig 从文件加载配置
func LoadConfig(filePath string) (*Config, error) {
	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("配置文件不存在: %s", filePath)
	}

	// 读取文件
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("无法读取配置文件: %v", err)
	}

	// 解析JSON
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("无法解析配置文件: %v", err)
	}

	return &config, nil
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Host: "0.0.0.0",
			Port: 50051,
		},
		Bybit: BybitConfig{
			BaseURL:   "https://api.bybit.com",
			APIKey:    "",
			APISecret: "",
			Debug:     false,
		},
		Logger: LoggerConfig{
			Level:  "info",
			Output: "stdout",
		},
	}
}

// SaveConfig 保存配置到文件
func SaveConfig(config *Config, filePath string) error {
	// 序列化为JSON
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("无法序列化配置: %v", err)
	}

	// 写入文件
	if err := ioutil.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("无法写入配置文件: %v", err)
	}

	return nil
}