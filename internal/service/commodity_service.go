package service

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/riicarus/loveshop/internal/constant"
	"github.com/riicarus/loveshop/internal/context"
	"github.com/riicarus/loveshop/internal/entity/dto"
	"github.com/riicarus/loveshop/internal/model"
	"github.com/riicarus/loveshop/pkg/connection"
	"github.com/riicarus/loveshop/pkg/e"
	"github.com/riicarus/loveshop/pkg/logic"
	"github.com/riicarus/loveshop/pkg/util"
)

type CommodityService struct {
	svcctx *context.ServiceContext
}

func NewCommodityService(svcctx *context.ServiceContext) *CommodityService {
	return &CommodityService{
		svcctx: svcctx,
	}
}

func (s *CommodityService) CastToSimpleView(commodity *model.Commodity) *dto.CommoditySimpleView {
	return &dto.CommoditySimpleView{
		Id:        commodity.Id,
		Type:      commodity.Type,
		Numbering: commodity.Numbering,
		Name:      commodity.Name,
		Amount:    commodity.Amount,
		Price:     commodity.Price,
		Extension: commodity.Extension,
	}
}

func (s *CommodityService) CastToSimpleViewSlice(commoditySlice []*model.Commodity) []*dto.CommoditySimpleView {
	simpleViewSlice := make([]*dto.CommoditySimpleView, 0, len(commoditySlice))
	for _, c := range commoditySlice {
		simpleViewSlice = append(simpleViewSlice, &dto.CommoditySimpleView{
			Id:        c.Id,
			Type:      c.Type,
			Numbering: c.Numbering,
			Name:      c.Name,
			Amount:    c.Amount,
			Price:     c.Price,
			Extension: c.Extension,
		})
	}

	return simpleViewSlice
}

func (s *CommodityService) CastToDetailView(commodity *model.Commodity) *dto.CommodityDetailView {
	return &dto.CommodityDetailView{
		Id:        commodity.Id,
		Type:      commodity.Type,
		Numbering: commodity.Numbering,
		Name:      commodity.Name,
		Amount:    commodity.Amount,
		Price:     commodity.Price,
		Extension: commodity.Extension,
		Deleted:   commodity.Deleted,
	}
}

func (s *CommodityService) Add(ctx *gin.Context, param *dto.CommodityAddParam) error {
	// TODO check param
	commodity := &model.Commodity{
		Id:        uuid.New().String(),
		Type:      param.Type,
		Numbering: param.Numbering,
		Name:      param.Name,
		Amount:    param.Amount,
		Price:     param.Price,
		Extension: param.Extension,
		Deleted:   false,
	}

	err := s.svcctx.CommodityModel.Conn(s.svcctx.DB).Add(commodity)
	if err != nil {
		fmt.Println("CommodityService.Add(), database err: ", err)
		return err
	}

	return nil
}

// status change remove cache from redis
func (s *CommodityService) Update(ctx *gin.Context, param *dto.CommodityUpdateParam) error {
	if strings.TrimSpace(param.Id) == "" {
		return e.UNAUTHED_ERR
	}

	commodity := &model.Commodity{
		Id:        param.Id,
		Numbering: param.Numbering,
		Name:      param.Name,
		Type:      param.Type,
		Price:     param.Price,
		Extension: param.Extension,
	}

	err := s.svcctx.CommodityModel.Conn(s.svcctx.DB).Update(commodity)
	if err != nil {
		fmt.Println("CommodityService.Update(), database err: ", err)
		return err
	}

	// use another routine to remove cache
	go func() {
		err = connection.NewRedisConnection[string]().DoHashRemove(constant.REDIS_COMMODITY_INFO_HASH, constant.RedisCommodityHashKey(param.Id))
		if err != nil {
			fmt.Println("CommodityService.Update(), redis err: ", err)
		}
	}()

	return nil
}

