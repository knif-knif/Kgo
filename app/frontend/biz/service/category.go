package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/knif-knif/Kgo/app/frontend/infra/rpc"
	"github.com/knif-knif/Kgo/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	category "github.com/knif-knif/Kgo/app/frontend/hertz_gen/frontend/category"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {
	p, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{CategoryName: req.Category})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"title": "Category",
		"items": p.Products,
	}, nil
}
