package sql

var CommodityFindInfoPage = "SELECT c.id, c.name, c.type, c.numbering, c.amount, c.price, c.extension FROM commodity AS c INNER JOIN (SELECT id FROM commodity LIMIT ?, ?) AS i ON i.`id` = c.`id`"