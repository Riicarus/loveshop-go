package connection

import (
	"time"

	"github.com/riicarus/loveshop/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var SqlConn *gorm.DB

func InitSqlConn() {
	var err error
	SqlConn, err = gorm.Open(mysql.Open(conf.ServiceConf.Gorm.Mysql.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err1 := SqlConn.DB()
	if err1 != nil {
		panic(err1)
	}
	sqlDB.SetMaxIdleConns(conf.ServiceConf.Gorm.Pool.MaxIdle)
	sqlDB.SetMaxOpenConns(conf.ServiceConf.Gorm.Pool.MaxOpen)
	sqlDB.SetConnMaxLifetime(time.Duration(conf.ServiceConf.Gorm.Pool.MaxLifeTime))
}