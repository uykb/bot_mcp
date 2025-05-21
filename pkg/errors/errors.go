package errors

import (
	"fmt"
	"net/http"
)

// 错误码定义
const (
	// 成功
	Success = 0

	// 通用错误
	ErrUnknown           = 10000 // 未知错误
	ErrInvalidParameter  = 10001 // 无效参数
	ErrInvalidSignature  = 10002 // 无效签名
	ErrAuthFailed        = 10003 // 认证失败
	ErrRateLimitExceeded = 10004 // 超过频率限制
	ErrPermissionDenied  = 10005 // 权限不足
	ErrServerError       = 10006 // 服务器错误
	ErrServiceUnavailable = 10007 // 服务不可用

	// API错误
	ErrAPIRequestFailed  = 20001 // API请求失败
	ErrAPIResponseInvalid = 20002 // API响应无效
	ErrAPITimeout        = 20003 // API超时

	// 业务错误
	ErrOrderFailed       = 30001 // 下单失败
	ErrOrderCancelFailed = 30002 // 取消订单失败
	ErrPositionNotFound  = 30003 // 仓位不存在
	ErrInsufficientBalance = 30004 // 余额不足
)

// Error 表示MCP服务错误
type Error struct {
	Code    int    // 错误码
	Message string // 错误消息
	Cause   error  // 原始错误
}

// Error 实现error接口
func (e *Error) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("[%d] %s: %v", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// Unwrap 返回原始错误
func (e *Error) Unwrap() error {
	return e.Cause
}

// New 创建一个新的错误
func New(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// Wrap 包装一个错误
func Wrap(code int, message string, err error) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Cause:   err,
	}
}

// FromHTTPResponse 从HTTP响应创建错误
func FromHTTPResponse(statusCode int, body string) *Error {
	switch statusCode {
	case http.StatusOK:
		return nil
	case http.StatusBadRequest:
		return New(ErrInvalidParameter, fmt.Sprintf("无效请求: %s", body))
	case http.StatusUnauthorized:
		return New(ErrAuthFailed, "认证失败")
	case http.StatusForbidden:
		return New(ErrPermissionDenied, "权限不足")
	case http.StatusTooManyRequests:
		return New(ErrRateLimitExceeded, "超过频率限制")
	case http.StatusInternalServerError:
		return New(ErrServerError, "服务器错误")
	case http.StatusServiceUnavailable:
		return New(ErrServiceUnavailable, "服务不可用")
	default:
		return New(ErrUnknown, fmt.Sprintf("未知错误: %d - %s", statusCode, body))
	}
}

// FromBybitAPIError 从Bybit API错误创建错误
func FromBybitAPIError(retCode int, retMsg string) *Error {
	if retCode == 0 {
		return nil
	}

	// 映射Bybit错误码到MCP错误码
	var code int
	switch {
	case retCode >= 10001 && retCode <= 10003:
		code = ErrInvalidParameter
	case retCode == 10004:
		code = ErrInvalidSignature
	case retCode == 10005:
		code = ErrAuthFailed
	case retCode == 10006 || retCode == 10007:
		code = ErrPermissionDenied
	case retCode == 10010:
		code = ErrRateLimitExceeded
	case retCode >= 20001 && retCode <= 20044:
		code = ErrOrderFailed
	case retCode >= 30000 && retCode <= 30099:
		code = ErrInsufficientBalance
	case retCode >= 110001 && retCode <= 110999:
		code = ErrAPIRequestFailed
	default:
		code = ErrUnknown
	}

	return New(code, fmt.Sprintf("Bybit API错误: [%d] %s", retCode, retMsg))
}