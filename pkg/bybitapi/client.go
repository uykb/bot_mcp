package bybitapi

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	// API基础URL
	BaseURL = "https://api.bybit.com"
	
	// API版本
	APIVersion = "v5"

	// 产品类别
	CategorySpot    = "spot"    // 现货
	CategoryLinear  = "linear"  // USDT永续
	CategoryInverse = "inverse" // 反向合约
	CategoryOption  = "option"  // 期权
)

// Client 是Bybit API客户端
type Client struct {
	BaseURL    string
	APIKey     string
	APISecret  string
	HTTPClient *http.Client
	Debug      bool
}

// NewClient 创建一个新的Bybit API客户端
func NewClient(apiKey, apiSecret string) *Client {
	return &Client{
		BaseURL:    BaseURL,
		APIKey:     apiKey,
		APISecret:  apiSecret,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Debug:      false,
	}
}

// SetDebug 设置调试模式
func (c *Client) SetDebug(debug bool) {
	c.Debug = debug
}

// 生成签名
func (c *Client) generateSignature(params map[string]string, timestamp int64) string {
	// 按键排序
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 构建查询字符串
	var queryString strings.Builder
	queryString.WriteString(strconv.FormatInt(timestamp, 10))
	queryString.WriteString(c.APIKey)

	for _, k := range keys {
		queryString.WriteString(k)
		queryString.WriteString(params[k])
	}

	// 计算HMAC-SHA256签名
	h := hmac.New(sha256.New, []byte(c.APISecret))
	h.Write([]byte(queryString.String()))
	return hex.EncodeToString(h.Sum(nil))
}

// 发送请求
func (c *Client) sendRequest(method, endpoint string, params map[string]string, auth bool) ([]byte, error) {
	// 构建URL
	apiURL := fmt.Sprintf("%s/%s/%s", c.BaseURL, APIVersion, endpoint)

	var req *http.Request
	var err error

	if method == "GET" {
		// 构建查询参数
		if len(params) > 0 {
			queryParams := url.Values{}
			for k, v := range params {
				queryParams.Add(k, v)
			}
			apiURL = fmt.Sprintf("%s?%s", apiURL, queryParams.Encode())
		}

		req, err = http.NewRequest(method, apiURL, nil)
		if err != nil {
			return nil, err
		}
	} else {
		// 构建请求体
		jsonParams, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}

		req, err = http.NewRequest(method, apiURL, bytes.NewBuffer(jsonParams))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
	}

	// 添加认证头
	if auth {
		timestamp := time.Now().UnixMilli()
		signature := c.generateSignature(params, timestamp)

		req.Header.Set("X-BAPI-API-KEY", c.APIKey)
		req.Header.Set("X-BAPI-TIMESTAMP", strconv.FormatInt(timestamp, 10))
		req.Header.Set("X-BAPI-SIGN", signature)
	}

	// 发送请求
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API错误: %s, 状态码: %d", string(body), resp.StatusCode)
	}

	return body, nil
}

// 市场数据API

// GetKline 获取K线数据
func (c *Client) GetKline(category, symbol, interval string, limit int) ([]byte, error) {
	params := map[string]string{
		"category": category,
		"symbol":   symbol,
		"interval": interval,
	}

	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	return c.sendRequest("GET", "market/kline", params, false)
}

// GetOrderbook 获取订单簿
func (c *Client) GetOrderbook(category, symbol string, limit int) ([]byte, error) {
	params := map[string]string{
		"category": category,
		"symbol":   symbol,
	}

	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	return c.sendRequest("GET", "market/orderbook", params, false)
}

// GetTickers 获取行情数据
func (c *Client) GetTickers(category, symbol string) ([]byte, error) {
	params := map[string]string{
		"category": category,
	}

	if symbol != "" {
		params["symbol"] = symbol
	}

	return c.sendRequest("GET", "market/tickers", params, false)
}

// 订单API

// CreateOrder 创建订单
func (c *Client) CreateOrder(category, symbol, side, orderType string, qty float64, price float64, timeInForce string) ([]byte, error) {
	params := map[string]string{
		"category":  category,
		"symbol":    symbol,
		"side":      side,
		"orderType": orderType,
		"qty":       strconv.FormatFloat(qty, 'f', -1, 64),
	}

	if price > 0 {
		params["price"] = strconv.FormatFloat(price, 'f', -1, 64)
	}

	if timeInForce != "" {
		params["timeInForce"] = timeInForce
	}

	return c.sendRequest("POST", "order/create", params, true)
}

// CancelOrder 取消订单
func (c *Client) CancelOrder(category, symbol, orderId string) ([]byte, error) {
	params := map[string]string{
		"category": category,
		"symbol":   symbol,
	}

	if orderId != "" {
		params["orderId"] = orderId
	}

	return c.sendRequest("POST", "order/cancel", params, true)
}

// GetOrders 获取订单列表
func (c *Client) GetOrders(category, symbol string, limit int) ([]byte, error) {
	params := map[string]string{
		"category": category,
	}

	if symbol != "" {
		params["symbol"] = symbol
	}

	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	return c.sendRequest("GET", "order/history", params, true)
}

// 仓位API

// GetPositions 获取仓位
func (c *Client) GetPositions(category, symbol string) ([]byte, error) {
	params := map[string]string{
		"category": category,
	}

	if symbol != "" {
		params["symbol"] = symbol
	}

	return c.sendRequest("GET", "position/list", params, true)
}

// 账户API

// GetWalletBalance 获取钱包余额
func (c *Client) GetWalletBalance(accountType, coin string) ([]byte, error) {
	params := map[string]string{
		"accountType": accountType,
	}

	if coin != "" {
		params["coin"] = coin
	}

	return c.sendRequest("GET", "account/wallet-balance", params, true)
}

// 资产API

// GetAssetInfo 获取资产信息
func (c *Client) GetAssetInfo(accountType string) ([]byte, error) {
	params := map[string]string{}

	if accountType != "" {
		params["accountType"] = accountType
	}

	return c.sendRequest("GET", "asset/transfer/query-asset-info", params, true)
}