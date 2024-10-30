package logic

import (
	"context"
	"errors"
	"strings"

	"zeroblog/application/applet/internal/svc"
	"zeroblog/application/applet/internal/types"
	"zeroblog/application/user/rpc/user"
	"zeroblog/pkg/encrypt"
	"zeroblog/pkg/jwt"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	prefixActivation = "biz#activation#%s"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (*types.RegisterResponse, error) {
	//去除多余空格并检查输入有效性
	req.Name = strings.TrimSpace(req.Name)

	req.Mobile = strings.TrimSpace(req.Mobile)
	if len(req.Mobile) == 0 {
		return nil, errors.New("error moblie") //错误码：手机号不能为空
	}

	req.Password = strings.TrimSpace(req.Password)
	if len(req.Password) == 0 {
		return nil, errors.New("error passwd") //错误码：密码不能为空
	} else {
		req.Password = encrypt.EncPassword(req.Password) //密码需要加密存储
	}

	req.VerificationCode = strings.TrimSpace(req.VerificationCode)
	if len(req.VerificationCode) == 0 {
		return nil, errors.New("error verification code") //错误码：短信验证码不能为空
	}
	err := checkVerificationCode(l.svcCtx.BizRedis, req.Mobile, req.VerificationCode)
	if err != nil { //验证码错误
		logx.Errorf("checkVerificationCode error: %v", err)
		return nil, err
	}

	mobile, err := encrypt.EncMobile(req.Mobile) //手机号需要加密存储
	if err != nil {
		logx.Errorf("EncMobile mobile: %s error: %v", req.Mobile, err)
		return nil, err
	}

	u, err := l.svcCtx.UserRPC.FindByMobile(l.ctx, &user.FindByMobileRequest{ //检查手机号是否未注册
		Mobile: mobile,
	})
	if err != nil {
		logx.Errorf("FindByMobile error: %v", err)
		return nil, err
	}
	if u != nil && u.UserId > 0 {
		return nil, errors.New("mobile registried") //错误码：手机号已注册
	}

	regRet, err := l.svcCtx.UserRPC.Register(l.ctx, &user.RegisterRequest{ //进行注册
		Username: req.Name,
		Mobile:   mobile,
	})
	if err != nil {
		logx.Errorf("Register error: %v", err)
		return nil, err
	}

	//注册成功后生成token
	token, err := jwt.BuildTokens(jwt.TokenOptions{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		Fields: map[string]interface{}{
			"userId": regRet.UserId,
		},
	})

	if err != nil {
		logx.Errorf("BuildTokens error: %v", err)
		return nil, err
	}

	_ = delActivationCache(req.Mobile, req.VerificationCode, l.svcCtx.BizRedis) //完成所有操作后删除缓存

	return &types.RegisterResponse{
		UserId: regRet.UserId,
		Token: types.Token{
			AccessToken:  token.AccessToken,
			AccessExpire: token.AccessExpire,
		},
	}, nil
}

// 验证码校验
func checkVerificationCode(rds *redis.Redis, moblie, code string) error {
	cacheCode, err := getActivationCache(moblie, rds)
	if err != nil {
		return err
	}
	if cacheCode == "" {
		return errors.New("verification code expired")
	}
	if cacheCode != code {
		return errors.New("verification code failed")
	}

	return nil
}
