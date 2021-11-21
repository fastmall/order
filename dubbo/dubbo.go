package dubbo

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	_ "github.com/dubbogo/triple/pkg/triple"
	cartApi "github.com/fastmall/cart/api"
	customerApi "github.com/fastmall/customer/api"
	goodsApi "github.com/fastmall/goods/api"
	"github.com/fastmall/order/service"
)

var GoodsService = &goodsApi.GoodsServiceClientImpl{}
var CustomerService = &customerApi.CustomerServiceClientImpl{}
var CartService = &cartApi.CartServiceClientImpl{}

func StartDubbo() {
	config.SetConsumerService(GoodsService)
	config.SetConsumerService(CustomerService)
	config.SetConsumerService(CartService)
	config.SetProviderService(&service.OrderService{})
	err := config.Load(config.WithPath("conf/dubbo.yaml"))
	if err != nil {
		logger.Fatal(err)
		return
	}
}
