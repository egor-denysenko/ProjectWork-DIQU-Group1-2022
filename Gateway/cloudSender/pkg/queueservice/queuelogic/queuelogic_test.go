package queuelogic

import (
	"context"
	"testing"
)

type MockQueue struct {
}

func (m *MockQueue) Enqueue(ctx context.Context, key string, data []byte) error {
	return nil
}

func TestEnqueue(t *testing.T) {
}
