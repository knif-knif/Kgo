package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/knif-knif/Kgo/app/frontend/hertz_gen/frontend/common"
	"github.com/knif-knif/Kgo/app/frontend/infra/rpc"
	"github.com/knif-knif/Kgo/rpc_gen/kitex_gen/product"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (res map[string]any, err error) {
	products, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{})
	if err != nil {
		return nil, err
	}
	var cartNum int
	return map[string]any{
		"Title":    "Hot Sales",
		"cart_num": cartNum,
		"Items":    products.Products,
	}, nil
}
