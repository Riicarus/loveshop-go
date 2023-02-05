package sql

var CommodityFindDetailByIsbn = "SELECT * FROM commodity WHERE extension->>'$.ISBN' = ?"
var CommodityFindPage = "SELECT c.id, c.name, c.type, c.numbering, c.amount, c.price, c.extension FROM commodity AS c INNER JOIN (SELECT id FROM commodity LIMIT ?, ?) AS i ON i.`id` = c.`id`"
var CommodityFindPageByType = "SELECT c.id, c.name, c.type, c.numbering, c.amount, c.price, c.extension FROM commodity AS c INNER JOIN (SELECT id FROM commodity WHERE `type` = ? LIMIT ?, ?) AS i ON i.`id` = c.`id`"
var CommodityFindPageByFuzzyName = "SELECT c.id, c.name, c.type, c.numbering, c.amount, c.price, c.extension FROM commodity AS c INNER JOIN (SELECT id FROM commodity WHERE `name` like CONCAT('%', ?, '%') LIMIT ?, ?) AS i ON i.`id` = c.`id`"
var CommodityFindPageByFuzzyNameAndType = "SELECT c.id, c.name, c.type, c.numbering, c.amount, c.price, c.extension FROM commodity AS c INNER JOIN (SELECT id FROM commodity WHERE `type` = ? AND `name` like CONCAT('%', ?, '%') LIMIT ?, ?) AS i ON i.`id` = c.`id`"
