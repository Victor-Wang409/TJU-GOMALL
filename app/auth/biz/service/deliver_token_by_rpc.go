package service

import (
	"context"
	"time"

	auth "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/auth"
	"github.com/golang-jwt/jwt/v4"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// Finish your business logic.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserId":  req.UserId,
		"Timeout": time.Hour,
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}
	// fmt.Println("token:", token)
	deliveryResp := &auth.DeliveryResp{
		Token: tokenString,
	}

	return deliveryResp, nil
}
