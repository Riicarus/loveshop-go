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
)

type OrderService struct {
	svcctx *context.ServiceContext
}

func NewOrderService(svcctx *context.ServiceContext) *OrderService {
	return &OrderService{
		svcctx: svcctx,
	}
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