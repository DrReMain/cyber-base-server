package auth_account_service

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/gorm"

	"github.com/DrReMain/cyber-base-server/biz/common/errc"
	"github.com/DrReMain/cyber-base-server/biz/dal/sys_model"
	"github.com/DrReMain/cyber-base-server/biz/hertz_gen/auth/account"
	cutils_crypto "github.com/DrReMain/cyber-base-server/cyber/utils/crypto"
	cutils_jwt "github.com/DrReMain/cyber-base-server/cyber/utils/jwt"
)

type Service struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewService(ctx context.Context, c *app.RequestContext) *Service {
	return &Service{ctx, c}
}

func (s *Service) Register(req *account.RegisterAccountReq) (err error) {
	if req.Password != req.Confirm {
		err = errc.ParamsInvalidErr.WithMsg("两次密码不一致")
		return
	}

	item, _ := sys_model.QueryByEmail(req.Email)
	if *item != (sys_model.SysUser{}) {
		err = errc.AlreadyExistErr.WithMsg("用户已存在")
		return
	}

	password, err := cutils_crypto.Crypt(req.Password)
	m := &sys_model.SysUser{
		Email:    &req.Email,
		Password: &password,
	}
	err = sys_model.CreateUser(m)
	return
}

func (s *Service) Login(req *account.LoginAccountReq) (accessToken string, err error) {
	item, err := sys_model.QueryByEmail(req.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errc.AuthorizationFailedErr.WithMsg("Email或密码错误")
		return
	}

	if ok := cutils_crypto.Verify(req.Password, *item.Password); !ok {
		err = errc.AuthorizationFailedErr.WithMsg("Email或密码错误")
		return
	}

	if item.Ban == 1 {
		err = errc.AuthorizationFailedErr.WithMsg("当前用户被禁止登录")
		return
	}
	accessToken, err = s.generate(item)
	return
}

func (s *Service) generate(u *sys_model.SysUser) (token string, err error) {
	j := cutils_jwt.NewJsonWebToken()
	token, err = j.GenToken(j.CreateClaims(cutils_jwt.BaseClaims{
		UserID: u.ID,
		Email:  *u.Email,
	}))
	if err != nil {
		err = errc.InternalErr.WithMsg("生成token失败")
		return
	}

	return
}
