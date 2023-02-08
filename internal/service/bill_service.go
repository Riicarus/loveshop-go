package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/riicarus/loveshop/internal/context"
	"github.com/riicarus/loveshop/internal/entity/dto"
	"github.com/riicarus/loveshop/internal/model"
	"github.com/riicarus/loveshop/pkg/e"
	"github.com/riicarus/loveshop/pkg/logic"
	"github.com/riicarus/loveshop/pkg/util"
	"gorm.io/gorm"
)

type BillService struct {
	svcctx *context.ServiceContext
}

func NewBillService(svcctx *context.ServiceContext) *BillService {
	return &BillService{
		svcctx: svcctx,
	}
}

func (s *BillService) CastToDetailAdminView(bill *model.Bill, admin *model.Admin, orderDetail *dto.OrderDetailAdminView) *dto.BillDetailAdminView {
	return &dto.BillDetailAdminView{
		Id:             bill.Id,
		BillCreateTime: time.Unix(bill.Time, 0).Format("2006-01-02 15:04:05"),
		OrderView:      orderDetail,
		AdminId:        admin.Id,
		Adminname:      admin.Name,
	}
}

func (s *BillService) CastToDetailAdminViewSlice(billSlice []*model.Bill, adminSlice []*model.Admin, orderDetailSlice []*dto.OrderDetailAdminView) []*dto.BillDetailAdminView {
	detailViewSlice := make([]*dto.BillDetailAdminView, 0, len(billSlice))

	for i, bill := range billSlice {
		detailViewSlice = append(detailViewSlice, s.CastToDetailAdminView(bill, adminSlice[i], orderDetailSlice[i]))
	}

	return detailViewSlice
}

func (s *BillService) Add(ctx *gin.Context, param *dto.BillAddParam) error {
	if strings.TrimSpace(param.AdminId) == "" || strings.TrimSpace(param.OrderId) == "" {
		return e.VALIDATE_ERR
	}

	bill := &model.Bill{
		Id:        uuid.New().String(),
		Time:      time.Now().UnixMilli(),
		AdminId:   param.AdminId,
		OrderId:   param.OrderId,
		OrderType: param.OrderType,
	}
	if err := s.svcctx.BillModel.Conn(s.svcctx.DB).Add(bill); err != nil {
		fmt.Println("BillService.Add(), database err: ", err)
		return err
	}

	return nil
}

func (s *BillService) AddTx(ctx *gin.Context, param *dto.BillAddParam) logic.TxFunc {
	return func(tx *gorm.DB) error {
		if strings.TrimSpace(param.AdminId) == "" || strings.TrimSpace(param.OrderId) == "" {
			return e.VALIDATE_ERR
		}

		bill := &model.Bill{
			Id:        uuid.New().String(),
			Time:      time.Now().UnixMilli(),
			AdminId:   param.AdminId,
			OrderId:   param.OrderId,
			OrderType: param.OrderType,
		}
		if err := s.svcctx.BillModel.Conn(tx).Add(bill); err != nil {
			fmt.Println("BillService.Add(), database err: ", err)
			return err
		}

		return nil
	}
}

func (s *BillService) FindDetailAdminViewById(ctx *gin.Context, id string) (*dto.BillDetailAdminView, error) {
	bill, err := s.svcctx.BillModel.Conn(s.svcctx.DB).FindById(id)
	if err != nil {
		fmt.Println("BillService.FindDetailAdminViewById(), database err: ", err)
		return nil, err
	}

	if bill == nil {
		return nil, nil
	}

	adminChan := make(chan *model.Admin)
	orderChan := make(chan *dto.OrderDetailAdminView)
	go func() {
		adminService := NewAdminService(s.svcctx)
		a, err2 := adminService.FindById(ctx.Copy(), bill.AdminId)
		if err2 != nil {
			fmt.Println("BillService.FindDetailAdminViewById(), admin service err: ", err2)
		}
		adminChan <- a
	}()

	go func() {
		orderSerivce := NewOrderService(s.svcctx)
		orderDetailView, err2 := orderSerivce.FindDetailAdminViewById(ctx.Copy(), bill.OrderId)
		if err2 != nil {
			fmt.Println("BillService.FindDetailAdminViewById(), order service err: ", err2)
		}
		orderChan <- orderDetailView
	}()

	orderDetail := <-orderChan
	admin := <-adminChan

	return s.CastToDetailAdminView(bill, admin, orderDetail), nil
}

