package service

import (
	"context"
	"testing"

	"github.com/cloudwego/biz-demo/gomall/app/cart/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/cart/infra/rpc"
	cart "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/joho/godotenv"
)

func TestAddItem_Run(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		klog.Error(err.Error())
	}
	mysql.Init()
	rpc.Init()
	ctx := context.Background()
	s := NewAddItemService(ctx)

	req := &cart.AddItemReq{
		UserId: uint32(1),
		Item: &cart.CartItem{
			ProductId: uint32(2),
			Quantity:  int32(10),
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
