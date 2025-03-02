package rpc

import (
	"sync"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/conf"
	FrontendUtils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	once           sync.Once
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
	})
}

func initUserClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegisterAddr)
	FrontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	UserClient, err = userservice.NewClient("user", opts...)
	FrontendUtils.MustHandleError(err)
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegisterAddr)
	FrontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	FrontendUtils.MustHandleError(err)
}

func initCartClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegisterAddr)
	FrontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	CartClient, err = cartservice.NewClient("cart", opts...)
	FrontendUtils.MustHandleError(err)
}

func initCheckoutClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegisterAddr)
	FrontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	CheckoutClient, err = checkoutservice.NewClient("checkout", opts...)
	FrontendUtils.MustHandleError(err)
}

func initOrderClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegisterAddr)
	FrontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	OrderClient, err = orderservice.NewClient("order", opts...)
	FrontendUtils.MustHandleError(err)
}
