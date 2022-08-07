package postgres

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"gitlab.ozon.dev/davidokk/reminder-manager/internal/models"
)

// CreateReminder adds a new Reminder into storage
func (r *Repository) CreateReminder(ctx context.Context, date time.Time, text string) (*models.Reminder, error) {
	query, args, err := squirrel.Insert("reminders").
		Columns("date, text").
		Values(date.Format("2006-01-02"), text).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return nil, err
	}
	row := r.pool.QueryRow(ctx, query, args...)

	rem := &models.Reminder{
		Date: date,
		Text: text,
	}

	if err := row.Scan(&rem.ID); err != nil {
		return nil, err
	}
	return rem, nil
}

// ListReminders returns all reminders as slice
func (r *Repository) ListReminders(ctx context.Context) ([]*models.Reminder, error) {
	query, args, err := squirrel.Select("*").
		From("reminders").
		OrderBy("date").
		ToSql()
	if err != nil {
		return nil, err
	}
	var rem []*models.Reminder
	if err := pgxscan.Select(ctx, r.pool, &rem, query, args...); err != nil {
		return nil, err
	}
	return rem, nil
}

// GetReminder returns a reminder with given id
func (r *Repository) GetReminder(ctx context.Context, id uint64) (*models.Reminder, error) {
	query, args, err := squirrel.Select("id, date, text").
		From("reminders").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	rem := &models.Reminder{}
	row, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	if err := pgxscan.ScanOne(rem, row); err != nil {
		return nil, err
	}
	return rem, nil
}

// UpdateReminder allows to change the text of Reminder with given ID
func (r *Repository) UpdateReminder(ctx context.Context, id uint64, text string) error {
	query, args, err := squirrel.Update("reminders").
		Set("text", text).
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

// RemoveReminder removes Reminder with given ID
func (r *Repository) RemoveReminder(ctx context.Context, id uint64) error {
	query, args, err := squirrel.Delete("reminders").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}
	if _, err := r.pool.Exec(ctx, query, args...); err != nil {
		return err
	}
	return nil
}
