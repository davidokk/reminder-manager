package local

import (
	"context"
	"sort"
	"sync"
	"time"

	"gitlab.ozon.dev/davidokk/reminder-manager/internal/models"

	"gitlab.ozon.dev/davidokk/reminder-manager/utils"

	"github.com/pkg/errors"

	"gitlab.ozon.dev/davidokk/reminder-manager/config"
)

type storage struct {
	lastID uint64
	data   []*models.Reminder
	poolCh chan struct{}
	mutex  sync.RWMutex
}

// ErrorIDNotExists possible error
var ErrorIDNotExists = errors.New("given id doesn't exist")

// New initializes the storage
func New() *storage {
	return &storage{
		data:   make([]*models.Reminder, 0),
		poolCh: make(chan struct{}, config.App.Storage.PoolSize),
	}
}

// firstAfterOrEqual returns index of min date after or equal to given
func (s *storage) firstAfterOrEqual(date time.Time) int {
	return sort.Search(len(s.data), func(i int) bool {
		return s.data[i].Date.After(date) || s.data[i].Date.Equal(date)
	})
}

type operationType uint8

const (
	read  operationType = 0
	write operationType = 1
)

func (s *storage) takeWorker(c context.Context, t operationType) error {
	ctx, cancel := context.WithTimeout(c, config.App.Storage.WaitingTime)
	defer cancel()
	select {
	case s.poolCh <- struct{}{}:
		if t == read {
			s.mutex.RLock()
		} else if t == write {
			s.mutex.Lock()
		}
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *storage) returnWorker(t operationType) {
	<-s.poolCh
	if t == read {
		s.mutex.RUnlock()
	} else if t == write {
		s.mutex.Unlock()
	}
}

// ListReminders returns all reminders as slice
func (s *storage) ListReminders(ctx context.Context) ([]*models.Reminder, error) {
	if err := s.takeWorker(ctx, read); err != nil {
		return nil, err
	}
	defer s.returnWorker(read)

	return utils.Clone(s.data), nil
}

// GetReminder returns a reminder with given id
func (s *storage) GetReminder(ctx context.Context, id uint64) (*models.Reminder, error) {
	if err := s.takeWorker(ctx, read); err != nil {
		return nil, err
	}
	defer s.returnWorker(read)

	index, err := s.indexByID(id)
	if err != nil {
		return nil, err
	}
	return s.data[index], nil
}

// CreateReminder adds a new Reminder into storage
func (s *storage) CreateReminder(ctx context.Context, date time.Time, text string) (*models.Reminder, error) {
	if err := s.takeWorker(ctx, write); err != nil {
		return nil, err
	}
	defer s.returnWorker(write)

	s.lastID++
	rem := &models.Reminder{
		ID:   s.lastID,
		Text: text,
		Date: date,
	}

	index := s.firstAfterOrEqual(rem.Date)
	s.data = utils.Insert(s.data, rem, index)
	return rem, nil
}

// RemoveReminder removes Reminder with given ID
func (s *storage) RemoveReminder(ctx context.Context, id uint64) error {
	if err := s.takeWorker(ctx, write); err != nil {
		return err
	}
	defer s.returnWorker(write)

	index, err := s.indexByID(id)
	if err == nil {
		s.data = utils.Remove(s.data, index)
	}
	return err
}

// UpdateReminder allows to change the text of Reminder with given ID
func (s *storage) UpdateReminder(ctx context.Context, id uint64, text string) error {
	if err := s.takeWorker(ctx, write); err != nil {
		return err
	}
	defer s.returnWorker(write)

	index, err := s.indexByID(id)
	if err == nil {
		s.data[index].Text = text
	}
	return err
}

// AsStrings applies Reminder.String to each Reminder
// and return the resulting slice
func AsStrings(rem []*models.Reminder) []string {
	if rem == nil {
		return nil
	}
	str := make([]string, 0, len(rem))
	for _, cur := range rem {
		str = append(str, cur.String())
	}
	return str
}

func (s *storage) indexByID(id uint64) (int, error) {
	for i, cur := range s.data {
		if cur.ID == id {
			return i, nil
		}
	}
	return -1, ErrorIDNotExists
}
