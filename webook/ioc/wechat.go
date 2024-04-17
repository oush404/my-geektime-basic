package ioc

import (
	"my-geektime-basic/webook/internal/service/oauth2/wechat"
	"my-geektime-basic/webook/pkg/logger"
)

func InitWechatService(l logger.LoggerV1) wechat.Service {
	//appID, ok := os.LookupEnv("WECHAT_APP_ID")
	//if !ok {
	//	panic("找不到环境变量 WECHAT_APP_ID")
	//}
	//appSecret, ok := os.LookupEnv("WECHAT_APP_SECRET")
	//if !ok {
	//	panic("找不到环境变量 WECHAT_APP_SECRET")
	//}
	appID := "123"
	appSecret := "123"
	return wechat.NewService(appID, appSecret, l)
}
