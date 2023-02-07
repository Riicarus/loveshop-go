package connection

import (
	"fmt"
	"sync"
	"time"

	"github.com/riicarus/loveshop/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewSqlConn() (*gorm.DB, error) {
	conn, err := gorm.Open(mysql.Open(conf.ServiceConf.Gorm.Mysql.Dsn), &gorm.Config{
		SkipDefaultTransaction: true,
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

type TxContext struct {
	Db *gorm.DB

	txLock sync.RWMutex
	Tx     *gorm.DB
	inTx   bool
}

func NewTxContext() (*TxContext, error) {
	DB, err := NewSqlConn()
	if err != nil {
		return nil, err
	}
	return &TxContext{
		Db:   DB,
		inTx: false,
	}, nil
}

// get gorm db connection, if in transaction, return tx
func (c *TxContext) DB() *gorm.DB {
	var db *gorm.DB

	c.txLock.RLock()
	if c.inTx {
		db = c.Tx
	} else {
		db = c.Db
	}
	c.txLock.RUnlock()

	return db
}

func (c *TxContext) StartTx() {
	c.txLock.Lock()
	c.Tx = c.Db.Begin()
	c.inTx = true
	c.txLock.Unlock()
}

func (c *TxContext) CommitTx() {
	c.txLock.Lock()
	err := c.Db.Commit().Error
	if err != nil {
		fmt.Println(err)
	}
	c.Tx = nil
	c.inTx = false
	c.txLock.Unlock()
}

func (c *TxContext) RollBackTx() {
	c.txLock.Lock()
	c.Db.Rollback()
	c.Tx = nil
	c.inTx = false
	c.txLock.Unlock()
}

func (c *TxContext) IsInTx() bool {
	c.txLock.RLock()
	isIn := c.inTx
	c.txLock.RUnlock()
	return isIn
}
