package queueservice

import (
	"context"
	"serialReciever/pkg/queueservice/queueaccess"
	"serialReciever/pkg/queueservice/queuelogic"
)

type ConsumerService queuelogic.QueueBusinnessLogic

type QueueService struct {
	service *queuelogic.QueueBusinnessLogic
}

func QueueServiceFactory() *QueueService {
	return &QueueService{
		service: queuelogic.NewQueue(queueaccess.NewMessageQueue()),
	}
}

func (q *QueueService) Enqueue(ctx context.Context, key string, data []byte) error {
	return q.service.Enqueue(ctx, key, data)
}
