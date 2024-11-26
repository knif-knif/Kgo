package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/knif-knif/Kgo/app/frontend/infra/rpc"
	frontendUtils "github.com/knif-knif/Kgo/app/frontend/utils"
	rpcpayment "github.com/knif-knif/Kgo/rpc_gen/kitex_gen/payment"

	"github.com/cloudwego/hertz/pkg/app"
	checkout "github.com/knif-knif/Kgo/app/frontend/hertz_gen/frontend/checkout"
	rpccheckout "github.com/knif-knif/Kgo/rpc_gen/kitex_gen/checkout"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	_, err = rpc.CheckoutClient.Checkout(h.Context, &rpccheckout.CheckoutReq{
		UserId:    userId,
		Email:     req.Email,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Address: &rpccheckout.Address{
			City:          req.City,
			Country:       req.Country,
			ZipCode:       req.Zipcode,
			State:         req.Province,
			StreetAddress: req.Street,
		},
		CreditCard: &rpcpayment.CreditCardInfo{
			CreditCardNumber:          req.CardNum,
			CreditCardExpirationYear:  req.ExpirationYear,
			CreditCardCvv:             req.Cvv,
			CreditCardExpirationMonth: req.ExpirationMonth,
		},
	})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"Title":    "waiting",
		"redirect": "/checkout/result",
	}, nil
}
