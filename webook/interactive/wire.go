//go:build wireinject

package main

import (
	"github.com/google/wire"
	"my-geektime-basic/webook/interactive/events"
	"my-geektime-basic/webook/interactive/grpc"
	"my-geektime-basic/webook/interactive/ioc"
	repository2 "my-geektime-basic/webook/interactive/repository"
	cache2 "my-geektime-basic/webook/interactive/repository/cache"
	dao2 "my-geektime-basic/webook/interactive/repository/dao"
	service2 "my-geektime-basic/webook/interactive/service"
)

var thirdPartySet = wire.NewSet(ioc.InitSrcDB,
	ioc.InitDstDB,
	ioc.InitDoubleWritePool,
	ioc.InitBizDB,
	ioc.InitLogger,
	ioc.InitSaramaClient,
	ioc.InitSaramaSyncProducer,
	ioc.InitRedis)

var interactiveSvcSet = wire.NewSet(dao2.NewGORMInteractiveDAO,
	cache2.NewInteractiveRedisCache,
	repository2.NewCachedInteractiveRepository,
	service2.NewInteractiveService,
)

func InitApp() *App {
	wire.Build(thirdPartySet,
		interactiveSvcSet,
		grpc.NewInteractiveServiceServer,
		events.NewInteractiveReadEventConsumer,
		ioc.InitInteractiveProducer,
		ioc.InitFixerConsumer,
		ioc.InitConsumers,
		ioc.NewGrpcxServer,
		ioc.InitGinxServer,
		wire.Struct(new(App), "*"),
	)
	return new(App)
}
