package service

import (
	"context"
	"testing"

	auth "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/auth"
)

func TestDeliverTokenByRPC_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeliverTokenByRPCService(ctx)

	req := &auth.DeliverTokenReq{
		UserId: 1,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
