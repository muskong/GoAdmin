package logic

import (
	"github.com/muskong/GoAdmin/internal/entity"
)

type _asyncWorkerCard struct {
	jobCard chan *entity.Order
	jobPay  chan *entity.Order
}

var Job *_asyncWorkerCard

// var jobCard chan *entity.Order

func init() {
	// Job.jobCard = make(chan *entity.Order, 99)
	// Job.jobPay = make(chan *entity.Order, 99)
	// go cardWorker(Job.jobCard)
	// go payWorker(Job.jobPay)
}

func cardWorker(jobCard <-chan *entity.Order) {
	const N = 5
	// 启动 5 个工作协程
	for i := 0; i < N; i++ {
		go func(id int) {
			for {
				data := <-jobCard
				// log.Printf("finish task: %#+v by worker %d\n", task, id)
				// time.Sleep(time.Second)
				OrderAsync.CardPublish(data)
			}
		}(i)
	}
}

func payWorker(jobCard <-chan *entity.Order) {
	const N = 5
	// 启动 5 个工作协程
	for i := 0; i < N; i++ {
		go func(id int) {
			for {
				data := <-jobCard
				// log.Printf("finish task: %#+v by worker %d\n", task, id)
				// time.Sleep(time.Second)
				OrderAsync.PayPublish(data.OrderNumber)
			}
		}(i)
	}
}
func (job *_asyncWorkerCard) PublishCard(data *entity.Order) {
	job.jobCard <- data
}

func (job *_asyncWorkerCard) PublishPay(data *entity.Order) {
	job.jobPay <- data
}
