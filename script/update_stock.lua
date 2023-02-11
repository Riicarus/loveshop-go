local stock = redis.call("HGET", "commodity-stock", KEYS[1])

if stock ~= nil then
	stock = tostring(tonumber(stock) - tonumber(ARGV[1]))
	redis.call("HSET", "commodity-stock", KEYS[1],  stock)
else
	return "not exist"
end

return "success"