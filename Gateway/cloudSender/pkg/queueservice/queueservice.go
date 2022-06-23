package queueservice

import (
	"cloudSender/pkg/queueservice/queueaccess"
	"cloudSender/pkg/queueservice/queuelogic"
	"context"
)

type ConsumerService queuelogic.QueueBusinnessLogic

type QueueService struct {
	service *queuelogic.QueueBusinnessLogic
}

func FactoryQueueService() *QueueService {
	return &QueueService{
		service: queuelogic.NewQueue(queueaccess.NewMessageQueue()),
	}
}

func (q *QueueService) Dequeue(ctx context.Context, key string) ([]byte, error) {
	return q.service.Dequeue(ctx, key)
}

func (q *QueueService) Enqueue(ctx context.Context, key string, message []byte) error {
	return q.service.Enqueue(ctx, key, message)
}