// status change remove cache from redis
func (s *CommodityService) UpdateAmount(ctx *gin.Context, id string, number int) error {
	if strings.TrimSpace(id) == "" {
		return e.VALIDATE_ERR
	}

	err := s.svcctx.CommodityModel.Conn(s.svcctx.DB).UpdateAmount(id, number)
	if err != nil {
		fmt.Println("CommodityService.UpdateAmount(), database err: ", err)
		return err
	}

	// use another routine to remove cache
	go func() {
		err = connection.NewRedisConnection[string]().DoHashRemove(constant.REDIS_COMMODITY_INFO_HASH, constant.RedisCommodityHashKey(id))
		if err != nil {
			fmt.Println("CommodityService.UpdateAmount(), redis err: ", err)
		}
	}()

	return err
}

func (s *CommodityService) UpdateAmountTx(ctx *gin.Context, id string, number int) logic.TxFunc {
	return func(tx *gorm.DB) error {
		if strings.TrimSpace(id) == "" {
			return e.VALIDATE_ERR
		}

		err := s.svcctx.CommodityModel.Conn(tx).UpdateAmount(id, number)
		if err != nil {
			fmt.Println("CommodityService.UpdateAmountTx(), database err: ", err)
			return err
		}

		// use another routine to remove cache
		go func() {
			err = connection.NewRedisConnection[string]().DoHashRemove(constant.REDIS_COMMODITY_INFO_HASH, constant.RedisCommodityHashKey(id))
			if err != nil {
				fmt.Println("CommodityService.UpdateAmountTx(), redis err: ", err)
			}
		}()

		return err
	}
}

// status change remove cache from redis
func (s *CommodityService) Delete(ctx *gin.Context, id string) error {
	if strings.TrimSpace(id) == "" {
		return e.VALIDATE_ERR
	}

	err := s.svcctx.CommodityModel.Conn(s.svcctx.DB).Delete(id)
	if err != nil {
		fmt.Println("CommodityService.Delete(), database err: ", err)
		return err
	}

	// use another routine to remove cache
	go func() {
		err = connection.NewRedisConnection[string]().DoHashRemove(constant.REDIS_COMMODITY_INFO_HASH, constant.RedisCommodityHashKey(id))
		if err != nil {
			fmt.Println("CommodityService.Delete(), redis err: ", err)
		}
	}()

	return nil
}

// status change remove cache from redis
func (s *CommodityService) Undelete(ctx *gin.Context, id string) error {
	if strings.TrimSpace(id) == "" {
		return e.VALIDATE_ERR
	}

	err := s.svcctx.CommodityModel.Conn(s.svcctx.DB).Undelete(id)
	if err != nil {
		fmt.Println("CommodityService.Undelete(), database err: ", err)
		return err
	}

	// use another routine to remove cache
	go func() {
		err = connection.NewRedisConnection[string]().DoHashRemove(constant.REDIS_COMMODITY_INFO_HASH, constant.RedisCommodityHashKey(id))
		if err != nil {
			fmt.Println("CommodityService.Undelete(), redis err: ", err)
		}
	}()

	return nil
}

func (s *CommodityService) FindDetailViewById(ctx *gin.Context, id string) (*dto.CommodityDetailView, error) {
	commodity := &model.Commodity{}
	var err error

	// read from redis
	err = connection.NewRedisConnection[*model.Commodity]().DoHashGet(constant.REDIS_COMMODITY_INFO_HASH, constant.RedisCommodityHashKey(id), commodity)
	if err != nil {
		fmt.Println("CommodityService.FindDetailById(), redis err: ", err)
		return nil, err
	}

	// get from db
	if reflect.DeepEqual(*commodity, model.Commodity{}) {
		commodity, err = s.svcctx.CommodityModel.Conn(s.svcctx.DB).FindById(id)
		if err != nil {
			fmt.Println("CommodityService.FindDetailById(), database err: ", err)
			return nil, err
		}

		// commodity not exists
		if reflect.DeepEqual(*commodity, model.Commodity{}) {
			return nil, nil
		}

		// cache to redis use another routine
		go func() {
			err = connection.NewRedisConnection[*model.Commodity]().DoHashSet(constant.REDIS_COMMODITY_INFO_HASH, constant.RedisCommodityHashKey(id), commodity, 0)
			if err != nil {
				fmt.Println("CommodityService.FindDetailById(), redis err: ", err)
			}
		}()
	}

	return s.CastToDetailView(commodity), nil
}

