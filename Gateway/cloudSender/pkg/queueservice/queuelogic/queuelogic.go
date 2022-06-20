package queuelogic

import (
	"context"
)

type ITrainQueue interface {
	Dequeue(ctx context.Context, key string) []byte
}
type QueueBusinnessLogic struct {
	queueAbs ITrainQueue
}

func NewQueue(queueaccess ITrainQueue) *QueueBusinnessLogic {
	return &QueueBusinnessLogic{
		queueAbs: queueaccess,
	}
}

func (q *QueueBusinnessLogic) Dequeue(ctx context.Context, key string) []byte {
	err := q.queueAbs.Dequeue(ctx, key)
	if err != nil {
		return err
	}
	return nil
}
