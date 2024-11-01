package logic

import (
	"context"
	"encoding/json"

	"zeroblog/application/article/api/internal/code"
	"zeroblog/application/article/api/internal/svc"
	"zeroblog/application/article/api/internal/types"
	"zeroblog/application/article/rpc/article"
	"zeroblog/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

const minContentLen = 80

type PublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishLogic) Publish(req *types.PublishRequest) (resp *types.PublishResponse, err error) {
	if len(req.Title) == 0 {
		return nil, code.ArtitleTitleEmpty
	}
	if len(req.Content) < minContentLen {
		return nil, code.ArticleContentTooFewWords
	}
	if len(req.Cover) == 0 {
		return nil, code.ArticleCoverEmpty
	}

	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		logx.Errorf("l.ctx.Value error: %v", err)
		return nil, xcode.NoLogin
	}

	pret, err := l.svcCtx.ArticleRPC.Publish(l.ctx, &article.PublishRequest{
		UserId:      userId,
		Title:       req.Title,
		Content:     req.Content,
		Cover:       req.Cover,
		Description: req.Description,
	})
	if err != nil {
		logx.Errorf("l.svcCtx.ArticleRPC.Publish req: %v userId: %d error: %v", req, userId, err)
		return nil, err
	}

	return &types.PublishResponse{ArticleId: pret.ArticleId}, nil
}
