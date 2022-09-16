package logic

import (
	"container/list"
	"log"
)

type _orderQueue struct {
	Logic
	card *list.List
}

var OrderQueue *_orderQueue = &_orderQueue{
	card: list.New(),
}

func InitCard() {
	// OrderQueue = &_orderQueue{
	// 	card: list.New(),
	// }
	OrderQueue.Card()
}
func (o *_orderQueue) Card() {
	for {
		log.Println(o.card.Len())
	}
}

// func QueueTestGaobei() {
// 	plug, err := plugin.Open("./plugins/plugin_GaoBei.so")
// 	if err != nil {
// 		log.Println("queueTest", err)
// 	}
// 	gaoBei, err := plug.Lookup("GaoBei")
// 	if err != nil {
// 		log.Println("queueTest Lookup", err)
// 	}
// 	api, ok := gaoBei.(Api)
// 	if !ok {
// 		log.Println("queueTest ok fail")
// 		return
// 	}

// 	param := map[string]any{
// 		"test": 1233,
// 	}
// 	var res map[string]any

// 	err = api.Send(param, &res)
// 	if err != nil {
// 		log.Println("queueTest Send fail", err)
// 		return
// 	}

// 	log.Println("queueTest Send success")
// 	log.Println(res)
// }

// func QueueTestShouKaYun() {
// 	plug, err := plugin.Open("./plugins/shoukayun.so")
// 	if err != nil {
// 		log.Println("queueTest", err)
// 	}
// 	sky, err := plug.Lookup("ShouKaYun")
// 	if err != nil {
// 		log.Println("queueTest Lookup", err)
// 	}
// 	api, ok := sky.(ApiInterface)
// 	if !ok {
// 		log.Println("queueTest ok fail")
// 		return
// 	}

// 	var param, respond map[string]any
// 	param["test"] = "ShouKaYun"
// 	err = api.Send(param, respond)
// 	if err != nil {
// 		log.Println("queueTest Send fail", err)
// 		return
// 	}

// 	log.Println("queueTest Send fail", param, respond)
// }
