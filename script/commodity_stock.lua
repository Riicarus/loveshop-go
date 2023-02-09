for i, v in ipairs(KEYS) do
    local stock = redis.call("HGET", "commodity-stock", v)

    if stock == true then
        if stock >= AMOUNTS[i] then
            stock = stock - AMOUNTS[i]
        else
            return false
        end
    else
        return false
    end
end

return true
