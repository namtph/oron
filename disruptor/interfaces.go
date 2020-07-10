package disruptor

type Event interface {

}

type Cursor interface {
	Store(e Event) error
	Load() (Event, error)
	Index() int64
	Next(step int64)
}

type WaitStrategy interface {
	Gate(count int64)
	Idle(count int64)
}

type Ring interface {
	Get(index int64) (Event, error)
	Set(index int64, event Event) error
	Length() int64
}