for i, v in ipairs(KEYS) do
    local stock = redis.call("HGET", "commodity-stock", v)

    if stock ~= nil then
        if tonumber(stock) >= tonumber(ARGV[i]) then
            stock = tostring(tonumber(stock) - tonumber(ARGV[i]))
            redis.call("HSET", "commodity-stock", v,  stock)
        else
               return "not enough"
        end
    else
        return "not exist"
    end
end

return "success"