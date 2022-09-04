package consumer

import (
	"context"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/internal/storage"
)

func checkParams(params map[string]string, needed ...string) error {
	for _, field := range needed {
		if _, ok := params[field]; !ok {
			return errors.Errorf("invalid argument: filed %s doesn't set", field)
		}
	}
	return nil
}

func update(ctx context.Context, storage storage.RemindersStorage, params map[string]string) error {
	if err := checkParams(params, "id", "text"); err != nil {
		return err
	}
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		return errors.Errorf("parse id error: %s", err.Error())
	}
	err = storage.UpdateReminder(ctx, id, params["text"])
	if err != nil {
		return err
	}
	return nil
}

func remove(ctx context.Context, storage storage.RemindersStorage, params map[string]string) error {
	if err := checkParams(params, "id"); err != nil {
		return err
	}
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		return errors.Errorf("parse id error: %s", err.Error())
	}
	err = storage.RemoveReminder(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func create(ctx context.Context, storage storage.RemindersStorage, params map[string]string) (uint64, error) {
	if err := checkParams(params, "date", "text"); err != nil {
		return 0, err
	}
	date, err := time.Parse(time.RFC3339, params["date"])
	if err != nil {
		return 0, errors.Errorf("parse date error: %s", err.Error())
	}
	rem, err := storage.CreateReminder(ctx, date, params["text"])
	if err != nil {
		return 0, err
	}
	return rem.ID, nil
}
