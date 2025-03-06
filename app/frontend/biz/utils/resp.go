package utils

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	FrontendUtils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	userId := FrontendUtils.GetUserIdFromCtx(ctx)
	token := FrontendUtils.GetTokenFromCtx(ctx)
	content["user_id"] = userId
	content["token"] = token
	// fmt.Println("user_id", content["user_id"])
	// fmt.Println("token:", content["token"])
	if userId > 0 {
		cartResp, err := rpc.CartClient.GetCart(ctx, &cart.GetCartReq{UserId: uint32(userId)})

		if err == nil && cartResp != nil {
			content["cart_num"] = len(cartResp.Cart.Items)
			// content["user_id"] = ctx.Value(FrontendUtils.UserIdKey)
		}
	}
	return content
}
