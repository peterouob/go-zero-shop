package logic

import (
	"context"
	"go-mirco-zero/product/rpc/types/product"

	"go-mirco-zero/product/api/internal/svc"
	"go-mirco-zero/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.DeleteRequest) (resp *types.DeleteResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.ProductRpc.Delete(l.ctx, &product.DeleteRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.DeleteResponse{}, nil
}
