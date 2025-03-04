package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudwego/biz-demo/gomall/app/user/biz/dal/mysql"
	user "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/joho/godotenv"
)

func TestRegister_Run(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("env error!")
		klog.Error(err.Error())
	}
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)

	req := &user.RegisterReq{
		Email:           "123@qq.com",
		Password:        "123",
		PasswordConfirm: "123",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
