package connection

import (
	"time"

	"github.com/riicarus/loveshop/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewSqlConn() (*gorm.DB, error) {
	conn, err := gorm.Open(mysql.Open(conf.ServiceConf.Gorm.Mysql.Dsn), &gorm.Config{
		//SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err1 := conn.DB()
	if err1 != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(conf.ServiceConf.Gorm.Pool.MaxIdle)
	sqlDB.SetMaxOpenConns(conf.ServiceConf.Gorm.Pool.MaxOpen)
	sqlDB.SetConnMaxLifetime(time.Duration(conf.ServiceConf.Gorm.Pool.MaxLifeTime))

	return conn, nil
}