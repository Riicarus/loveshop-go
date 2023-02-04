package service

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/riicarus/loveshop/internal/constant"
	"github.com/riicarus/loveshop/internal/context"
	"github.com/riicarus/loveshop/internal/entity/dto"
	"github.com/riicarus/loveshop/internal/model"
	"github.com/riicarus/loveshop/pkg/connection"
	"github.com/riicarus/loveshop/pkg/e"
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

func (s *CommodityService) Add(ctx *gin.Context, param *dto.CommodityAddParam) error {
	// TODO check param
	err := s.svcctx.CommodityModel.Add(param)
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

	err := s.svcctx.CommodityModel.Update(param)
	if err != nil {
		fmt.Println("CommodityService.Update(), database err: ", err)
		return err
	}

	// use another routine to remove cache
	go func() {
		err = connection.RedisConn.DoHashRemove(constant.REDIS_COMMODITY_HASH, constant.RedisCommodityHashKey(param.Id))
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

	err := s.svcctx.CommodityModel.UpdateAmount(id, number)
	if err != nil {
		return err
	}

	// use another routine to remove cache
	go func() {
		err = connection.RedisConn.DoHashRemove(constant.REDIS_COMMODITY_HASH, constant.RedisCommodityHashKey(id))
		if err != nil {
			fmt.Println("CommodityService.UpdateAmount(), redis err: ", err)
		}
	}()

	return err
}

// status change remove cache from redis
func (s *CommodityService) Delete(ctx *gin.Context, id string) error {
	if strings.TrimSpace(id) == "" {
		return e.VALIDATE_ERR
	}

	err := s.svcctx.CommodityModel.Delete(id)
	if err != nil {
		return err
	}

	// use another routine to remove cache
	go func() {
		err = connection.RedisConn.DoHashRemove(constant.REDIS_COMMODITY_HASH, constant.RedisCommodityHashKey(id))
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

	err := s.svcctx.CommodityModel.Undelete(id)
	if err != nil {
		return err
	}

	// use another routine to remove cache
	go func() {
		err = connection.RedisConn.DoHashRemove(constant.REDIS_COMMODITY_HASH, constant.RedisCommodityHashKey(id))
		if err != nil {
			fmt.Println("CommodityService.Undelete(), redis err: ", err)
		}
	}()

	return nil
}

// TODO add redis cache
// use elasticsearch to store data, sort through it and get ids of target commodities
// then use the id to get from redis
func (s *CommodityService) FindInfoPage(ctx *gin.Context, page *util.Page[*dto.CommoditySimpleInfo]) error {
	// sort through es

	// get cache from redis

	// if redis not cached, get from db
	commoditySlice, err := s.svcctx.CommodityModel.FindInfoPage(page.Num, page.Size)
	if err != nil {
		fmt.Println("CommodityService.FindInfoPage(), database err: ", err)
		return err
	}

	infoSlice := make([]*dto.CommoditySimpleInfo, 0, len(commoditySlice))
	for _, c := range commoditySlice {
		infoSlice = append(infoSlice, &dto.CommoditySimpleInfo{
			Id:        c.Id,
			Type:      c.Type,
			Numbering: c.Numbering,
			Name:      c.Name,
			Amount:    c.Amount,
			Price:     c.Price,
			Extension: c.Extension,
		})
	}

	page.Data = infoSlice

	return nil
}

func (s *CommodityService) FindDetailById(ctx *gin.Context, id string) (*dto.CommodityDetailInfo, error) {
	commodity := &model.Commodity{}
	var err error

	// read from redis
	err = connection.RedisConn.DoHashGet(constant.REDIS_COMMODITY_HASH, constant.RedisCommodityHashKey(id), commodity)
	if err != nil {
		fmt.Println("CommodityService.FindDetailById(), redis err: ", err)
		return nil, err
	}

	// get from db
	if reflect.DeepEqual(*commodity, model.Commodity{}) {
		commodity, err = s.svcctx.CommodityModel.FindById(id)
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
			err = connection.RedisConn.DoHashSet(constant.REDIS_COMMODITY_HASH, constant.RedisCommodityHashKey(id), commodity, 0)
			if err != nil {
				fmt.Println("CommodityService.FindDetailById(), redis err: ", err)
			}
		}()
	}

	return &dto.CommodityDetailInfo{
		Id:        commodity.Id,
		Type:      commodity.Type,
		Numbering: commodity.Numbering,
		Name:      commodity.Name,
		Amount:    commodity.Amount,
		Price:     commodity.Price,
		Extension: commodity.Extension,
		Deleted:   commodity.Deleted,
	}, nil
}
