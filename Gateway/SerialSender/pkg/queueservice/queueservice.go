package queueservice

import (
	"SerialSender/pkg/queueservice/queueaccess"
	"SerialSender/pkg/queueservice/queuelogic"
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

func (q *QueueService) Connect() error {
	return q.service.Connect()
}
func (q *QueueService) Dequeue(ctx context.Context, key string, queueDataChan chan<- *string) {
	q.service.Dequeue(ctx, key, queueDataChan)
}

func (q *QueueService) Enqueue(ctx context.Context, key string, message []byte) error {
	return q.service.Enqueue(ctx, key, message)
}
