package service

import (
	"fmt"
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

type AdminService struct {
	svcctx *context.ServiceContext
}

func NewAdminService(svcctx *context.ServiceContext) *AdminService {
	return &AdminService{
		svcctx: svcctx,
	}
}

func (s *AdminService) LoginWithPass(ctx *gin.Context, loginParam *dto.AdminLoginParam) (string, error) {
	admin, err1 := s.svcctx.AdminModel.FindByStudentId(loginParam.StudentId)
	if err1 != nil {
		fmt.Println("AdminService.LoginWithPass(), db err: ", err1)
		return "", err1
	}

	// no such admin
	if admin == nil {
		return "", e.UNAUTHED_ERR
	}

	// not enabled admin
	if !admin.Enabled {
		return "", e.UNAUTHED_ERR
	}

	md5Pass := util.Md5(loginParam.Password, admin.Salt, 1024)
	if md5Pass == admin.Password {
		token, err2 := util.GenToken(admin.StudentId, constant.ADMIN_LOGIN_TYPE)
		if err2 != nil {
			fmt.Println("AdminService.LoginWithPass(), encoding jwt err: ", err2)
			return "", err2
		}

		return token, nil
	}

	// password not match
	return "", e.UNAUTHED_ERR
}

func (s *AdminService) Unable(ctx *gin.Context, id string) error {
	err := s.svcctx.AdminModel.Unable(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *AdminService) Register(ctx *gin.Context, registerParam *dto.AdminRegisterParam) error {
	salt := fmt.Sprintf("%d", time.Now().Unix())

	roleIds := make(util.JSONStringSlice, 0)
	//TODO get role ids from role table by role name

	admin := &model.Admin{
		Id:          uuid.New().String(),
		Name:        registerParam.Name,
		StudentId:   registerParam.StudentId,
		Password:    util.Md5(registerParam.Password, salt, 1024),
		Email:       registerParam.Email,
		Salt:        salt,
		Group:       registerParam.Group,
		Integration: 0,
		RoleIds:     roleIds,
		Enabled:     true,
	}

	err := s.svcctx.AdminModel.Register(admin)
	if err != nil {
		fmt.Println("AdminService.Register() err: ", err)
		return err
	}

	return nil
}
