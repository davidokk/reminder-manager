//go:build integration
// +build integration

package tests

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/pkg/api"
	"gitlab.ozon.dev/davidokk/reminder-manager/utils"
)

var ID uint64 = 1
var text = "test"
var date = utils.UpToDay(time.Now())

func TestReminderCreate(t *testing.T) {
	Db.SetUp(t)
	defer Db.TearDown()

	resp, err := Client.ReminderCreate(context.Background(), &api.ReminderCreateRequest{
		Date: utils.TimeToTimestamp(date),
		Text: text,
	})

	require.NoError(t, err)
	assert.Equal(t, resp.GetId(), ID)
}

func TestReminderGet(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		_, _ = Client.ReminderCreate(context.Background(), &api.ReminderCreateRequest{
			Date: utils.TimeToTimestamp(date),
			Text: text,
		})

		resp, err := Client.ReminderGet(context.Background(), &api.ReminderGetRequest{
			Id: ID,
		})

		require.NoError(t, err)
		assert.Equal(t, resp.GetReminder().GetText(), text)
		assert.Equal(t, resp.GetReminder().GetDate(), utils.TimeToTimestamp(date))
		assert.Equal(t, resp.GetReminder().GetId(), ID)
	})

	t.Run("error", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		_, err := Client.ReminderGet(context.Background(), &api.ReminderGetRequest{
			Id: ID,
		})

		require.Error(t, err)
	})
}

func TestReminderUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		_, _ = Client.ReminderCreate(context.Background(), &api.ReminderCreateRequest{
			Date: utils.TimeToTimestamp(date),
			Text: text,
		})

		_, err := Client.ReminderUpdate(context.Background(), &api.ReminderUpdateRequest{
			Id:   ID,
			Text: "new_test",
		})

		getResp, getErr := Client.ReminderGet(context.Background(), &api.ReminderGetRequest{
			Id: ID,
		})

		require.NoError(t, err)
		require.NoError(t, getErr)
		assert.Equal(t, getResp.GetReminder().GetText(), "new_test")
	})

	t.Run("error", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		_, err := Client.ReminderUpdate(context.Background(), &api.ReminderUpdateRequest{
			Id:   ID,
			Text: "new_test",
		})

		require.Error(t, err)
	})
}

func TestReminderRemove(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		_, _ = Client.ReminderCreate(context.Background(), &api.ReminderCreateRequest{
			Date: utils.TimeToTimestamp(date),
			Text: text,
		})

		_, err := Client.ReminderRemove(context.Background(), &api.ReminderRemoveRequest{
			Id: ID,
		})

		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		_, err := Client.ReminderRemove(context.Background(), &api.ReminderRemoveRequest{
			Id: ID,
		})

		require.Error(t, err)
	})
}

func TestReminderList(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		resp, err := Client.ReminderList(context.Background(), &api.ReminderListRequest{})

		require.NoError(t, err)
		assert.Equal(t, len(resp.GetReminders()), 0)
	})

	t.Run("not empty", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		_, _ = Client.ReminderCreate(context.Background(), &api.ReminderCreateRequest{
			Date: utils.TimeToTimestamp(date),
			Text: text,
		})

		resp, err := Client.ReminderList(context.Background(), &api.ReminderListRequest{})

		require.NoError(t, err)
		require.Equal(t, len(resp.GetReminders()), 1)
		assert.Equal(t, resp.GetReminders()[0].GetId(), ID)
		assert.Equal(t, resp.GetReminders()[0].GetDate(), utils.TimeToTimestamp(date))
		assert.Equal(t, resp.GetReminders()[0].GetText(), text)

	})
}
