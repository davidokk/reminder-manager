package repository

import (
	"context"
	"time"

	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/internal/models"
)

// Interface provides interface for work with db
type Interface interface {
	CreateReminder(ctx context.Context, date time.Time, text string) (*models.Reminder, error)
	GetReminder(ctx context.Context, id uint64) (*models.Reminder, error)
	UpdateReminder(ctx context.Context, id uint64, text string) error
	RemoveReminder(ctx context.Context, id uint64) error
	ListReminders(ctx context.Context) ([]*models.Reminder, error)

	Close()
}
