package utils

import (
	"strconv"
	"sync/atomic"
)

type Counter struct {
	cnt uint64
}

func (c *Counter) Inc() {
	atomic.AddUint64(&c.cnt, 1)
}

func (c *Counter) Value() uint64 {
	return c.cnt
}

func (c *Counter) String() string {
	return strconv.FormatUint(c.cnt, 10)
}
