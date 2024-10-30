package svc

import (
	"zeroblog/application/applet/internal/config"
	"zeroblog/application/user/rpc/user"
	"zeroblog/pkg/interceptors"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	BizRedis *redis.Redis
	UserRPC  user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 自定义拦截器
	userRPC := zrpc.MustNewClient(c.UserRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:   c,
		BizRedis: redis.MustNewRedis(c.BizRedis, redis.WithPass(c.BizRedis.Pass)),
		UserRPC:  user.NewUser(userRPC),
	}
}
