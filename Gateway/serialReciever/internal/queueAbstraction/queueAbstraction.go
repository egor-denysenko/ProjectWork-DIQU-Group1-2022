package queueAbstraction

type QueueActions interface {
	Enqueue(VagonMessage []byte) error
}

func NewQueue() QueueActions {
	return QueueActions
}
