package service

import (
	"context"
	"testing"

	"github.com/cloudwego/biz-demo/gomall/app/checkout/infra/rpc"
	checkout "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/payment"
)

func TestCheckout_Run(t *testing.T) {

	rpc.InitClient()
	ctx := context.Background()
	s := NewCheckoutService(ctx)

	req := &checkout.CheckoutReq{
		UserId:    uint32(1),
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "123@qq.com",
		Address: &checkout.Address{
			StreetAddress: "123 Main St",
			City:          "San Francisco",
			Country:       "USA",
			State:         "CA",
			ZipCode:       "94101",
		},
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          "1234567890123456",
			CreditCardCvv:             int32(123),
			CreditCardExpirationMonth: int32(1),
			CreditCardExpirationYear:  int32(2026),
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
