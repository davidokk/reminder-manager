package utils

import (
	"strconv"
	"sync/atomic"
)

// Counter thread-safe counter
type Counter struct {
	cnt uint64
}

// Inc increments value
func (c *Counter) Inc() {
	atomic.AddUint64(&c.cnt, 1)
}

// Value returns a counter value
func (c *Counter) Value() uint64 {
	return c.cnt
}

func (c *Counter) String() string {
	return strconv.FormatUint(c.cnt, 10)
}
