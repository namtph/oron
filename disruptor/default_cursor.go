package disruptor

import "sync/atomic"

type DefaultCursor struct {
	ringLength int64
	rings      Ring
	idx        int64
}

func (c *DefaultCursor) StoreAndMove(e Event) error {
	err := c.rings.Set(c.idx, e)
	c.Next(1)
	return err
}

func (c *DefaultCursor) Store(e Event) error {
	return c.rings.Set(c.idx, e)
}

func (c *DefaultCursor) Load() (Event, error) {
	return c.rings.Get(c.idx)
}

func (c *DefaultCursor) Index() int64 {
	return c.idx
}

func (c *DefaultCursor) Next(step int64) {
	atomic.AddInt64(&c.idx,  (c.idx + step) % c.rings.Length())
}
