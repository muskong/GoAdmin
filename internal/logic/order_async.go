package logic

import (
	"fmt"

	"github.com/muskong/GoAdmin/internal/entity"
	"github.com/muskong/GoCore/respond"
	"github.com/muskong/GoPkg/zaplog"
)

// 发起第三方请求
func CardPublish(order *entity.Order) {
	zaplog.Sugar.Info("CardPublish Start")
	defer func() {
		zaplog.Sugar.Info("CardPublish END")
	}()

	service, err := Product.WeightService(order.ProductCardId, order.ProductAmountId, order.ProductChannelId)
	if err != nil {
		zaplog.Sugar.Error(err)
		return
	}

	req := map[string]string{
		"order_number":   order.Uuid,
		"product_number": order.Product.ProductNumber,
		"card_number":    order.CardNumber,
		"card_password":  order.CardPassword,
		"notify_url":     fmt.Sprintf("public/notify/%s", OrderToken.GenerateToken(order)),
	}

	var rsp respond.Response
	err = Plugin.New(service.Class, service.Content).Send(req, &rsp)
	if err != nil {
		zaplog.Sugar.Error(err)
		return
	}

	if rsp.Message != "" {
		zaplog.Sugar.Error(rsp.Message)
		return
	}

	order.Queue = entity.OrderEntity.QueueCardRequest()
	err = Order.Update(*order)
	if err != nil {
		zaplog.Sugar.Error(err)
	}
}

// 第三方通知
func CardNotify(orderNumber string, data any) {
	order, err := Order.Detail(orderNumber)
	if err != nil {
		zaplog.Sugar.Error(err)
	}

	if order.State == entity.OrderEntity.StateSuccess() {
		zaplog.Sugar.Infof("订单已完成:%#+v", order)
	}

	var rsp respond.Response
	err = Plugin.New(order.Service.Class, order.Service.Content).Notify(data, &rsp)
	if err != nil {
		zaplog.Sugar.Error(err)
	}

	if rsp.Message != "" {
		zaplog.Sugar.Error(rsp.Message)
		order.Queue = entity.OrderEntity.QueueCardError()
	} else {
		order.Queue = entity.OrderEntity.QueueCardSuccess()
	}
	order.ResultCard = []any{rsp.Data}

	err = Order.Update(*order)
	if err != nil {
		zaplog.Sugar.Error(err)
	}

	go PayPublish(order.OrderNumber)
}

// 第三方通知
func PayPublish(orderNumber string) {
	order, err := Order.Detail(orderNumber)
	if err != nil {
		zaplog.Sugar.Error(err)
	}

	if order.State == entity.OrderEntity.StateSuccess() {
		zaplog.Sugar.Infof("订单已完成:%#+v", order)
	}

	// 获取支付
	pay, err := Pay.WeightService()
	if err != nil {
		zaplog.Sugar.Error(err)
	}

	data := map[string]string{
		"orderid":  "",
		"userid":   "userid",
		"money":    "money",
		"remarks":  "remarks",
		"username": "user",
	}
	var rsp respond.Response
	err = Plugin.New(pay.Class, pay.Content).Send(data, &rsp)
	if err != nil {
		zaplog.Sugar.Error(err)
	}

	if rsp.Message != "" {
		zaplog.Sugar.Error(rsp.Message)
		order.Queue = entity.OrderEntity.QueuePayError()
		order.State = entity.OrderEntity.StateError()
	} else {
		order.Queue = entity.OrderEntity.QueuePaySuccess()
		order.State = entity.OrderEntity.StateSuccess()
	}
	order.ResultPay = []any{rsp.Data}

	err = Order.Update(*order)
	if err != nil {
		zaplog.Sugar.Error(err)
	}
}
