package service

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/knif-knif/Kgo/app/user/biz/dal/mysql"
	user "github.com/knif-knif/Kgo/rpc_gen/kitex_gen/user"
	"testing"
)

func TestRegister_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:     "123@13.com",
		Password:  "123456",
		PasswordC: "123456",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
