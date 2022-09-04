package cached

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/config"
	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/internal/models"
	storagePkg "gitlab.ozon.dev/davidokk/reminder-manager/data-service/internal/storage"
	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/internal/storage/repository"
)

type storage struct {
	repo  repository.Interface
	cache *redis.Client
}

// New returns storage with given pool and cache
func New(r repository.Interface) storagePkg.RemindersStorage {
	return &storage{
		repo: r,
		cache: redis.NewClient(&redis.Options{
			Addr:     config.App.Cache.Address,
			DB:       0,
			Password: "",
		}),
	}
}

func (s *storage) cacheSet(rem *models.Reminder) {
	if rem != nil {
		res := s.cache.Set(strconv.FormatUint(rem.ID, 10), rem, config.App.Cache.Expiration)
		if res.Err() != nil {
			log.Printf("cache set: %s", res.Err().Error())
		}
	}
}

func (s *storage) cacheGet(id uint64) *models.Reminder {
	res := s.cache.Get(strconv.FormatUint(id, 10))
	if res.Err() != nil {
		if res.Err() == redis.Nil {
			miss.Inc()
		} else {
			log.Printf("cache get: %s", res.Err().Error())
		}
		return nil
	}
	hit.Inc()

	var rem models.Reminder
	err := res.Scan(&rem)

	if err != nil {
		log.Printf("cache scan: %s", err)
		return nil
	}
	return &rem
}

func (s *storage) CreateReminder(ctx context.Context, date time.Time, text string) (*models.Reminder, error) {
	rem, err := s.repo.CreateReminder(ctx, date, text)
	if err == nil {
		s.cacheSet(rem)
	}
	return rem, err
}

func (s *storage) GetReminder(ctx context.Context, id uint64) (*models.Reminder, error) {
	cached := s.cacheGet(id)
	if cached != nil {
		return cached, nil
	}
	rem, err := s.repo.GetReminder(ctx, id)
	if err == nil {
		s.cacheSet(rem)
	}
	return rem, err
}

func (s *storage) UpdateReminder(ctx context.Context, id uint64, text string) error {
	err := s.repo.UpdateReminder(ctx, id, text)
	if err == nil {
		s.cache.Del(strconv.FormatUint(id, 10))
	}
	return err
}

func (s *storage) RemoveReminder(ctx context.Context, id uint64) error {
	err := s.repo.RemoveReminder(ctx, id)
	if err == nil {
		s.cache.Del(strconv.FormatUint(id, 10))
	}
	return err
}

func (s *storage) ListReminders(ctx context.Context) ([]*models.Reminder, error) {
	return s.repo.ListReminders(ctx)
}
