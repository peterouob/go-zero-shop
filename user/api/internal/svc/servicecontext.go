package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-mirco-zero/user/api/internal/config"
	"go-mirco-zero/user/rpc/userclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
