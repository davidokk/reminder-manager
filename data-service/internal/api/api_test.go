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

var expectedID uint64 = 1
var expectedText = "test"
var expectedDate = utils.UpToDay(time.Now())

func TestReminderList(t *testing.T) {
	t.Run("returns reminders list", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().ListReminders(gomock.Any()).Return([]*models.Reminder{
			{
				Text: expectedText,
				Date: expectedDate,
				ID:   expectedID,
			},
		}, nil)

		resp, err := f.service.ReminderList(f.ctx, &api.ReminderListRequest{})

		require.NoError(t, err)
		assert.Equal(t, resp, &api.ReminderListResponse{
			Reminders: []*api.Reminder{
				{
					Id:   expectedID,
					Text: expectedText,
					Date: utils.TimeToTimestamp(expectedDate),
				},
			},
		})
	})

	t.Run("handle error from storage", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().ListReminders(gomock.Any()).Return(nil, errors.New(""))

		resp, err := f.service.ReminderList(f.ctx, &api.ReminderListRequest{})

		require.Error(t, err)
		assert.EqualError(t, err, "rpc error: code = Internal desc = ")
		assert.Nil(t, resp)
	})
}

func TestReminderGet(t *testing.T) {
	t.Run("get existing element", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().GetReminder(gomock.Any(), expectedID).Return(&models.Reminder{
			ID:   expectedID,
			Text: expectedText,
			Date: expectedDate,
		}, nil)

		resp, err := f.service.ReminderGet(f.ctx, &api.ReminderGetRequest{
			Id: expectedID,
		})

		require.NoError(t, err)
		assert.Equal(t, resp, &api.ReminderGetResponse{
			Reminder: &api.Reminder{
				Id:   expectedID,
				Text: expectedText,
				Date: utils.TimeToTimestamp(expectedDate),
			},
		})
	})

	t.Run("get non existing element", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().GetReminder(gomock.Any(), expectedID).Return(nil, errors.New(""))

		resp, err := f.service.ReminderGet(f.ctx, &api.ReminderGetRequest{
			Id: expectedID,
		})

		require.Error(t, err)
		assert.EqualError(t, err, "rpc error: code = Internal desc = ")
		assert.Nil(t, resp)
	})
}

func TestReminderCreate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().CreateReminder(gomock.Any(), expectedDate, expectedText).Return(&models.Reminder{
			ID:   expectedID,
			Text: expectedText,
			Date: expectedDate,
		}, nil)

		resp, err := f.service.ReminderCreate(f.ctx, &api.ReminderCreateRequest{
			Text: expectedText,
			Date: utils.TimeToTimestamp(expectedDate),
		})

		require.NoError(t, err)
		assert.Equal(t, resp, &api.ReminderCreateResponse{
			Id: expectedID,
		})
	})

	t.Run("handle error from storage", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().CreateReminder(gomock.Any(), expectedDate, expectedText).Return(nil, errors.New(""))

		resp, err := f.service.ReminderCreate(f.ctx, &api.ReminderCreateRequest{
			Text: expectedText,
			Date: utils.TimeToTimestamp(expectedDate),
		})

		require.Error(t, err)
		assert.EqualError(t, err, "rpc error: code = Internal desc = ")
		assert.Nil(t, resp)
	})
}

func TestReminderUpdate(t *testing.T) {
	t.Run("update existing element", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().UpdateReminder(gomock.Any(), expectedID, expectedText).Return(nil)

		resp, err := f.service.ReminderUpdate(f.ctx, &api.ReminderUpdateRequest{
			Text: expectedText,
			Id:   expectedID,
		})

		require.NoError(t, err)
		assert.Equal(t, resp, &api.ReminderUpdateResponse{})
	})

	t.Run("update non existing element", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().UpdateReminder(gomock.Any(), expectedID, expectedText).Return(errors.New(""))

		resp, err := f.service.ReminderUpdate(f.ctx, &api.ReminderUpdateRequest{
			Text: expectedText,
			Id:   expectedID,
		})

		require.Error(t, err)
		assert.EqualError(t, err, "rpc error: code = Internal desc = ")
		assert.Nil(t, resp)
	})
}

func TestReminderRemove(t *testing.T) {
	t.Run("remove existing element", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().RemoveReminder(gomock.Any(), expectedID).Return(nil)

		resp, err := f.service.ReminderRemove(f.ctx, &api.ReminderRemoveRequest{
			Id: expectedID,
		})

		require.NoError(t, err)
		assert.Equal(t, resp, &api.ReminderRemoveResponse{})
	})

	t.Run("remove non existing element", func(t *testing.T) {
		f := setUp(t)

		f.storage.EXPECT().RemoveReminder(gomock.Any(), expectedID).Return(errors.New(""))

		resp, err := f.service.ReminderRemove(f.ctx, &api.ReminderRemoveRequest{
			Id: expectedID,
		})

		require.Error(t, err)
		assert.EqualError(t, err, "rpc error: code = Internal desc = ")
		assert.Nil(t, resp)
	})
}
