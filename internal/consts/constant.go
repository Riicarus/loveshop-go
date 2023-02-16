package consts

// auth
const (
	ADMIN_LOGIN_TYPE string = "ADMIN"
	USER_LOGIN_TYPE  string = "USER"
)

// commodity
const (
	REDIS_COMMODITY_INFO_HASH  = "commodity-info"
	REDIS_COMMODITY_STOCK_HASH = "commodity-stock"

	BOOK_TYPE                = "BOOK"
	CULTURAL_CREATIVITY_TYPE = "CULTURAL_CREATIVITY"
	DAILY_NECESSITY_TYPE     = "DAILY_NECESSITY"
	SPORTS_GOODS_TYPE        = "SPORTS_GOODS"
	BOARD_GAME_TYPE          = "BOARD_GAME"
)

func RedisCommodityInfoHashKey(id string) string {
	return "commodity_" + id
}

func RedisCommodityStockHashKey(id string) string {
	return "commodity_" + id
}

// order
const (
	ONLINE  = "ONLINE"
	OFFLINE = "OFFLINE"

	ORDER_STATUS_CREATED  = "CREATED"
	ORDER_STATUS_CANCELED = "CANCELED"
	ORDER_STATUS_PAYED    = "PAYED"
	ORDER_STATUS_FINISHED = "FINISHED"
)

// kafka
const (
	KAFKA_ORDER_TOPIC = "loveshop-order"
	KAFKA_LOG_TOPIC   = "loveshop-log"

	KAFKA_ORDER_GROUP = "loveshop-order-group"
)

// regexp
const (
	POSITION_REG = "^0[0-9]-[AB]-0[0-9]$"
	ISBN_REG     = "^(?:ISBN(?:-1[03])?:? )?(?=[0-9X]{10}$|(?=(?:[0-9]+[- ]){3})[- 0-9X]{13}$|97[89][0-9]{10}$|(?=(?:[0-9]+[- ]){4})[- 0-9]{17}$)(?:97[89][- ]?)?[0-9]{1,5}[- ]?[0-9]+[- ]?[0-9]+[- ]?[0-9X]$"
)
