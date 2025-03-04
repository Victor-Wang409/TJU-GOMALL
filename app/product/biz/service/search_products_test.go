package service

import (
	"context"
	"testing"

	"github.com/cloudwego/biz-demo/gomall/app/product/biz/dal/mysql"
	product "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/joho/godotenv"
)

func TestSearchProducts_Run(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		klog.Error(err.Error())
	}
	mysql.Init()
	ctx := context.Background()
	s := NewSearchProductsService(ctx)

	req := &product.SearchProductsReq{Query: "shirt"}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