func (s *CommodityService) FindDetailViewByIsbn(ctx *gin.Context, isbn string) (*dto.CommodityDetailView, error) {
	commodity := &model.Commodity{}
	var err error

	// read from redis
	commoditySlice := make([]model.Commodity, 0)
	err = connection.NewRedisConnection[model.Commodity]().DoHashGetAll(constant.REDIS_COMMODITY_INFO_HASH, *commodity, commoditySlice)
	if err != nil {
		fmt.Println("CommodityService.FindDetailByIsbn(), redis err: ", err)
		return nil, err
	}
	for _, c := range commoditySlice {
		if c.Extension["ISBN"] == isbn {
			commodity = &c
			break
		}
	}

	// get from db
	if reflect.DeepEqual(*commodity, model.Commodity{}) {
		commodity, err = s.svcctx.CommodityModel.Conn(s.svcctx.DB).FindByIsbn(isbn)
		if err != nil {
			fmt.Println("CommodityService.FindDetailByIsbn(), database err: ", err)
			return nil, err
		}

		// commodity not exists
		if reflect.DeepEqual(*commodity, model.Commodity{}) {
			return nil, nil
		}

		// cache to redis use another routine
		go func() {
			err = connection.NewRedisConnection[*model.Commodity]().DoHashSet(constant.REDIS_COMMODITY_INFO_HASH, constant.RedisCommodityHashKey(commodity.Id), commodity, 0)
			if err != nil {
				fmt.Println("CommodityService.FindDetailByIsbn(), redis err: ", err)
			}
		}()
	}

	return s.CastToDetailView(commodity), nil
}

// TODO add redis cache
// use elasticsearch to store data, sort through it and get ids of target commodities
// then use the id to get from redis
func (s *CommodityService) FindSimpleViewPage(ctx *gin.Context, page *util.Page[*dto.CommoditySimpleView]) error {
	// sort through es

	// get cache from redis

	// if redis not cached, get from db
	commoditySlice, err := s.svcctx.CommodityModel.Conn(s.svcctx.DB).FindPage(page.Num, page.Size)
	if err != nil {
		fmt.Println("CommodityService.FindSimpleViewPage(), database err: ", err)
		return err
	}

	simpleViewSlice := s.CastToSimpleViewSlice(commoditySlice)
	page.Data = simpleViewSlice

	return nil
}

func (s *CommodityService) FindSimpleViewPageByType(ctx *gin.Context, t string, page *util.Page[*dto.CommoditySimpleView]) error {
	commoditySlice, err := s.svcctx.CommodityModel.Conn(s.svcctx.DB).FindPageByType(t, page.Num, page.Size)
	if err != nil {
		fmt.Println("CommodityService.FindSimpleViewPageByType(), database err: ", err)
		return err
	}

	simpleViewSlice := s.CastToSimpleViewSlice(commoditySlice)
	page.Data = simpleViewSlice

	return nil
}

func (s *CommodityService) FindSimpleViewPageByFuzzyName(ctx *gin.Context, name string, page *util.Page[*dto.CommoditySimpleView]) error {
	commoditySlice, err := s.svcctx.CommodityModel.Conn(s.svcctx.DB).FindPageByFuzzyName(name, page.Num, page.Size)
	if err != nil {
		fmt.Println("CommodityService.FindSimpleViewPageByFuzzyName(), database err: ", err)
		return err
	}

	simpleViewSlice := s.CastToSimpleViewSlice(commoditySlice)
	page.Data = simpleViewSlice

	return nil
}

func (s *CommodityService) FindSimpleViewPageByFuzzyNameAndType(ctx *gin.Context, name string, t string, page *util.Page[*dto.CommoditySimpleView]) error {
	commoditySlice, err := s.svcctx.CommodityModel.Conn(s.svcctx.DB).FindPageByFuzzyNameAndType(name, t, page.Num, page.Size)
	if err != nil {
		fmt.Println("CommodityService.FindSimpleViewPageByFuzzyNameAndType(), database err: ", err)
		return err
	}

	simpleViewSlice := s.CastToSimpleViewSlice(commoditySlice)
	page.Data = simpleViewSlice

	return nil
}
