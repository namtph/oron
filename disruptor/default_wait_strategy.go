package disruptor

import "time"

type DefaultWaitStrategy struct {

}

func (w *DefaultWaitStrategy) Gate(count int64) {
	time.Sleep(time.Nanosecond)
}

func (w *DefaultWaitStrategy) Idle(count int64) {
	time.Sleep(time.Second)
}
