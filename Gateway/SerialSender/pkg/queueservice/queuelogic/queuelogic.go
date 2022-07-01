package queuelogic

import (
	"context"
)

type ITrainQueue interface {
	Connect() error
	Dequeue(ctx context.Context, key string, queueDataChan chan<- *string)
	Enqueue(ctx context.Context, key string, message []byte) error
}
type QueueBusinnessLogic struct {
	queueAbs ITrainQueue
}

func NewQueue(queueaccess ITrainQueue) *QueueBusinnessLogic {
	return &QueueBusinnessLogic{
		queueAbs: queueaccess,
	}
}

func (q *QueueBusinnessLogic) Connect() error {
	return q.queueAbs.Connect()
}

func (q *QueueBusinnessLogic) Dequeue(ctx context.Context, key string, queueDataChan chan<- *string) {
	q.queueAbs.Dequeue(ctx, key, queueDataChan)
}

func (q *QueueBusinnessLogic) Enqueue(ctx context.Context, key string, message []byte) error {
	err := q.queueAbs.Enqueue(ctx, key, message)
	if err != nil {
		return err
	}
	return nil
}
