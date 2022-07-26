package storage

import (
	"errors"
	"sort"
	"sync"
	"time"

	"gitlab.ozon.dev/davidokk/reminder-manager/utils"
)

const poolSize = 10

var data []*Reminder
var poolCh chan struct{}
var mutex sync.RWMutex

// possible errors
var (
	ErrorIDNotExists     = errors.New("given id doesn't exist")
	ErrorIDAlreadyExists = errors.New("given id already exist")
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

// Add adds a new Reminder into storage
func Add(rem *Reminder) error {
	poolCh <- struct{}{}
	mutex.Lock()
	defer func() {
		mutex.Unlock()
		<-poolCh
	}()

	if _, err := indexById(rem.ID); err == nil {
		return ErrorIDAlreadyExists
	}
	index := firstAfterOrEqual(rem.Date)
	data = utils.Insert(data, rem, index)
	return nil
}

// RemindersForDays returns list of reminders for the next count days
func RemindersForDays(count int) []*Reminder {
	if count < 1 {
		return nil
	}

	poolCh <- struct{}{}
	mutex.RLock()
	defer func() {
		mutex.RUnlock()
		<-poolCh
	}()

	l := firstAfterOrEqual(utils.UpToDay(time.Now()))
	r := firstAfterOrEqual(utils.UpToDay(time.Now()).Add(24 * time.Hour * time.Duration(count)))
	if l == r {
		return nil
	}
	return utils.Clone(data[l:r])
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

// RemoveOutdated removes from storage all outdated entries
// return count of deleted entries
func RemoveOutdated() int {
	poolCh <- struct{}{}
	mutex.Lock()
	defer func() {
		mutex.Unlock()
		<-poolCh
	}()

	outdated := outdatedCount()
	data = data[outdated:]
	return outdated
}

func outdatedCount() (cnt int) {
	return firstAfterOrEqual(utils.UpToDay(time.Now()))
}

// OutdatedCount returns count of outdated records
func OutdatedCount() int {
	poolCh <- struct{}{}
	mutex.RLock()
	defer func() {
		mutex.RUnlock()
		<-poolCh
	}()

	return firstAfterOrEqual(utils.UpToDay(time.Now()))
}

// RemoveById removes Reminder with given ID
func RemoveById(id uint64) error {
	poolCh <- struct{}{}
	mutex.Lock()
	defer func() {
		mutex.Unlock()
		<-poolCh
	}()

	index, err := indexById(id)
	if err == nil {
		data = utils.Remove(data, index)
	}
	return err
}

// Edit allows to change the text of Reminder with given ID
func Edit(id uint64, newText string) error {
	poolCh <- struct{}{}
	mutex.Lock()
	defer func() {
		mutex.Unlock()
		<-poolCh
	}()

	index, err := indexById(id)
	if err == nil {
		data[index].Text = newText
	}
	return err
}

func indexById(id uint64) (int, error) {
	for i, cur := range data {
		if cur.ID == id {
			return i, nil
		}
	}
	return -1, ErrorIDNotExists
}

// Data returns all reminders as slice
func Data() []*Reminder {
	poolCh <- struct{}{}
	mutex.RLock()
	defer func() {
		mutex.RUnlock()
		<-poolCh
	}()

	return utils.Clone(data)
}
