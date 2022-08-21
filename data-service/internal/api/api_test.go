package api

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/internal/models"
	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/pkg/api"
	"gitlab.ozon.dev/davidokk/reminder-manager/utils"
)

var ID uint64 = 1
var text = "test"
var date = utils.UpToDay(time.Now())

func TestReminderList(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().ListReminders(gomock.Any()).Return([]*models.Reminder{
			{
				Text: text,
				Date: date,
				ID:   ID,
			},
		}, nil)

		resp, err := f.service.ReminderList(f.ctx, &api.ReminderListRequest{})

		require.NoError(t, err)
		assert.Equal(t, resp, &api.ReminderListResponse{
			Reminders: []*api.Reminder{
				{
					Id:   ID,
					Text: text,
					Date: utils.TimeToTimestamp(date),
				},
			},
		})
	})

	t.Run("error", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().ListReminders(gomock.Any()).Return(nil, errors.New(""))

		resp, err := f.service.ReminderList(f.ctx, &api.ReminderListRequest{})

		require.Error(t, err)
		assert.Equal(t, resp, (*api.ReminderListResponse)(nil))
	})
}

func TestReminderGet(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().GetReminder(gomock.Any(), ID).Return(&models.Reminder{
			ID:   ID,
			Text: text,
			Date: date,
		}, nil)

		resp, err := f.service.ReminderGet(f.ctx, &api.ReminderGetRequest{
			Id: ID,
		})

		require.NoError(t, err)
		assert.Equal(t, resp, &api.ReminderGetResponse{
			Reminder: &api.Reminder{
				Id:   ID,
				Text: text,
				Date: utils.TimeToTimestamp(date),
			},
		})
	})

	t.Run("error", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().GetReminder(gomock.Any(), ID).Return(nil, errors.New(""))

		resp, err := f.service.ReminderGet(f.ctx, &api.ReminderGetRequest{
			Id: ID,
		})

		require.Error(t, err)
		assert.Equal(t, resp, (*api.ReminderGetResponse)(nil))
	})
}

func TestReminderCreate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().CreateReminder(gomock.Any(), date, text).Return(&models.Reminder{
			ID:   ID,
			Text: text,
			Date: date,
		}, nil)

		resp, err := f.service.ReminderCreate(f.ctx, &api.ReminderCreateRequest{
			Text: text,
			Date: utils.TimeToTimestamp(date),
		})

		require.NoError(t, err)
		assert.Equal(t, resp, &api.ReminderCreateResponse{
			Id: ID,
		})
	})

	t.Run("error", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().CreateReminder(gomock.Any(), date, text).Return(nil, errors.New("some error"))

		resp, err := f.service.ReminderCreate(f.ctx, &api.ReminderCreateRequest{
			Text: text,
			Date: utils.TimeToTimestamp(date),
		})

		require.Error(t, err)
		assert.Equal(t, resp, (*api.ReminderCreateResponse)(nil))
	})
}

func TestReminderUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().UpdateReminder(gomock.Any(), ID, text).Return(nil)

		resp, err := f.service.ReminderUpdate(f.ctx, &api.ReminderUpdateRequest{
			Text: text,
			Id:   ID,
		})

		require.NoError(t, err)
		assert.Equal(t, resp, &api.ReminderUpdateResponse{})
	})

	t.Run("error", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().UpdateReminder(gomock.Any(), ID, text).Return(errors.New(""))

		resp, err := f.service.ReminderUpdate(f.ctx, &api.ReminderUpdateRequest{
			Text: text,
			Id:   ID,
		})

		require.Error(t, err)
		assert.Equal(t, resp, (*api.ReminderUpdateResponse)(nil))
	})
}

func TestReminderRemove(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().RemoveReminder(gomock.Any(), ID).Return(nil)

		resp, err := f.service.ReminderRemove(f.ctx, &api.ReminderRemoveRequest{
			Id: ID,
		})

		require.NoError(t, err)
		assert.Equal(t, resp, &api.ReminderRemoveResponse{})
	})

	t.Run("error", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().RemoveReminder(gomock.Any(), ID).Return(errors.New(""))

		resp, err := f.service.ReminderRemove(f.ctx, &api.ReminderRemoveRequest{
			Id: ID,
		})

		require.Error(t, err)
		assert.Equal(t, resp, (*api.ReminderRemoveResponse)(nil))
	})
}
