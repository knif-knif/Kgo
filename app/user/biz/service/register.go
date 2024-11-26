package service

import (
	"context"
	"errors"
	"github.com/knif-knif/Kgo/app/user/biz/dal/mysql"
	"github.com/knif-knif/Kgo/app/user/biz/model"
	user "github.com/knif-knif/Kgo/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	if req.Email == "" || req.Password == "" || req.PasswordC == "" {
		return nil, errors.New("email or password is empty")
	}

	if req.Password != req.PasswordC {
		return nil, errors.New("password not match")
	}
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(passwordHashed),
	}
	err = model.Create(mysql.DB, newUser)
	if err != nil {
		return nil, err
	}
	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}
