package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-mirco-zero/product/model"
	"go-mirco-zero/product/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	Product model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:  c,
		Product: model.NewProductModel(conn, c.CacheRedis),
	}
}
