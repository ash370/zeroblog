package logic

import (
	"context"
	"errors"
	"time"

	"zeroblog/application/user/rpc/internal/model"
	"zeroblog/application/user/rpc/internal/svc"
	"zeroblog/application/user/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *service.RegisterRequest) (*service.RegisterResponse, error) {
	// 当注册名字为空的时候，返回业务自定义错误码
	if len(in.Username) == 0 {
		return nil, errors.New("error name") //错误码：注册名不能为空
	}

	ret, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Username:   in.Username,
		Mobile:     in.Mobile,
		Avatar:     in.Avatar,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	if err != nil {
		logx.Errorf("Register req: %v error: %v", in, err)
		return nil, err
	}
	userId, err := ret.LastInsertId()
	if err != nil {
		logx.Errorf("LastInsertId error: %v", err)
		return nil, err
	}

	return &service.RegisterResponse{UserId: uint64(userId)}, nil
}