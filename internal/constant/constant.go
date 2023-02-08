package constant

// auth
const (
	ADMIN_LOGIN_TYPE string = "ADMIN"
	USER_LOGIN_TYPE  string = "USER"
)

// commodity
const (
	REDIS_COMMODITY_HASH = "commodity"

	BOOK_TYPE                = "BOOK"
	CULTURAL_CREATIVITY_TYPE = "CULTURAL_CREATIVITY"
	DAILY_NECESSITY_TYPE     = "DAILY_NECESSITY"
	SPORTS_GOODS_TYPE        = "SPORTS_GOODS"
	BOARD_GAME_TYPE          = "BOARD_GAME"
)

func RedisCommodityHashKey(id string) string {
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
