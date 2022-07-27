package storage

import (
	"context"
	"github.com/pkg/errors"
	"sort"
	"sync"
	"time"

	"gitlab.ozon.dev/davidokk/reminder-manager/utils"
)

const poolSize = 10
const waitingTime = 100 * time.Millisecond

var data []*Reminder
var poolCh chan struct{}
var mutex sync.RWMutex

// possible errors
var (
	ErrorIDNotExists     = errors.New("given id doesn't exist")
	ErrorIDAlreadyExists = errors.New("given id already exist")
	ErrorTimeoutExceeded = errors.New("timeout exceeded")
)

func init() {
	data = make([]*Reminder, 0)
	poolCh = make(chan struct{}, poolSize)
}

// firstAfterOrEqual returns index of min date after or equal to given
func firstAfterOrEqual(date time.Time) int {
	return sort.Search(len(data), func(i int) bool {
		return data[i].Date.After(date) || data[i].Date.Equal(date)
	})
}

type operationType uint8

const (
	READ  operationType = 0
	WRITE               = 1
)

func takeWorker(t operationType) error {
	ctx, cancel := context.WithTimeout(context.Background(), waitingTime)
	defer cancel()
	select {
	case poolCh <- struct{}{}:
		if t == READ {
			mutex.RLock()
		} else if t == WRITE {
			mutex.Lock()
		}
		return nil
	case <-ctx.Done():
		return ErrorTimeoutExceeded
	}
}

func returnWorker(t operationType) {
	<-poolCh
	if t == READ {
		mutex.RUnlock()
	} else if t == WRITE {
		mutex.Unlock()
	}
}

// Data returns all reminders as slice
func Data() ([]*Reminder, error) {
	if err := takeWorker(READ); err != nil {
		return nil, err
	}
	defer returnWorker(READ)

	return utils.Clone(data), nil
}

// Add adds a new Reminder into storage
func Add(rem *Reminder) error {
	if err := takeWorker(WRITE); err != nil {
		return err
	}
	defer returnWorker(WRITE)

	if _, err := indexById(rem.ID); err == nil {
		return ErrorIDAlreadyExists
	}
	index := firstAfterOrEqual(rem.Date)
	data = utils.Insert(data, rem, index)
	return nil
}

// RemoveById removes Reminder with given ID
func RemoveById(id uint64) error {
	if err := takeWorker(WRITE); err != nil {
		return err
	}
	defer returnWorker(WRITE)

	index, err := indexById(id)
	if err == nil {
		data = utils.Remove(data, index)
	}
	return err
}

// Edit allows to change the text of Reminder with given ID
func Edit(id uint64, newText string) error {
	if err := takeWorker(WRITE); err != nil {
		return err
	}
	defer returnWorker(WRITE)

	index, err := indexById(id)
	if err == nil {
		data[index].Text = newText
	}
	return err
}

// AsStrings applies Reminder.ToString to each Reminder
// and return the resulting slice
func AsStrings(rem []*Reminder) []string {
	if rem == nil {
		return nil
	}
	str := make([]string, 0, len(rem))
	for _, cur := range rem {
		str = append(str, cur.ToString())
	}
	return str
}

func indexById(id uint64) (int, error) {
	for i, cur := range data {
		if cur.ID == id {
			return i, nil
		}
	}
	return -1, ErrorIDNotExists
}
