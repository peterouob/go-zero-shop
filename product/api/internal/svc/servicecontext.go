package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-mirco-zero/product/api/internal/config"
	"go-mirco-zero/product/rpc/productclient"
)

type ServiceContext struct {
	Config     config.Config
	ProductRpc productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
