package rpc

import (
	"sync"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/conf"
	FrontendUtils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient userservice.Client
	once       sync.Once
)

func Init() {
	once.Do(func() {
		initUserClient()
	})
}

func initUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegisterAddr)
	FrontendUtils.MustHandleError(err)

	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	FrontendUtils.MustHandleError(err)
}
