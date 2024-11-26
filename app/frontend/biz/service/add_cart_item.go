package service

import (
	"context"
	"github.com/knif-knif/Kgo/app/frontend/infra/rpc"
	frontendUtils "github.com/knif-knif/Kgo/app/frontend/utils"

	"github.com/cloudwego/hertz/pkg/app"
	cart "github.com/knif-knif/Kgo/app/frontend/hertz_gen/frontend/cart"
	common "github.com/knif-knif/Kgo/app/frontend/hertz_gen/frontend/common"
	rpccart "github.com/knif-knif/Kgo/rpc_gen/kitex_gen/cart"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartItemReq) (resp *common.Empty, err error) {
	_, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		UserId: frontendUtils.GetUserIdFromCtx(h.Context),
		Item: &rpccart.CartItem{
			ProductId: req.ProductId,
			Quantity:  uint32(req.ProductNum),
		},
	})
	if err != nil {
		return nil, err
	}
	return
}