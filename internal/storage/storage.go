package storage

import (
	"errors"
	"sort"
	"time"

	"gitlab.ozon.dev/davidokk/reminder-manager/utils"
)

var data []*Reminder

var IdNotExistsError = errors.New("given id doesn't exist")
var IdAlreadyExistsError = errors.New("given id already exist")

func init() {
	data = make([]*Reminder, 0)
}

// returns index of min date after or equal to given
func firstAfterOrEqual(date time.Time) int {
	return sort.Search(len(data), func(i int) bool {
		return data[i].Date.After(date) || data[i].Date.Equal(date)
	})
}

func Add(rem *Reminder) error {
	if _, err := indexById(rem.ID); err == nil {
		return IdAlreadyExistsError
	}
	index := firstAfterOrEqual(rem.Date)
	data = utils.Insert(data, rem, index)
	return nil
}

func RemindersForDays(count int) []*Reminder {
	if count < 1 {
		return nil
	}
	l := firstAfterOrEqual(utils.UpToDay(time.Now()))
	r := firstAfterOrEqual(utils.UpToDay(time.Now()).Add(24 * time.Hour * time.Duration(count)))
	if l == r {
		return nil
	}
	return utils.Clone(data[l:r])
}

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

func RemoveOutdated() int {
	outdated := OutdatedCount()
	data = data[outdated:]
	return outdated
}

func OutdatedCount() (cnt int) {
	return firstAfterOrEqual(utils.UpToDay(time.Now()))
}

func RemoveById(id uint64) error {
	index, err := indexById(id)
	if err == nil {
		data = utils.Remove(data, index)
	}
	return err
}

func Edit(id uint64, newText string) error {
	index, err := indexById(id)
	if err == nil {
		data[index].What = newText
	}
	return err
}

func indexById(id uint64) (int, error) {
	for i, cur := range data {
		if cur.ID == id {
			return i, nil
		}
	}
	return -1, IdNotExistsError
}

func Data() []*Reminder {
	return utils.Clone(data)
}
