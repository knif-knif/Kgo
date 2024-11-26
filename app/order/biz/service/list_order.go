package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/knif-knif/Kgo/app/order/biz/dal/mysql"
	"github.com/knif-knif/Kgo/app/order/biz/model"
	"github.com/knif-knif/Kgo/rpc_gen/kitex_gen/cart"
	order "github.com/knif-knif/Kgo/rpc_gen/kitex_gen/order"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	list, err := model.ListOrder(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(500001, err.Error())
	}
	var orders []*order.Order
	for _, v := range list {
		var items []*order.OrderItem
		for _, oi := range v.OrderItems {
			items = append(items, &order.OrderItem{
				Item: &cart.CartItem{
					ProductId: oi.ProductId,
					Quantity:  oi.Quantity,
				},
				Cost: oi.Cost,
			})
		}
		orders = append(orders, &order.Order{
			OrderId:      v.OrderId,
			UserId:       v.UserId,
			UserCurrency: v.UserCurrency,
			Email:        v.Consignee.Email,
			Address: &order.Address{
				StreetAddress: v.Consignee.StreetAddress,
				State:         v.Consignee.State,
				City:          v.Consignee.City,
				Country:       v.Consignee.Country,
				ZipCode:       int32(v.Consignee.ZipCode),
			},
			Items: items,
		})
	}
	resp = &order.ListOrderResp{
		Orders: orders,
	}
	return
}
