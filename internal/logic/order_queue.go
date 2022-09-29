package logic

import (
	"container/list"
	"log"
	"time"
)

type _orderQueue struct {
	Logic
	card *list.List
}

var OrderQueue *_orderQueue

func InitCard() {
	OrderQueue = &_orderQueue{
		card: list.New(),
	}
	go OrderQueue.Card()
	OrderQueue.TestPush()
}
func (o *_orderQueue) Card() {
	for q := o.card.Front(); q != nil; q = q.Next() {
		log.Println(q.Value)
	}
	// for {
	// 	q := o.card.Front()
	// 	if q != nil {
	// 		log.Println(q.Value)
	// 		q = q.Next()
	// 	}
	// }
}

func (o *_orderQueue) TestPush() {
	for i := 0; i < 10; i++ {
		log.Println(i)
		o.card.PushBack(i)
		time.Sleep(time.Second * 10)
	}
}
