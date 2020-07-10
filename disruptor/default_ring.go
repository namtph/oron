package disruptor

import "sync"

type DefaultRing struct {
	ringLength int64
	events     []Event
	lock       sync.Mutex
}

func (r *DefaultRing) Get(index int64) (Event, error) {
	if index > r.ringLength || index < 0 {
		return nil, errOutOfRange
	}
	r.lock.Lock()
	e := r.events[index]
	r.lock.Unlock()
	return e, nil
}

func (r *DefaultRing) Set(index int64, e Event) error {
	if index > r.ringLength || index < 0{
		return errOutOfRange
	}
	r.lock.Lock()
	r.events[index] = e
	r.lock.Unlock()

	return nil
}

func (r *DefaultRing) Length() int64 {
	return r.ringLength
}