package utils

import (
	"fmt"
	"sync"
)

func NewErrorCounter() ErrorCounter {
	return ErrorCounter{
		cnt: make(map[string]*Counter),
	}
}

type ErrorCounter struct {
	cnt   map[string]*Counter
	mutex sync.Mutex
}

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
