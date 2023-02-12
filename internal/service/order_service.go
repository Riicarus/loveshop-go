package service

import (
	c "context"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/riicarus/loveshop/internal/consts"
	"github.com/riicarus/loveshop/internal/context"
	"github.com/riicarus/loveshop/internal/entity/dto"
	"github.com/riicarus/loveshop/internal/model"
	"github.com/riicarus/loveshop/pkg/connection"
	"github.com/riicarus/loveshop/pkg/e"
	"github.com/riicarus/loveshop/pkg/logic"
	"github.com/riicarus/loveshop/pkg/util"
	"gorm.io/gorm"
)

type OrderService struct {
	svcctx *context.ServiceContext
}

func NewOrderService(svcctx *context.ServiceContext) *OrderService {
	return &OrderService{
		svcctx: svcctx,
	}
}

func (s *OrderService) CastToDetailAdminView(order *model.OrderDetail) *dto.OrderDetailAdminView {
	return &dto.OrderDetailAdminView{
		Id:          order.Id,
		AdminId:     order.AdminId,
		Adminname:   order.Adminname,
		UserId:      order.UserId,
		Username:    order.Username,
		Time:        time.Unix(order.Time, 0).Format("2006-01-02 15:04:05"),
		Timestamp:   order.Time,
		Commodities: order.Commodities,
		Payment:     order.Payment,
		Status:      order.Status,
		Type:        order.Type,
	}
}

func (s *OrderService) CastToDetailAdminViewSlice(orderSlice []*model.OrderDetail) []*dto.OrderDetailAdminView {
	viewSlice := make([]*dto.OrderDetailAdminView, 0)
	for _, o := range orderSlice {
		viewSlice = append(viewSlice, &dto.OrderDetailAdminView{
			Id:          o.Id,
			AdminId:     o.AdminId,
			Adminname:   o.Adminname,
			UserId:      o.UserId,
			Username:    o.Username,
			Time:        time.Unix(o.Time, 0).Format("2006-01-02 15:04:05"),
			Timestamp:   o.Time,
			Commodities: o.Commodities,
			Payment:     o.Payment,
			Status:      o.Status,
			Type:        o.Type,
		})
	}

	return viewSlice
}

func (s *OrderService) CastToDetailUserView(order *model.Order) *dto.OrderDetailUserView {
	return &dto.OrderDetailUserView{
		Id:          order.Id,
		UserId:      order.UserId,
		Time:        time.Unix(order.Time, 0).Format("2006-01-02 15:04:05"),
		Commodities: order.Commodities,
		Payment:     order.Payment,
	}
}

func (s *OrderService) CastToDetailUserViewSlice(orderSlice []*model.Order) []*dto.OrderDetailUserView {
	viewSlice := make([]*dto.OrderDetailUserView, 0)
	for _, o := range orderSlice {
		viewSlice = append(viewSlice, &dto.OrderDetailUserView{
			Id:          o.Id,
			UserId:      o.UserId,
			Time:        time.Unix(o.Time, 0).Format("2006-01-02 15:04:05"),
			Commodities: o.Commodities,
			Payment:     o.Payment,
		})
	}

	return viewSlice
}

