package queuelogic

import (
	"context"
)

type ITrainQueue interface {
	Enqueue(ctx context.Context, key string, data []byte) error
}
type QueueBusinnessLogic struct {
	queueAbs ITrainQueue
}

func NewQueue(queueaccess ITrainQueue) *QueueBusinnessLogic {
	return &QueueBusinnessLogic{
		queueAbs: queueaccess,
	}
}

func (q *QueueBusinnessLogic) Enqueue(ctx context.Context, key string, data []byte) error {
	q.queueAbs.Enqueue(ctx, key, data)
	return nil
}
