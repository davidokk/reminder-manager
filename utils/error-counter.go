package utils

import (
	"fmt"
	"sync"
)

// NewErrorCounter initialized new ErrorCounter
func NewErrorCounter() ErrorCounter {
	return ErrorCounter{
		cnt: make(map[string]*Counter),
	}
}

// ErrorCounter thread-safe errors counter
type ErrorCounter struct {
	cnt   map[string]*Counter
	mutex sync.Mutex
}

// Inc increments the given error counter
func (ec *ErrorCounter) Inc(e string) {
	ec.mutex.Lock()
	defer ec.mutex.Unlock()

	if _, ok := ec.cnt[e]; !ok {
		ec.cnt[e] = &Counter{}
	}
	ec.cnt[e].Inc()
}

func (ec *ErrorCounter) String() string {
	ec.mutex.Lock()
	defer ec.mutex.Unlock()

	return fmt.Sprint(ec.cnt)
}
