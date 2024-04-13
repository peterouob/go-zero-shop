package logic

import (
	"context"
	"errors"
	"go-mirco-zero/user/model"
	"google.golang.org/grpc/status"

	"go-mirco-zero/product/rpc/internal/svc"
	"go-mirco-zero/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLogic) Delete(in *product.DeleteRequest) (*product.DeleteResponse, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.Product.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(100, "產品不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	err = l.svcCtx.Product.Delete(l.ctx, res.Id)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &product.DeleteResponse{}, nil
}