func (s *BillService) FindDetailAdminViewPageOrderByTime(ctx *gin.Context, desc bool, page *util.Page[*dto.BillDetailAdminView]) error {
	billSlice, err := s.svcctx.BillModel.Conn(s.svcctx.DB).FindPageOrderByTime(desc, page.Num, page.Size)
	if err != nil {
		fmt.Println("BillService.FindDetailAdminViewPageOrderByTime(), database err: ", err)
		return err
	}

	if len(billSlice) == 0 {
		return nil
	}

	adminChan := make(chan []*model.Admin)
	orderChan := make(chan []*dto.OrderDetailAdminView)

	go func() {
		adminSlice := make([]*model.Admin, 0, len(billSlice))
		adminService := NewAdminService(s.svcctx)
		for _, bill := range billSlice {
			admin, err2 := adminService.FindById(ctx.Copy(), bill.AdminId)
			if err2 != nil {
				fmt.Println("BillService.FindDetailAdminViewById(), admin service err: ", err2)
			}
			adminSlice = append(adminSlice, admin)
		}

		adminChan <- adminSlice
	}()

	go func() {
		orderDetailSlice := make([]*dto.OrderDetailAdminView, 0, len(billSlice))
		orderSerivce := NewOrderService(s.svcctx)
		for _, bill := range billSlice {
			orderDetailView, err2 := orderSerivce.FindDetailAdminViewById(ctx.Copy(), bill.OrderId)
			if err2 != nil {
				fmt.Println("BillService.FindDetailAdminViewById(), order service err: ", err2)
			}
			orderDetailSlice = append(orderDetailSlice, orderDetailView)
		}

		orderChan <- orderDetailSlice
	}()

	adminSlice := <- adminChan
	orderDetailSlice := <-orderChan

	page.Data = s.CastToDetailAdminViewSlice(billSlice, adminSlice, orderDetailSlice)

	return nil
}

func (s *BillService) FindDetailAdminViewPageByOrderTypeOrderByTime(ctx *gin.Context, orderType string, desc bool, page *util.Page[*dto.BillDetailAdminView]) error {
	billSlice, err := s.svcctx.BillModel.Conn(s.svcctx.DB).FindPageByOrderTypeOrderByTime(orderType, desc, page.Num, page.Size)
	if err != nil {
		fmt.Println("BillService.FindDetailAdminViewPageByOrderTypeOrderByTime(), database err: ", err)
		return err
	}

	if len(billSlice) == 0 {
		return nil
	}

	adminChan := make(chan []*model.Admin)
	orderChan := make(chan []*dto.OrderDetailAdminView)

	go func() {
		adminSlice := make([]*model.Admin, 0, len(billSlice))
		adminService := NewAdminService(s.svcctx)
		for _, bill := range billSlice {
			admin, err2 := adminService.FindById(ctx.Copy(), bill.AdminId)
			if err2 != nil {
				fmt.Println("BillService.FindDetailAdminViewById(), admin service err: ", err2)
			}
			adminSlice = append(adminSlice, admin)
		}

		adminChan <- adminSlice
	}()

	go func() {
		orderDetailSlice := make([]*dto.OrderDetailAdminView, 0, len(billSlice))
		orderSerivce := NewOrderService(s.svcctx)
		for _, bill := range billSlice {
			orderDetailView, err2 := orderSerivce.FindDetailAdminViewById(ctx.Copy(), bill.OrderId)
			if err2 != nil {
				fmt.Println("BillService.FindDetailAdminViewById(), order service err: ", err2)
			}
			orderDetailSlice = append(orderDetailSlice, orderDetailView)
		}
		orderChan <- orderDetailSlice
	}()

	adminSlice := <- adminChan
	orderDetailSlice := <-orderChan

	page.Data = s.CastToDetailAdminViewSlice(billSlice, adminSlice, orderDetailSlice)

	return nil
}