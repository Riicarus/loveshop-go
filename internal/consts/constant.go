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
	ORDER_STATUS_CANCLED  = "CANCLED"
	ORDER_STATUS_PAYED    = "PAYED"
	ORDER_STATUS_FINISHED = "FINISHED"
)

// kafka
const (
	KAFKA_ORDER_TOPIC = "loveshop-order"
	KAKFA_LOG_TOPIC   = "loveshop-log"

	KAFKA_ORDER_GROUP = "loveshop-order-group"
)
