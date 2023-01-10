package poller

import (
	"context"
	"fmt"
	"go.uber.org/goleak"
	"testing"
)

// https://github.com/uber-go/goleak
// Goroutine 泄漏检测器，帮助避免 Goroutine 泄漏
func TestMain(m *testing.M) {
	fmt.Println("start")
	goleak.VerifyTestMain(m)
}

func TestPoller(t *testing.T) {
	producer := NewPoller(5)
	producer.Poll(context.Background())
}
