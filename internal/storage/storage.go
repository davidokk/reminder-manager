package storage

import (
	"context"
	"sort"
	"sync"
	"time"

	"gitlab.ozon.dev/davidokk/reminder-manager/utils"

	"github.com/pkg/errors"

	"gitlab.ozon.dev/davidokk/reminder-manager/config"
)

var data []*Reminder
var poolCh chan struct{}
var mutex sync.RWMutex

// possible errors
var (
	ErrorIDNotExists     = errors.New("given id doesn't exist")
	ErrorIDAlreadyExists = errors.New("given id already exist")
)

// Init initializes the storage
func Init() {
	data = make([]*Reminder, 0)
	poolCh = make(chan struct{}, config.App.Storage.PoolSize)
}

// firstAfterOrEqual returns index of min date after or equal to given
func firstAfterOrEqual(date time.Time) int {
	return sort.Search(len(data), func(i int) bool {
		return data[i].Date.After(date) || data[i].Date.Equal(date)
	})
}

type operationType uint8

const (
	read  operationType = 0
	write operationType = 1
)

func takeWorker(c context.Context, t operationType) error {
	ctx, cancel := context.WithTimeout(c, config.App.Storage.WaitingTime)
	defer cancel()
	select {
	case poolCh <- struct{}{}:
		if t == read {
			mutex.RLock()
		} else if t == write {
			mutex.Lock()
		}
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func returnWorker(t operationType) {
	<-poolCh
	if t == read {
		mutex.RUnlock()
	} else if t == write {
		mutex.Unlock()
	}
}

// Data returns all reminders as slice
func Data(ctx context.Context) ([]*Reminder, error) {
	if err := takeWorker(ctx, read); err != nil {
		return nil, err
	}
	defer returnWorker(read)

	return utils.Clone(data), nil
}

// Add adds a new Reminder into storage
func Add(ctx context.Context, rem *Reminder) error {
	if err := takeWorker(ctx, write); err != nil {
		return err
	}
	defer returnWorker(write)

	if _, err := indexByID(rem.ID); err == nil {
		return ErrorIDAlreadyExists
	}
	index := firstAfterOrEqual(rem.Date)
	data = utils.Insert(data, rem, index)
	return nil
}

// RemoveByID removes Reminder with given ID
func RemoveByID(ctx context.Context, id uint64) error {
	if err := takeWorker(ctx, write); err != nil {
		return err
	}
	defer returnWorker(write)

	index, err := indexByID(id)
	if err == nil {
		data = utils.Remove(data, index)
	}
	return err
}

// Edit allows to change the text of Reminder with given ID
func Edit(ctx context.Context, id uint64, newText string) error {
	if err := takeWorker(ctx, write); err != nil {
		return err
	}
	defer returnWorker(write)

	index, err := indexByID(id)
	if err == nil {
		data[index].Text = newText
	}
	return err
}

// AsStrings applies Reminder.String to each Reminder
// and return the resulting slice
func AsStrings(rem []*Reminder) []string {
	if rem == nil {
		return nil
	}
	str := make([]string, 0, len(rem))
	for _, cur := range rem {
		str = append(str, cur.String())
	}
	return str
}

func indexByID(id uint64) (int, error) {
	for i, cur := range data {
		if cur.ID == id {
			return i, nil
		}
	}
	return -1, ErrorIDNotExists
}
