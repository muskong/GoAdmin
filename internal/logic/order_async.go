package logic

import (
	"github.com/muskong/GoAdmin/internal/entity"
	"github.com/muskong/GoCore/respond"
	"github.com/muskong/GoPkg/zaplog"
)

// 发起第三方请求
func ExternalPublish(order *entity.Order) {
	service, err := Product.WeightService(order.ProductCardId, order.ProductAmountId, order.ProductChannelId)
	if err != nil {
		zaplog.Sugar.Error(err)
	}

	api := Plugin(service.Class)

	req := map[string]any{
		"card_number":   order.CardNumber,
		"card_password": order.CardPassword,
		"card_cvv":      order.CardCvv,
	}
	var rsp respond.Response
	err = api.Send(req, &rsp)
	if err != nil {
		zaplog.Sugar.Error(err)
	}

	if rsp.Message != "" {
		zaplog.Sugar.Error(rsp.Message)
	}

	order.Queue = entity.OrderEntity.QueueCardRequest()
	err = Order.Update(*order)
	if err != nil {
		zaplog.Sugar.Error(err)
	}
}

// 第三方通知
func ExternalNotify(orderNumber string) {
	order, err := Order.Detail(orderNumber)
	if err != nil {
		zaplog.Sugar.Error(err)
	}

	if order.State != entity.OrderEntity.StateHang() {

	}
}
