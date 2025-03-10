package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cloudwego/biz-demo/gomall/app/checkout/infra/mq"
	"github.com/cloudwego/biz-demo/gomall/app/checkout/infra/rpc"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	checkout "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/email"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {

	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		// fmt.Println("Empty Cart 1!")
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}
	if cartResult == nil || cartResult.Cart == nil || len(cartResult.Cart.Items) == 0 {
		// fmt.Println("Empty Cart 2!")
		return nil, kerrors.NewGRPCBizStatusError(5004001, "Cart is Empty!")
	}

	var (
		oi    []*order.OrderItem
		total float32
	)
	for _, cartItem := range cartResult.Cart.Items {
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: cartItem.ProductId})
		if resultErr != nil {
			return nil, resultErr
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		cost := p.Price * float32(cartItem.Quantity)
		total += cost
		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{
				ProductId: cartItem.ProductId,
				Quantity:  cartItem.Quantity},
			Cost: cost,
		})
	}

	var orderId string
	// u, _ := uuid.NewRandom()
	// orderId = u.String()
	zipcode, _ := strconv.Atoi(req.Address.ZipCode)
	orderResp, err := rpc.OrderClient.PlaceOrder(s.ctx, &order.PlaceOrderReq{
		UserId:     req.UserId,
		OrderItems: oi,
		Email:      req.Email,
		Address: &order.Address{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			Country:       req.Address.Country,
			State:         req.Address.State,
			ZipCode:       int32(zipcode),
		},
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5004002, err.Error())
	}
	if orderResp != nil && orderResp.Order != nil {
		orderId = orderResp.Order.OrderId
	}

	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
		},
	}

	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		return nil, err
	}

	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		fmt.Println("Payment ERROR!", payReq)
		return nil, err
	}

	data, _ := proto.Marshal(&email.EmailReq{
		From:        "example@qq.com",
		To:          req.Email,
		ContentType: "text/plain",
		Subject:     "Order Confirmation",
		Content:     "Order Confirmation",
		// Content:     fmt.Sprintf("Your order %s has been placed successfully!", orderId),
	})
	msg := nats.Msg{
		Subject: "email",
		Data:    data,
	}
	_ = mq.Nc.PublishMsg(&msg)

	klog.Info(paymentResult)

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}

	return
}
