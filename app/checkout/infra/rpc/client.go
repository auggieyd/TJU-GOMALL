package rpc

import (
	"sync"

	"github.com/cloudwego/biz-demo/gomall/app/checkout/conf"
	CartUtils "github.com/cloudwego/biz-demo/gomall/app/checkout/utils"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client
	once          sync.Once
)

func InitClient() {
	once.Do(func() {

		initCartClient()
		initProductClient()
		initPaymentClient()
		// initOrderClient()
	})
}

func initCartClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	CartUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
	)

	CartClient, err = cartservice.NewClient("cart", opts...)
	CartUtils.MustHandleError(err)
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	CartUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
	)

	CartClient, err = cartservice.NewClient("product", opts...)
	CartUtils.MustHandleError(err)
}

func initPaymentClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	CartUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
	)

	CartClient, err = cartservice.NewClient("payment", opts...)
	CartUtils.MustHandleError(err)
}

// func initOrderClient() {
// 	OrderClient, err = orderservice.NewClient("order", commonSuite)
// 	checkoututils.MustHandleError(err)
// }
