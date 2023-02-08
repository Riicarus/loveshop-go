## 在 Gin 中 集成 Gorm 事务
Gorm 提供了比较方便的事务处理, 但是其中可能还有一些小毛病, 有待发现. 但是 Gorm 的事务是在单个 Model 中处理比较方便, 如果需要多个 Model 或者 Service 来共同协作完成一个事务, 就需要一些封装了.  

这里记录一个个人探索出的封装方法:  

### 业务封装
```golang
// 用于在 db.transaction() 中执行的事务方法
type TxFunc func(tx *gorm.DB) error

// 开启事务
func Transaction(db *gorm.DB, fcs []TxFunc) (bizErr, txErr error) {
	tx := db.Begin()

	for _, fc := range fcs {
		if bizErr = fc(tx); bizErr != nil {
			return bizErr, tx.Rollback().Error
		}
	}

	return bizErr, tx.Commit().Error
}

// model 需要实现接口
type IDBModel[T interface{}] interface {
	Conn(db *gorm.DB) T
}

// model 需要继承 DBModel
type DBModel struct {
	DB *gorm.DB
}
```

ServiceContext 贯穿全局 Service
```golang
type ServiceContext struct {
	Conf *conf.ServiceConfig

	// gorm
	DB *gorm.DB

	// model
	AdminModel     model.AdminModel
	UserModel      model.UserModel
	CommodityModel model.CommodityModel
	OrderModel     model.OrderModel
	BillModel      model.BillModel
}
```

### Model 层实现
Model 层就正常写就行, 主要是在 Service 岑处理
```golang
type BillModel interface {
	logic.IDBModel[BillModel]

	Add(bill *Bill) error

	FindById(id string) (*Bill, error)
	FindPageOrderByTime(desc bool, num, size int) ([]*Bill, error)
	FindPageByOrderTypeOrderByTime(orderType string, desc bool, num, size int) ([]*Bill, error)
}

type DefaultBillModel struct {
	logic.DBModel
}

func (m *DefaultBillModel) Conn(db *gorm.DB) BillModel {
	m.DB = db

	return m
}

// ... 实现 BillModel 的方法
```

### Service 层实现
```golang
type BillService struct {
	svcctx *context.ServiceContext
}
```

```golang
// use transaction to protect
func (s *OrderService) Add(ctx *gin.Context, param *dto.OrderAddParam) error {
	txfcs := make([]logic.TxFunc, 0)
	// add order
	txfcs = append(txfcs, s.AddTx(ctx, param))
	// decrease stock
	commodityService := NewCommodityService(s.svcctx)
	for _, c := range param.Commodities {
		txfcs = append(txfcs, commodityService.UpdateAmountTx(ctx, c.CommodityId, -c.Amount))
	}

	// 开启事务支持
	bizErr, txErr := logic.Transaction(s.svcctx.DB, txfcs)
	if bizErr != nil {
		return bizErr
	} else if txErr != nil {
		return txErr
	}

	return nil
}

func (s *OrderService) AddTx(ctx *gin.Context, param *dto.OrderAddParam) logic.TxFunc {
	return func(tx *gorm.DB) error {
        // ...

		err := s.svcctx.OrderModel.Conn(tx).Add(order)

        // ...
	}
}
```