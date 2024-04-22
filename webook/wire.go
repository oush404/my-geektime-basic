//go:build wireinject

package main

import (
	"github.com/google/wire"
	"my-geektime-basic/webook/internal/events/article"
	"my-geektime-basic/webook/internal/repository"
	"my-geektime-basic/webook/internal/repository/cache"
	"my-geektime-basic/webook/internal/repository/dao"
	"my-geektime-basic/webook/internal/service"
	"my-geektime-basic/webook/internal/web"
	ijwt "my-geektime-basic/webook/internal/web/jwt"
	"my-geektime-basic/webook/ioc"
)

var interactiveSvcSet = wire.NewSet(dao.NewGORMInteractiveDAO,
	cache.NewInteractiveRedisCache,
	repository.NewCachedInteractiveRepository,
	service.NewInteractiveService,
)

func InitWebServer() *App {
	wire.Build(
		// 第三方依赖
		ioc.InitRedis, ioc.InitDB,
		ioc.InitLogger,
		ioc.InitSaramaClient,
		ioc.InitSyncProducer,
		// DAO 部分
		dao.NewUserDAO,
		dao.NewArticleGORMDAO,

		interactiveSvcSet,

		article.NewSaramaSyncProducer,
		article.NewInteractiveReadEventConsumer,
		ioc.InitConsumers,

		// cache 部分
		cache.NewCodeCache, cache.NewUserCache,
		cache.NewArticleRedisCache,

		// repository 部分
		repository.NewCachedUserRepository,
		repository.NewCodeRepository,
		repository.NewCachedArticleRepository,

		// Service 部分
		ioc.InitSMSService,
		ioc.InitWechatService,
		service.NewUserService,
		service.NewCodeService,
		service.NewArticleService,

		// handler 部分
		web.NewUserHandler,
		web.NewArticleHandler,
		ijwt.NewRedisJWTHandler,
		web.NewOAuth2WechatHandler,
		ioc.InitGinMiddlewares,
		ioc.InitWebServer,

		wire.Struct(new(App), "*"),
	)
	return new(App)
}
