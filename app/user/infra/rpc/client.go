package rpc

import (
	"sync"

	"github.com/cloudwego/biz-demo/gomall/app/user/conf"
	UserUtils "github.com/cloudwego/biz-demo/gomall/app/user/utils"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/auth/authservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	AuthClient authservice.Client
	once       sync.Once
)

func Init() {
	once.Do(func() {
		initAuthClient()
	})
}

func initAuthClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	UserUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	AuthClient, err = authservice.NewClient("auth", opts...)
	UserUtils.MustHandleError(err)
}
