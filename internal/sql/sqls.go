package sql

var (
	CommodityFindDetailByIsbn           = "SELECT * FROM commodity WHERE extension->>'$.ISBN' = ?"
	CommodityFindPage                   = "SELECT c.id, c.name, c.type, c.numbering, c.amount, c.price, c.extension FROM commodity AS c INNER JOIN (SELECT id FROM commodity LIMIT ?, ?) AS i ON i.`id` = c.`id`"
	CommodityFindPageByType             = "SELECT c.id, c.name, c.type, c.numbering, c.amount, c.price, c.extension FROM commodity AS c INNER JOIN (SELECT id FROM commodity WHERE `type` = ? LIMIT ?, ?) AS i ON i.`id` = c.`id`"
	CommodityFindPageByFuzzyName        = "SELECT c.id, c.name, c.type, c.numbering, c.amount, c.price, c.extension FROM commodity AS c INNER JOIN (SELECT id FROM commodity WHERE `name` like CONCAT('%', ?, '%') LIMIT ?, ?) AS i ON i.`id` = c.`id`"
	CommodityFindPageByFuzzyNameAndType = "SELECT c.id, c.name, c.type, c.numbering, c.amount, c.price, c.extension FROM commodity AS c INNER JOIN (SELECT id FROM commodity WHERE `type` = ? AND `name` like CONCAT('%', ?, '%') LIMIT ?, ?) AS i ON i.`id` = c.`id`"
)

var (
	OrderFindDetailAdminViewById                       = "SELECT o.id, o.user_id, IFNULL(u.`name`, 'OFFLINE') AS username, o.admin_id, IFNULL(a.`name`, 'ONLINE') AS adminname, o.time, o.commodities, o.payment, o.`status`, o.type FROM `order` AS o INNER JOIN (SELECT id FROM `order` WHERE `id` = ?) AS i ON i.id = o.id LEFT JOIN `user` AS u ON u.id = o.user_id LEFT JOIN `admin` AS a ON a.id = o.admin_id"
	OrderFindPageOrderByTimeDesc                       = "SELECT o.id, o.user_id, IFNULL(u.`name`, 'OFFLINE') AS username, o.admin_id, IFNULL(a.`name`, 'ONLINE') AS adminname, o.time, o.commodities, o.payment, o.`status`, o.type FROM `order` AS o INNER JOIN (SELECT id FROM `order` ORDER BY time DESC LIMIT ?, ?) AS i ON i.id = o.id LEFT JOIN `user` AS u ON u.id = o.user_id LEFT JOIN `admin` AS a ON a.id = o.admin_id"
	OrderFindPageOrderByTimeAsc                        = "SELECT o.id, o.user_id, IFNULL(u.`name`, 'OFFLINE') AS username, o.admin_id, IFNULL(a.`name`, 'ONLINE') AS adminname, o.time, o.commodities, o.payment, o.`status`, o.type FROM `order` AS o INNER JOIN (SELECT id FROM `order` ORDER BY time ASC LIMIT ?, ?) AS i ON i.id = o.id LEFT JOIN `user` AS u ON u.id = o.user_id LEFT JOIN `admin` AS a ON a.id = o.admin_id"
	OrderFindPageByStatusOrderByTimeDesc               = "SELECT o.id, o.user_id, IFNULL(u.`name`, 'OFFLINE') AS username, o.admin_id, IFNULL(a.`name`, 'ONLINE') AS adminname, o.time, o.commodities, o.payment, o.`status`, o.type FROM `order` AS o INNER JOIN (SELECT id FROM `order` WHERE `status` = ? ORDER BY time DESC LIMIT ?, ?) AS i ON i.id = o.id LEFT JOIN `user` AS u ON u.id = o.user_id LEFT JOIN `admin` AS a ON a.id = o.admin_id"
	OrderFindPageByStatusOrderByTimeAsc                = "SELECT o.id, o.user_id, IFNULL(u.`name`, 'OFFLINE') AS username, o.admin_id, IFNULL(a.`name`, 'ONLINE') AS adminname, o.time, o.commodities, o.payment, o.`status`, o.type FROM `order` AS o INNER JOIN (SELECT id FROM `order` WHERE `status` = ? ORDER BY time ASC LIMIT ?, ?) AS i ON i.id = o.id LEFT JOIN `user` AS u ON u.id = o.user_id LEFT JOIN `admin` AS a ON a.id = o.admin_id"
	OrderFindUserViewPageByUidOrderByTimeDesc          = "SELECT o.id, o.user_id, o.time, o.commodities, o.payment FROM `order` AS o INNER JOIN (SELECT id FROM `order` WHERE `user_id` = ? ORDER BY time DESC LIMIT ?, ?) AS i ON i.id = o.id"
	OrderFindUserViewPageByUidOrderByTimeAsc           = "SELECT o.id, o.user_id, o.time, o.commodities, o.payment FROM `order` AS o INNER JOIN (SELECT id FROM `order` WHERE `user_id` = ? ORDER BY time Asc LIMIT ?, ?) AS i ON i.id = o.id"
	OrderFindUserViewPageByUidAndStatusOrderByTimeDesc = "SELECT o.id, o.user_id, o.time, o.commodities, o.payment FROM `order` AS o INNER JOIN (SELECT id FROM `order` WHERE `user_id` = ? AND `status` = ? ORDER BY time DESC LIMIT ?, ?) AS i ON i.id = o.id"
	OrderFindUserViewPageByUidAndStatusOrderByTimeAsc  = "SELECT o.id, o.user_id, o.time, o.commodities, o.payment FROM `order` AS o INNER JOIN (SELECT id FROM `order` WHERE `user_id` = ? AND `status` = ? ORDER BY time Asc LIMIT ?, ?) AS i ON i.id = o.id"
)

var (
	BillFindPageOrderByTimeDesc            = "SELECT b.id, b.time, b.admin_id, b.order_id, b.order_type FROM bill AS b INNER JOIN (SELECT id FROM bill ORDER BY time DESC LIMIT ?, ?) AS i ON i.`id` = b.`id`"
	BillFindPageOrderByTimeAsc             = "SELECT b.id, b.time, b.admin_id, b.order_id, b.order_type FROM bill AS b INNER JOIN (SELECT id FROM bill ORDER BY time ASC LIMIT ?, ?) AS i ON i.`id` = b.`id`"
	BillFindPageByOrderTypeOrderByTimeDesc = "SELECT b.id, b.time, b.admin_id, b.order_id, b.order_type FROM bill AS b INNER JOIN (SELECT id FROM bill WHERE `order_type` = ? ORDER BY time DESC LIMIT ?, ?) AS i ON i.`id` = b.`id`"
	BillFindPageByOrderTypeOrderByTimeAsc  = "SELECT b.id, b.time, b.admin_id, b.order_id, b.order_type FROM bill AS b INNER JOIN (SELECT id FROM bill WHERE `order_type` = ? ORDER BY time ASC LIMIT ?, ?) AS i ON i.`id` = b.`id`"
)
