package queueservice

import (
	"MqttReceiver/pkg/queueservice/queueaccess"
	"MqttReceiver/pkg/queueservice/queuelogic"
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

func (q *QueueService) Enqueue(ctx context.Context, key string, message []byte) error {
	return q.service.Enqueue(ctx, key, message)
}
