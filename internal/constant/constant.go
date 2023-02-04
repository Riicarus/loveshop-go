package constant

const (
	ADMIN_LOGIN_TYPE string = "ADMIN"
	USER_LOGIN_TYPE  string = "USER"
)

const (
	REDIS_COMMODITY_HASH = "commodity"
)

func RedisCommodityHashKey(id string) string {
	return "commodity_" + id
}