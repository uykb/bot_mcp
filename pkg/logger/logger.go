package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// 日志级别
const (
	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
	FatalLevel = "fatal"
)

// 日志级别映射
var levelMap = map[string]int{
	DebugLevel: 0,
	InfoLevel:  1,
	WarnLevel:  2,
	ErrorLevel: 3,
	FatalLevel: 4,
}

// Logger 是日志记录器
type Logger struct {
	level  string
	output io.Writer
	logger *log.Logger
}

// Output 返回日志输出位置
func (l *Logger) Output() io.Writer {
	return l.output
}

// New 创建一个新的日志记录器
func New(level, output string) *Logger {
	var out io.Writer
	switch strings.ToLower(output) {
	case "stdout":
		out = os.Stdout
	case "stderr":
		out = os.Stderr
	default:
		// 尝试打开文件
		file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("无法打开日志文件 %s: %v, 使用标准输出\n", output, err)
			out = os.Stdout
		} else {
			out = file
		}
	}

	return &Logger{
		level:  strings.ToLower(level),
		output: out,
		logger: log.New(out, "", log.LstdFlags),
	}
}

// 检查是否应该记录该级别的日志
func (l *Logger) shouldLog(level string) bool {
	currentLevel, ok := levelMap[l.level]
	if !ok {
		currentLevel = levelMap[InfoLevel] // 默认为info级别
	}

	msgLevel, ok := levelMap[level]
	if !ok {
		msgLevel = levelMap[InfoLevel] // 默认为info级别
	}

	return msgLevel >= currentLevel
}

// 格式化日志消息
func (l *Logger) formatMessage(level, format string, args ...interface{}) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	message := fmt.Sprintf(format, args...)
	return fmt.Sprintf("%s [%s] %s", timestamp, strings.ToUpper(level), message)
}

// Debug 记录调试级别日志
func (l *Logger) Debug(format string, args ...interface{}) {
	if l.shouldLog(DebugLevel) {
		l.logger.Println(l.formatMessage(DebugLevel, format, args...))
	}
}

// Info 记录信息级别日志
func (l *Logger) Info(format string, args ...interface{}) {
	if l.shouldLog(InfoLevel) {
		l.logger.Println(l.formatMessage(InfoLevel, format, args...))
	}
}

// Warn 记录警告级别日志
func (l *Logger) Warn(format string, args ...interface{}) {
	if l.shouldLog(WarnLevel) {
		l.logger.Println(l.formatMessage(WarnLevel, format, args...))
	}
}

// Error 记录错误级别日志
func (l *Logger) Error(format string, args ...interface{}) {
	if l.shouldLog(ErrorLevel) {
		l.logger.Println(l.formatMessage(ErrorLevel, format, args...))
	}
}

// Fatal 记录致命级别日志并退出程序
func (l *Logger) Fatal(format string, args ...interface{}) {
	if l.shouldLog(FatalLevel) {
		l.logger.Println(l.formatMessage(FatalLevel, format, args...))
		os.Exit(1)
	}
}

// 默认日志记录器
var defaultLogger = New(InfoLevel, "stdout")

// SetDefaultLogger 设置默认日志记录器
func SetDefaultLogger(level, output string) {
	defaultLogger = New(level, output)
}

// Debug 使用默认日志记录器记录调试级别日志
func Debug(format string, args ...interface{}) {
	defaultLogger.Debug(format, args...)
}

// Info 使用默认日志记录器记录信息级别日志
func Info(format string, args ...interface{}) {
	defaultLogger.Info(format, args...)
}

// Warn 使用默认日志记录器记录警告级别日志
func Warn(format string, args ...interface{}) {
	defaultLogger.Warn(format, args...)
}

// Error 使用默认日志记录器记录错误级别日志
func Error(format string, args ...interface{}) {
	defaultLogger.Error(format, args...)
}

// Fatal 使用默认日志记录器记录致命级别日志并退出程序
func Fatal(format string, args ...interface{}) {
	defaultLogger.Fatal(format, args...)
}