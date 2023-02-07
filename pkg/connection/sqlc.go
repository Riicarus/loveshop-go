package connection

import (
	"sync"
	"time"

	"github.com/riicarus/loveshop/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewSqlConn() (*gorm.DB, error) {
	conn, err := gorm.Open(mysql.Open(conf.ServiceConf.Gorm.Mysql.Dsn), &gorm.Config{})
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

type TxContext struct {
	Tx   *gorm.DB

	txLock sync.RWMutex
	inTx bool
}

func NewTxContext() (*TxContext, error) {
	DB, err := NewSqlConn()
	if err != nil {
		return nil, err
	}
	return &TxContext{
		Tx: DB,
		inTx: false,
	}, nil
}

func (c *TxContext) StartTx() {
	c.txLock.Lock()
	c.inTx = true
	c.txLock.Unlock()
}

func (c *TxContext) IsInTx() bool {
	c.txLock.RLock()
	isIn := c.inTx
	c.txLock.RUnlock()
	return isIn
}
