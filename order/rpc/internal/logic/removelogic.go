package logic

import (
	"context"

	"go-mirco-zero/order/rpc/internal/svc"
	"go-mirco-zero/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveLogic) Remove(in *order.RemoveRequest) (*order.RemoveResponse, error) {
	// todo: add your logic here and delete this line

	return &order.RemoveResponse{}, nil
}
