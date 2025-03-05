package service

import (
	"context"
	"testing"

	auth "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/auth"
)

func TestVerifyTokenByRPC_Run(t *testing.T) {
	ctx := context.Background()
	s := NewVerifyTokenByRPCService(ctx)

	req := &auth.VerifyTokenReq{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJUaW1lb3V0IjozNjAwMDAwMDAwMDAwLCJVc2VySWQiOjF9.FTplAyK1JKCCdmaMJLLNQlTq-rMciGb_r280uBzwFlI",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

}
