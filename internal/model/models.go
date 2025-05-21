package model

// 通用响应结构
type Response struct {
	RetCode    int         `json:"retCode"`    // 返回码
	RetMsg     string      `json:"retMsg"`     // 返回消息
	Result     interface{} `json:"result"`     // 结果数据
	RetExtInfo interface{} `json:"retExtInfo"` // 扩展信息
	Time       int64       `json:"time"`       // 时间戳
}

// 市场数据模型

// K线数据
type Kline struct {
	Category string   `json:"category"` // 产品类型
	Symbol   string   `json:"symbol"`   // 交易对
	List     [][]string `json:"list"`     // K线数据列表
}

// 订单簿数据
type Orderbook struct {
	Symbol string     `json:"symbol"` // 交易对
	Bids   [][]string `json:"bids"`   // 买单列表
	Asks   [][]string `json:"asks"`   // 卖单列表
	Ts     int64      `json:"ts"`     // 时间戳
	UpdateId int64    `json:"updateId"` // 更新ID
}

// 行情数据
type Ticker struct {
	Symbol        string `json:"symbol"`        // 交易对
	LastPrice     string `json:"lastPrice"`     // 最新价格
	IndexPrice    string `json:"indexPrice"`    // 指数价格
	MarkPrice     string `json:"markPrice"`     // 标记价格
	PrevPrice24h  string `json:"prevPrice24h"`  // 24小时前价格
	Price24hPcnt  string `json:"price24hPcnt"`  // 24小时价格变化百分比
	HighPrice24h  string `json:"highPrice24h"`  // 24小时最高价
	LowPrice24h   string `json:"lowPrice24h"`   // 24小时最低价
	PrevPrice1h   string `json:"prevPrice1h"`   // 1小时前价格
	OpenInterest  string `json:"openInterest"`  // 未平仓合约数量
	OpenInterestValue string `json:"openInterestValue"` // 未平仓合约价值
	Turnover24h   string `json:"turnover24h"`   // 24小时成交额
	Volume24h     string `json:"volume24h"`     // 24小时成交量
	FundingRate   string `json:"fundingRate"`   // 资金费率
	NextFundingTime string `json:"nextFundingTime"` // 下次资金费用时间
}

// 订单模型

// 订单请求
type OrderRequest struct {
	Category    string  `json:"category"`    // 产品类型
	Symbol      string  `json:"symbol"`      // 交易对
	Side        string  `json:"side"`        // 方向: Buy, Sell
	OrderType   string  `json:"orderType"`   // 订单类型: Limit, Market
	Qty         float64 `json:"qty"`         // 数量
	Price       float64 `json:"price,omitempty"` // 价格
	TimeInForce string  `json:"timeInForce,omitempty"` // 有效期: GTC, IOC, FOK
	OrderLinkId string  `json:"orderLinkId,omitempty"` // 自定义订单ID
	TakeProfit  float64 `json:"takeProfit,omitempty"`  // 止盈价格
	StopLoss    float64 `json:"stopLoss,omitempty"`    // 止损价格
	ReduceOnly  bool    `json:"reduceOnly,omitempty"`  // 只减仓
	CloseOnTrigger bool  `json:"closeOnTrigger,omitempty"` // 触发后平仓
}

// 订单信息
type Order struct {
	OrderId      string `json:"orderId"`      // 订单ID
	OrderLinkId  string `json:"orderLinkId"`  // 自定义订单ID
	Symbol       string `json:"symbol"`       // 交易对
	Side         string `json:"side"`         // 方向
	OrderType    string `json:"orderType"`    // 订单类型
	Price        string `json:"price"`        // 价格
	Qty          string `json:"qty"`          // 数量
	TimeInForce  string `json:"timeInForce"`  // 有效期
	OrderStatus  string `json:"orderStatus"`  // 订单状态
	CumExecQty   string `json:"cumExecQty"`   // 已成交数量
	CumExecValue string `json:"cumExecValue"` // 已成交价值
	CumExecFee   string `json:"cumExecFee"`   // 已成交手续费
	CreatedTime  string `json:"createdTime"`  // 创建时间
	UpdatedTime  string `json:"updatedTime"`  // 更新时间
}

