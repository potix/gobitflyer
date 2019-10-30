package types

type ProductCode string

type CurrencyCode string

type Pagination struct {
	Count  int64 `json:"count,omitempty" url:"count,omitempty"`
	Before int64 `json:"before,omitempty" url:"before,omitempty"`
	After  int64 `json:"after,omitempty" url:"after,omitempty"`
}

type OrderType string

const (
	OrderTypeLimit  OrderType = "LIMIT"
	OrderTypeMarket OrderType = "MARKET"
)

type Side string

const (
	SideBuy  Side = "BUY"
	SideSell Side = "SELL"
)

type TimeInForce string

const (
	TimeInForceNone TimeInForce = ""
	TimeInForceGTC  TimeInForce = "GTC"
	TimeInForceIOC  TimeInForce = "IOC"
	TimeInForceFOK  TimeInForce = "FOK"
)

type IdType int

const (
	IdTypeParentOrderId           IdType = 1
	IdTypeParentOrderAcceptanceId IdType = 2
	IdTypeChildOrderId            IdType = 3
	IdTypeChildOrderAcceptanceId  IdType = 4
)

type OrderState string

const (
	OrderStateNone      OrderState = ""
	OrderStateActive    OrderState = "ACTIVE"
	OrderStateCompleted OrderState = "COMPLETED"
	OrderStateCanceled  OrderState = "CANCELED"
	OrderStateExpired   OrderState = "EXPIRED"
	OrderStateRejected  OrderState = "REJECTED"
)

type TradeType string

const (
	TradeTypeBuy        TradeType = "BUY"
	TradeTypeSell       TradeType = "SELL"
	TradeTypeDeposit    TradeType = "DEPOSIT"
	TradeTypeWithdraw   TradeType = "WITHDRAW"
	TradeTypeFee        TradeType = "FEE"
	TradeTypePostCol    TradeType = "POST_COLL"
	TradeTypeCancelColl TradeType = "CANCEL_COLL"
	TradeTypePayment    TradeType = "PAYMENT"
	TradeTypeTransfer   TradeType = "TRANSFER"
)

type OrderMethod string

const (
	OrderMethodNone   OrderMethod = ""
	OrderMethodSimple OrderMethod = "SIMPLE"
	OrderMethodIFD    OrderMethod = "IFD"
	OrderMethodOCO    OrderMethod = "OCO"
	OrderMethodIFDOCO OrderMethod = "IFDOCO"
)

type ConditionType string

const (
	ConditionTypeLimit     ConditionType = "LIMIT"
	ConditionTypeMarket    ConditionType = "MARKET"
	ConditionTypeStop      ConditionType = "STOP"
	ConditionTypeStopLimit ConditionType = "STOP_LIMIT"
	ConditionTypeTrail     ConditionType = "TRAIL"
)

type RealtimeType int

const (
	RealtimeTypeBoardSnapshot RealtimeType = 1
	RealtimeTypeBoard         RealtimeType = 2
	RealtimeTypeTicker        RealtimeType = 3
	RealtimeTypeExecutions    RealtimeType = 4
)
