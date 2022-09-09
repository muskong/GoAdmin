package logic

import (
	"github.com/muskong/GoAdmin/internal/entity"
	"github.com/muskong/GoPkg/broker"
	"github.com/muskong/GoPkg/zaplog"
)

var Broker = broker.NewClient()

func OrderCardTopic() {
	defer Broker.Close()
	Broker.SetConditions(100)

	ch, err := Broker.Subscribe("OrderCard")
	if err != nil {
		zaplog.Sugar.Panic(err)
	}

	for {
		order := Broker.GetPayLoad(ch).(*entity.Order)

		if order != nil {
			service, err := Product.WeightService(order.ProductCardId, order.ProductAmountId, order.ProductChannelId)
			if err != nil {

			}

		}

		zaplog.Sugar.Info(order)
	}
}

func OrderCardPublish(payload any) {
	go Broker.Publish("OrderCard", payload)
}
