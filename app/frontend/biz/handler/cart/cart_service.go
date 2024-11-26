package cart

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	hertzUtils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/knif-knif/Kgo/app/frontend/biz/service"
	"github.com/knif-knif/Kgo/app/frontend/biz/utils"
	cart "github.com/knif-knif/Kgo/app/frontend/hertz_gen/frontend/cart"
	common "github.com/knif-knif/Kgo/app/frontend/hertz_gen/frontend/common"
)

// GetCart .
// @router /cart [GET]
func GetCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		c.HTML(consts.StatusOK, "cart", utils.WarpResponse(ctx, c, hertzUtils.H{"warning": err}))
		return
	}
	resp, err := service.NewGetCartService(ctx, c).Run(&req)
	if err != nil {
		fmt.Println(err)
		c.HTML(consts.StatusOK, "cart", utils.WarpResponse(ctx, c, hertzUtils.H{"error": err}))
		return
	}
	c.HTML(consts.StatusOK, "cart", utils.WarpResponse(ctx, c, resp))
}

// AddCartItem .
// @router /cart [POST]
func AddCartItem(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cart.AddCartItemReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewAddCartItemService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.Redirect(consts.StatusFound, []byte("/cart"))
	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
