package service

import (
	"context"
	"fmt"

	auth "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/auth"
	"github.com/golang-jwt/jwt/v4"
)

type VerifyTokenByRPCService struct {
	ctx context.Context
} // NewVerifyTokenByRPCService new VerifyTokenByRPCService
func NewVerifyTokenByRPCService(ctx context.Context) *VerifyTokenByRPCService {
	return &VerifyTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {

	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}

	// fmt.Println("header:", token.Header)
	// fmt.Println("claims:", token.Claims)
	// fmt.Println("signature:", base64.RawURLEncoding.EncodeToString([]byte(token.Signature)))
	verifyResp := &auth.VerifyResp{}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		verifyResp.Res = true
		fmt.Println(claims["UserId"], claims["Timeout"])
	} else {
		verifyResp.Res = false
	}

	return verifyResp, nil
}