// produce order msg to kafka
func (s *OrderService) ProduceToKafka(ctx *gin.Context, order *model.Order) error {
	// check commodity
	commodityService := NewCommodityService(s.svcctx)
	for _, c := range order.Commodities {
		detailView, err2 := commodityService.FindDetailViewById(ctx, c.CommodityId)
		if err2 != nil {
			return err2
		}

		if detailView.Amount < c.Amount {
			return e.STOCK_ERR
		}
	}

	kafkaHandler := connection.NewKakfaHandler[model.Order](consts.KAFKA_ORDER_GROUP, consts.KAFKA_ORDER_TOPIC)
	if err := kafkaHandler.Write(order.Id, *order); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// consumer order msg from kafka
// start when app start
func (s *OrderService) ConsumeFromKafka(ctx *gin.Context) {
	partitionConsume := func() {
		kafkaHandler := connection.NewKakfaHandler[model.Order](consts.KAFKA_ORDER_GROUP, consts.KAFKA_ORDER_TOPIC)
		for {
			order, commit, err := kafkaHandler.Fetch()
			if err != nil {
				fmt.Println("OrderService.ConsumerFromKafka(), get order from kafka err: ", err)
			}

			if err := s.Add(ctx.Copy(), order); err != nil {
				fmt.Println("OrderService.ConsumerFromKafka(), add order err: ", err)
			}
			if err := commit(); err != nil {
				fmt.Println("OrderService.ConsumerFromKafka(), commit to kafka err: ", err)
			}
		}
	}

	go partitionConsume()
	go partitionConsume()
	go partitionConsume()

	fmt.Println("kafka order consumer started")
}

// use LUA to ensure atomic redis operation
// LUA{find stock -> update stock if enough} -> push order to kafka
func (s *OrderService) Create(ctx *gin.Context, param *dto.OrderAddParam) error {
	if strings.TrimSpace(param.AdminId) == "" && strings.TrimSpace(param.UserId) == "" {
		return e.VALIDATE_ERR
	}
	if len(param.Commodities) == 0 {
		return e.VALIDATE_ERR
	}
	if param.Type != consts.ONLINE && param.Type != consts.OFFLINE {
		return e.VALIDATE_ERR
	}

	if strings.TrimSpace(param.AdminId) == "" {
		param.AdminId = consts.ONLINE
	}
	if strings.TrimSpace(param.UserId) == "" {
		param.UserId = consts.OFFLINE
	}

	var payment float64
	for _, c := range param.Commodities {
		payment += c.Discount * c.Price * float64(c.Amount)
	}

	order := model.Order{
		Id:          uuid.New().String(),
		UserId:      param.UserId,
		AdminId:     param.AdminId,
		Time:        time.Now().UnixMilli(),
		Commodities: param.Commodities,
		Payment:     payment,
		Status:      consts.ORDER_STATUS_CREATED,
		Type:        param.Type,
	}

	// KEYS: commodity_id, ARGV: number]
	desStock := redis.NewScript(`
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

	return "success"`)

	keys := make([]string, 0, len(order.Commodities))
	values := make([]interface{}, 0, len(order.Commodities))
	for _, c := range order.Commodities {
		keys = append(keys, consts.RedisCommodityStockHashKey(c.CommodityId))
		values = append(values, c.Amount)
	}

	rctx := c.Background()
	cmd := desStock.Run(rctx, connection.RedisClient, keys, values...)
	if err := cmd.Err(); err != nil {
		fmt.Println("OrderService.Create(), redis lua err: ", err)
		return err
	}

	switch cmd.Val() {
	case "success":
	case "not enough":
		return e.STOCK_ERR
	case "not exist":
		return e.NOT_EXIST_ERR
	}

	go func() {
		if err := s.ProduceToKafka(ctx, &order); err != nil {
			fmt.Printf("OrderService.Create(), kafka err: %v; Order info: %v", err, order)
		}
	}()

	return nil
}

func (s *OrderService) Add(ctx *gin.Context, order *model.Order) error {
	err := s.svcctx.OrderModel.Conn(s.svcctx.DB).Add(order)
	if err != nil {
		fmt.Println("OrderService.Add(), database err: ", err)
		return err
	}

	return nil
}

func (s *OrderService) CancleOrder(ctx *gin.Context, id string) error {
	err := s.svcctx.OrderModel.Conn(s.svcctx.DB).CancleOrder(id)
	if err != nil {
		fmt.Println("OrderService.CancleOrder(), database err: ", err)
		return err
	}

	return nil
}

func (s *OrderService) PayOrder(ctx *gin.Context, id string) error {
	err := s.svcctx.OrderModel.Conn(s.svcctx.DB).PayOrder(id)
	if err != nil {
		fmt.Println("OrderService.PayOrder(), database err: ", err)
		return err
	}

	return nil
}

func (s *OrderService) FinishOrder(ctx *gin.Context, id string) error {
	orderDetailView, err := s.FindDetailAdminViewById(ctx, id)
	if err != nil || orderDetailView == nil {
		fmt.Println("OrderService.FinishOrder(), database err: ", err)
		return err
	}

	txfcs := make([]logic.TxFunc, 0)
	// finish order
	txfcs = append(txfcs, s.FinishOrderTx(ctx, id))
	// create bill
	billService := NewBillService(s.svcctx)
	billAddParam := &dto.BillAddParam{
		Time:      time.Now().UnixMilli(),
		AdminId:   orderDetailView.AdminId,
		OrderId:   orderDetailView.Id,
		OrderType: orderDetailView.Type,
	}
	txfcs = append(txfcs, billService.AddTx(ctx, billAddParam))

	bizErr, txErr := logic.Transaction(s.svcctx.DB, txfcs)
	if bizErr != nil {
		return bizErr
	} else if txErr != nil {
		return txErr
	}

	return nil
}

func (s *OrderService) FinishOrderTx(ctx *gin.Context, id string) logic.TxFunc {
	return func(tx *gorm.DB) error {
		err := s.svcctx.OrderModel.Conn(tx).FinishOrder(id)
		if err != nil {
			fmt.Println("OrderService.FinishOrder(), database err: ", err)
			return err
		}

		return nil
	}
}

func (s *OrderService) FindDetailAdminViewById(ctx *gin.Context, id string) (*dto.OrderDetailAdminView, error) {
	orderDetail, err := s.svcctx.OrderModel.Conn(s.svcctx.DB).FindDetailById(id)
	if err != nil {
		return nil, nil
	}

	return s.CastToDetailAdminView(orderDetail), nil
}

func (s *OrderService) FindDetailAdminViewPageOrderByTime(ctx *gin.Context, desc bool, page *util.Page[*dto.OrderDetailAdminView]) error {
	orderSlice, err := s.svcctx.OrderModel.Conn(s.svcctx.DB).FindPageOrderByTime(desc, page.Num, page.Size)
	if err != nil {
		fmt.Println("OrderService.FindDetailAdminViewPageOrderByTime(), err: ", err)
		return err
	}

	page.Data = s.CastToDetailAdminViewSlice(orderSlice)

	return nil
}

func (s *OrderService) FindDetailAdminViewPageByStatusOrderByTime(ctx *gin.Context, status string, desc bool, page *util.Page[*dto.OrderDetailAdminView]) error {
	orderSlice, err := s.svcctx.OrderModel.Conn(s.svcctx.DB).FindPageByStatusOrderByTime(status, desc, page.Num, page.Size)
	if err != nil {
		fmt.Println("OrderService.FindDetailAdminViewPageByStatusOrderByTime(), err: ", err)
		return err
	}

	page.Data = s.CastToDetailAdminViewSlice(orderSlice)

	return nil
}

func (s *OrderService) FindDetailUserViewPageByUidOrderByTime(ctx *gin.Context, uid string, desc bool, page *util.Page[*dto.OrderDetailUserView]) error {
	orderSlice, err := s.svcctx.OrderModel.Conn(s.svcctx.DB).FindUserViewPageByUidOrderByTime(uid, desc, page.Num, page.Size)
	if err != nil {
		fmt.Println("OrderService.FindDetailUserViewPageByUidOrderByTime(), err: ", err)
		return err
	}

	page.Data = s.CastToDetailUserViewSlice(orderSlice)

	return nil
}

func (s *OrderService) FindDetailUserViewPageByUidAndStatusOrderByTime(ctx *gin.Context, uid string, status string, desc bool, page *util.Page[*dto.OrderDetailUserView]) error {
	orderSlice, err := s.svcctx.OrderModel.Conn(s.svcctx.DB).FindUserViewPageByUidAndStatusOrderByTime(uid, status, desc, page.Num, page.Size)
	if err != nil {
		fmt.Println("OrderService.FindDetailUserViewPageByUidAndStatusOrderByTime(), err: ", err)
		return err
	}

	page.Data = s.CastToDetailUserViewSlice(orderSlice)

	return nil
}
