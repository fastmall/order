package dubbo

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	goodsApi "github.com/fastmall/goods/api"
	"github.com/fastmall/order/service"
)

var GoodsService = &goodsApi.GoodsServiceClientImpl{}

func StartDubbo() {
	config.SetConsumerService(GoodsService)
	config.SetProviderService(&service.OrderService{})
	err := config.Load(config.WithPath("conf/dubbo.yaml"))
	if err != nil {
		logger.Fatal(err)
		return
	}
	err = config.Load(config.WithPath("conf/dubbo.yaml"))
	if err != nil {
		logger.Fatal(err)
		return
	}
}
