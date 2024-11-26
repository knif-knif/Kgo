package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/knif-knif/Kgo/app/frontend/infra/rpc"
	frontendUtils "github.com/knif-knif/Kgo/app/frontend/utils"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/knif-knif/Kgo/app/frontend/hertz_gen/frontend/common"
	rpccart "github.com/knif-knif/Kgo/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/knif-knif/Kgo/rpc_gen/kitex_gen/product"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	var items []map[string]string
	cartResp, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{
		UserId: frontendUtils.GetUserIdFromCtx(h.Context),
	})

	if err != nil {
		return nil, err
	}

	var total float32
	for _, item := range cartResp.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{
			Id: item.GetProductId(),
		})
		if err != nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]string{
			"Name":    p.Name,
			"Desc":    p.Description,
			"Price":   strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Picture": p.Picture,
			"Qty":     strconv.Itoa(int(item.Quantity)),
		})
		total += float32(item.Quantity) * p.Price
	}
	return utils.H{
		"Title": "Cart",
		"Items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