// 仓位模型

// 仓位信息
type Position struct {
	PositionIdx    int    `json:"positionIdx"`    // 仓位索引
	RiskId         int    `json:"riskId"`         // 风险ID
	Symbol         string `json:"symbol"`         // 交易对
	Side           string `json:"side"`           // 方向
	Size           string `json:"size"`           // 仓位大小
	EntryPrice     string `json:"entryPrice"`     // 入场价格
	Leverage       string `json:"leverage"`       // 杠杆
	PositionValue  string `json:"positionValue"`  // 仓位价值
	PositionBalance string `json:"positionBalance"` // 仓位余额
	MarkPrice      string `json:"markPrice"`      // 标记价格
	PositionIM     string `json:"positionIM"`     // 仓位初始保证金
	PositionMM     string `json:"positionMM"`     // 仓位维持保证金
	TakeProfit     string `json:"takeProfit"`     // 止盈价格
	StopLoss       string `json:"stopLoss"`       // 止损价格
	UnrealisedPnl  string `json:"unrealisedPnl"`  // 未实现盈亏
	CumRealisedPnl string `json:"cumRealisedPnl"` // 累计已实现盈亏
	CreatedTime    string `json:"createdTime"`    // 创建时间
	UpdatedTime    string `json:"updatedTime"`    // 更新时间
}

// 账户模型

// 钱包余额
type WalletBalance struct {
	AccountType    string `json:"accountType"`    // 账户类型
	AccountIMRate  string `json:"accountIMRate"`  // 账户初始保证金率
	AccountMMRate  string `json:"accountMMRate"`  // 账户维持保证金率
	TotalEquity    string `json:"totalEquity"`    // 总权益
	TotalWalletBalance string `json:"totalWalletBalance"` // 总钱包余额
	TotalMarginBalance string `json:"totalMarginBalance"` // 总保证金余额
	TotalAvailableBalance string `json:"totalAvailableBalance"` // 总可用余额
	Coin          []CoinBalance `json:"coin"`          // 币种余额列表
}

// 币种余额
type CoinBalance struct {
	Coin           string `json:"coin"`           // 币种
	Equity         string `json:"equity"`         // 权益
	WalletBalance  string `json:"walletBalance"`  // 钱包余额
	AvailableToWithdraw string `json:"availableToWithdraw"` // 可提现余额
	AvailableToBorrow string `json:"availableToBorrow"` // 可借余额
	BorrowAmount    string `json:"borrowAmount"`    // 已借金额
	AccruedInterest string `json:"accruedInterest"` // 应计利息
	TotalOrderIM    string `json:"totalOrderIM"`    // 总订单初始保证金
	TotalPositionIM string `json:"totalPositionIM"` // 总仓位初始保证金
	TotalPositionMM string `json:"totalPositionMM"` // 总仓位维持保证金
}

// 资产模型

// 资产信息
type AssetInfo struct {
	Spot    AssetBalance `json:"spot"`    // 现货资产
	Contract AssetBalance `json:"contract"` // 合约资产
	Option   AssetBalance `json:"option"`   // 期权资产
}

// 资产余额
type AssetBalance struct {
	TotalEquity    string `json:"totalEquity"`    // 总权益
	AccountMargin  string `json:"accountMargin"`  // 账户保证金
	UnrealisedPnl  string `json:"unrealisedPnl"`  // 未实现盈亏
	RealisedPnl    string `json:"realisedPnl"`    // 已实现盈亏
	AvailableBalance string `json:"availableBalance"` // 可用余额
	CumRealisedPnl string `json:"cumRealisedPnl"` // 累计已实现盈亏
	TotalBalance   string `json:"totalBalance"`   // 总余额
}

// MCP服务请求/响应模型

// MCP请求
type MCPRequest struct {
	RequestId string      `json:"requestId"` // 请求ID
	Method    string      `json:"method"`    // 方法名
	Params    interface{} `json:"params"`    // 参数
}

// MCP响应
type MCPResponse struct {
	RequestId string      `json:"requestId"` // 请求ID
	Code      int         `json:"code"`      // 状态码
	Message   string      `json:"message"`   // 消息
	Data      interface{} `json:"data"`      // 数据
}