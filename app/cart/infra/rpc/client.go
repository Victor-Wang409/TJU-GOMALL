package rpc

import (
	"sync"

	"github.com/cloudwego/biz-demo/gomall/app/cart/conf"
	CartUtils "github.com/cloudwego/biz-demo/gomall/app/cart/utils"
	"github.com/cloudwego/biz-demo/gomall/common/clientsuite"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
	ServiceName   = conf.GetConf().Kitex.Service
	RegistryAddr  = conf.GetConf().Registry.RegistryAddress[0]
	err           error
)

func Init() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	// var opts []client.Option
	// r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	// CartUtils.MustHandleError(err)
	// opts = append(opts, client.WithResolver(r))

	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	CartUtils.MustHandleError(err)
}
