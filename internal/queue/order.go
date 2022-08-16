package queue

type (
	QueueInterface interface {
		Push() error
	}

	QueueMysql struct{}
)

func (*QueueMysql) Push() error {
	// 获取队列数据
	// 如果有数据

	return nil
}
