package service

import (
	"context"
	"github.com/hertz-contrib/sessions"
	"github.com/knif-knif/Kgo/app/frontend/infra/rpc"
	"github.com/knif-knif/Kgo/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	auth "github.com/knif-knif/Kgo/app/frontend/hertz_gen/frontend/auth"
	common "github.com/knif-knif/Kgo/app/frontend/hertz_gen/frontend/common"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userResp, err := rpc.UserClient.Register(h.Context, &user.RegisterReq{
		Email:     req.Email,
		Password:  req.Password,
		PasswordC: req.PasswordC,
	})
	if err != nil {
		return nil, err
	}
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", userResp.UserId)
	err = session.Save()
	if err != nil {
		return nil, err
	}
	return
}
