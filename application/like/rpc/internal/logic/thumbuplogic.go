package logic

import (
	"context"
	"encoding/json"

	"zeroblog/application/like/rpc/internal/svc"
	"zeroblog/application/like/rpc/internal/types"
	"zeroblog/application/like/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
)

type ThumbupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewThumbupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbupLogic {
	return &ThumbupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ThumbupLogic) Thumbup(in *service.ThumbupRequest) (*service.ThumbupResponse, error) {
	// TODO:
	// 1. 查询是否已经点赞
	// 2. 计算当前内容的总点赞数和点踩数

	msg := &types.ThumbupMsg{
		BizId:    in.BizId,
		LikeType: in.LikeType,
		ObjId:    in.ObjId,
		UserId:   in.UserId,
	}

	// 异步，不用 go func
	threading.GoSafe(func() {
		data, err := json.Marshal(msg)
		if err != nil {
			l.Logger.Errorf("[Thumbup] marshal msg %v error: %v", msg, err)
			return
		}

		err = l.svcCtx.KqPusherClient.Push(l.ctx, string(data))
		if err != nil {
			l.Logger.Errorf("[Thumbup] push data %s error: %v", data, err)
		}
	})

	return &service.ThumbupResponse{}, nil
}
