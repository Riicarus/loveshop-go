package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/riicarus/loveshop/internal/constant"
	"github.com/riicarus/loveshop/internal/context"
	"github.com/riicarus/loveshop/internal/entity/dto"
	"github.com/riicarus/loveshop/internal/model"
	"github.com/riicarus/loveshop/pkg/e"
	"github.com/riicarus/loveshop/pkg/util"
)

type OrderService struct {
	svcctx *context.ServiceContext
}

func NewOrderService(svcctx *context.ServiceContext) *OrderService {
	return &OrderService{
		svcctx: svcctx,
	}
}

func (s *OrderService) CastToDetailAdminView(order *model.Order) *dto.OrderDetailAdminView {
	return &dto.OrderDetailAdminView{
		Id: order.Id,
		AdminId: order.AdminId,
		UserId: order.UserId,
		Time: time.Unix(order.Time, 0).Format("2006-01-02 15:04:05"),
		Commodities: order.Commodities,
		Payment: order.Payment,
		Status: order.Status,
		Type: order.Type,
	}
}

func (s *OrderService) CastToDetailAdminViewSlice(orderSlice []*model.Order) []*dto.OrderDetailAdminView {
	viewSlice := make([]*dto.OrderDetailAdminView, 0)
	for _, o := range orderSlice {
		viewSlice = append(viewSlice, &dto.OrderDetailAdminView{
			Id: o.Id,
			AdminId: o.AdminId,
			UserId: o.UserId,
			Time: time.Unix(o.Time, 0).Format("2006-01-02 15:04:05"),
			Commodities: o.Commodities,
			Payment: o.Payment,
			Status: o.Status,
			Type: o.Type,
		})
	}

	return viewSlice
}

func (s *OrderService) CastToDetailUserView(order *model.Order) *dto.OrderDetailUserView {
	return &dto.OrderDetailUserView{
		Id: order.Id,
		UserId: order.UserId,
		Time: time.Unix(order.Time, 0).Format("2006-01-02 15:04:05"),
		Commodities: order.Commodities,
		Payment: order.Payment,
	}
}

func (s *OrderService) CastToDetailUserViewSlice(orderSlice []*model.Order) []*dto.OrderDetailUserView {
	viewSlice := make([]*dto.OrderDetailUserView, 0)
	for _, o := range orderSlice {
		viewSlice = append(viewSlice, &dto.OrderDetailUserView{
			Id: o.Id,
			UserId: o.UserId,
			Time: time.Unix(o.Time, 0).Format("2006-01-02 15:04:05"),
			Commodities: o.Commodities,
			Payment: o.Payment,
		})
	}

	return viewSlice
}

// use transaction to protect
func (s *OrderService) Add(ctx *gin.Context, param *dto.OrderAddParam) error {
	if strings.TrimSpace(param.AdminId) == "" && strings.TrimSpace(param.UserId) == "" {
		return e.VALIDATE_ERR
	}
	if len(param.Commodities) == 0 {
		return e.VALIDATE_ERR
	}
	if param.Type != constant.ONLINE && param.Type != constant.OFFLINE {
		return e.VALIDATE_ERR
	}

	if strings.TrimSpace(param.AdminId) == "" {
		param.AdminId = constant.ONLINE
	}
	if strings.TrimSpace(param.UserId) == "" {
		param.UserId = constant.OFFLINE
	}

	var payment float64
	for _, c := range param.Commodities{
		payment += c.Discount * c.Price * float64(c.Amount)
	}

	order := &model.Order{
		Id: uuid.New().String(),
		UserId: param.UserId,
		AdminId: param.AdminId,
		Time: time.Now().UnixMilli(),
		Commodities: param.Commodities,
		Payment: payment,
		Status: constant.ORDER_STATUS_CREATED,
		Type: param.Type,
	}

	ctx.Set("tx", true)

	err := s.svcctx.OrderModel.Add(order)
	if err != nil {
		fmt.Println("OrderService.Add(), database err: ", err)
		return err
	}

	// TODO decrease stock

	return nil
}

func (s *OrderService) CancleOrder(ctx *gin.Context, id string) error {
	err := s.svcctx.OrderModel.CancleOrder(id)
	if err != nil {
		fmt.Println("OrderService.CancleOrder(), database err: ", err)
		return err
	}

	return nil
}

func (s *OrderService) PayOrder(ctx *gin.Context, id string) error {
	err := s.svcctx.OrderModel.PayOrder(id)
	if err != nil {
		fmt.Println("OrderService.PayOrder(), database err: ", err)
		return err
	}

	return nil
}

func (s *OrderService) FinishOrder(ctx *gin.Context, id string) error {
	err := s.svcctx.OrderModel.FinishOrder(id)
	if err != nil {
		fmt.Println("OrderService.FinishOrder(), database err: ", err)
		return err
	}

	return nil
}

func (s *OrderService) FindDetailAdminViewPageOrderByTime(ctx *gin.Context, desc bool, page *util.Page[*dto.OrderDetailAdminView]) error {
	orderSlice, err := s.svcctx.OrderModel.FindPageOrderByTime(desc, page.Num, page.Size)
	if err != nil {
		fmt.Println("OrderService.FindDetailAdminViewPageOrderByTime(), err: ", err)
		return err
	}

	page.Data = s.CastToDetailAdminViewSlice(orderSlice)

	return nil
}

func (s *OrderService) FindDetailAdminViewPageByStatusOrderByTime(ctx *gin.Context, status string, desc bool, page *util.Page[*dto.OrderDetailAdminView]) error {
	orderSlice, err := s.svcctx.OrderModel.FindPageByStatusOrderByTime(status, desc, page.Num, page.Size)
	if err != nil {
		fmt.Println("OrderService.FindDetailAdminViewPageByStatusOrderByTime(), err: ", err)
		return err
	}

	page.Data = s.CastToDetailAdminViewSlice(orderSlice)

	return nil
}

func (s *OrderService) FindDetailUserViewPageByUidOrderByTime(ctx *gin.Context, uid string, desc bool, page *util.Page[*dto.OrderDetailUserView]) error {
	orderSlice, err := s.svcctx.OrderModel.FindUserViewPageByUidOrderByTime(uid, desc, page.Num, page.Size)
	if err != nil {
		fmt.Println("OrderService.FindDetailUserViewPageByUidOrderByTime(), err: ", err)
		return err
	}

	page.Data = s.CastToDetailUserViewSlice(orderSlice)

	return nil
}

func (s *OrderService) FindDetailUserViewPageByUidAndStatusOrderByTime(ctx *gin.Context, uid string, status string, desc bool, page *util.Page[*dto.OrderDetailUserView]) error {
	orderSlice, err := s.svcctx.OrderModel.FindUserViewPageByUidAndStatusOrderByTime(uid, status, desc, page.Num, page.Size)
	if err != nil {
		fmt.Println("OrderService.FindDetailUserViewPageByUidAndStatusOrderByTime(), err: ", err)
		return err
	}

	page.Data = s.CastToDetailUserViewSlice(orderSlice)

	return nil
}