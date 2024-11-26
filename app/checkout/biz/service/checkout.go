package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/golang/protobuf/proto"
	"github.com/knif-knif/Kgo/app/checkout/infra/mq"
	"github.com/knif-knif/Kgo/app/checkout/infra/rpc"
	"github.com/knif-knif/Kgo/rpc_gen/kitex_gen/cart"
	checkout "github.com/knif-knif/Kgo/rpc_gen/kitex_gen/checkout"
	"github.com/knif-knif/Kgo/rpc_gen/kitex_gen/email"
	"github.com/knif-knif/Kgo/rpc_gen/kitex_gen/order"
	"github.com/knif-knif/Kgo/rpc_gen/kitex_gen/payment"
	"github.com/knif-knif/Kgo/rpc_gen/kitex_gen/product"
	"github.com/nats-io/nats.go"
	"strconv"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	cardResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}
	if cardResult == nil || cardResult.Items == nil {
		return nil, kerrors.NewGRPCBizStatusError(5004001, "cart is empty")
	}

	var (
		total float32
		oi    []*order.OrderItem
	)
	for _, cartItem := range cardResult.Items {
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
			Id: cartItem.ProductId,
		})
		if resultErr != nil {
			return nil, resultErr
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product.Price
		cost := p * float32(cartItem.Quantity)
		total += cost
		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{
				ProductId: cartItem.ProductId,
				Quantity:  cartItem.Quantity,
			},
			Cost: cost,
		})
	}
	var orderId string
	zipCodeInt, _ := strconv.Atoi(req.Address.ZipCode)
	orderResp, err := rpc.OrderClient.PlaceOrder(s.ctx, &order.PlaceOrderReq{
		UserId: req.UserId,
		Email:  req.Email,
		Address: &order.Address{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       int32(zipCodeInt),
		},
		Items: oi,
	})
	if err != nil {
		return nil, kerrors.NewBizStatusError(5004002, err.Error())
	}
	if orderResp.Order != nil || orderResp.Order != nil {
		orderId = orderResp.Order.OrderId
	}
	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
		},
	}
	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{
		UserId: req.UserId,
	})
	if err != nil {
		klog.Error(err.Error())
	}
	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, err
	}

	data, _ := proto.Marshal(&email.EmailReq{
		From:        "from@example.com",
		To:          req.Email,
		ContentType: "You have just created an order in the KGo",
		Content:     "Your order has been placed successfully",
	})
	msg := &nats.Msg{Subject: "email", Data: data}
	_ = mq.Nc.PublishMsg(msg)
	klog.Info(paymentResult)
	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}
	return
}
