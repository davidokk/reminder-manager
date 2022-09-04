//go:build integration
// +build integration

package tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/pkg/api"
	"gitlab.ozon.dev/davidokk/reminder-manager/utils"
)

var expectedID uint64 = 1
var expectedText = "test"
var expectedDate = utils.UpToDay(time.Now())

func TestReminderCreate(t *testing.T) {
	Db.SetUp(t)
	defer Db.TearDown()

	resp, err := Client.ReminderCreate(context.Background(), &api.ReminderCreateRequest{
		Date: utils.TimeToTimestamp(expectedDate),
		Text: expectedText,
	})

	require.NoError(t, err)
	assert.Equal(t, resp.GetId(), expectedID)
}

func TestReminderGet(t *testing.T) {
	t.Run("get existing element", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		_, _ = Client.ReminderCreate(context.Background(), &api.ReminderCreateRequest{
			Date: utils.TimeToTimestamp(expectedDate),
			Text: expectedText,
		})

		resp, err := Client.ReminderGet(context.Background(), &api.ReminderGetRequest{
			Id: expectedID,
		})

		require.NoError(t, err)
		assert.Equal(t, resp.GetReminder().GetText(), expectedText)
		assert.Equal(t, resp.GetReminder().GetDate(), utils.TimeToTimestamp(expectedDate))
		assert.Equal(t, resp.GetReminder().GetId(), expectedID)
	})

	t.Run("get non existing element", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		_, err := Client.ReminderGet(context.Background(), &api.ReminderGetRequest{
			Id: expectedID,
		})

		require.Error(t, err)
		assert.EqualError(t, err, "rpc error: code = Internal desc = no rows in result set")
	})
}

func TestReminderUpdate(t *testing.T) {
	t.Run("update existing element", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		_, _ = Client.ReminderCreate(context.Background(), &api.ReminderCreateRequest{
			Date: utils.TimeToTimestamp(expectedDate),
			Text: expectedText,
		})

		_, err := Client.ReminderUpdate(context.Background(), &api.ReminderUpdateRequest{
			Id:   expectedID,
			Text: "new_test",
		})

		getResp, getErr := Client.ReminderGet(context.Background(), &api.ReminderGetRequest{
			Id: expectedID,
		})

		require.NoError(t, err)
		require.NoError(t, getErr)
		assert.Equal(t, getResp.GetReminder().GetText(), "new_test")
	})

	t.Run("update non existing element", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		_, err := Client.ReminderUpdate(context.Background(), &api.ReminderUpdateRequest{
			Id:   expectedID,
			Text: "new_test",
		})

		require.Error(t, err)
		assert.EqualError(t, err, "rpc error: code = Internal desc = can't find given id")
	})
}

func TestReminderRemove(t *testing.T) {
	t.Run("remove existing element", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		_, _ = Client.ReminderCreate(context.Background(), &api.ReminderCreateRequest{
			Date: utils.TimeToTimestamp(expectedDate),
			Text: expectedText,
		})

		_, err := Client.ReminderRemove(context.Background(), &api.ReminderRemoveRequest{
			Id: expectedID,
		})

		require.NoError(t, err)
	})

	t.Run("remove non existent element", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		_, err := Client.ReminderRemove(context.Background(), &api.ReminderRemoveRequest{
			Id: expectedID,
		})

		require.Error(t, err)
		assert.EqualError(t, err, "rpc error: code = Internal desc = can't find given id")
	})
}

func TestReminderList(t *testing.T) {
	t.Run("returns empty list", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		resp, err := Client.ReminderList(context.Background(), &api.ReminderListRequest{})

		require.NoError(t, err)
		assert.Equal(t, len(resp.GetReminders()), 0)
	})

	t.Run("returns not empty list", func(t *testing.T) {
		Db.SetUp(t)
		defer Db.TearDown()

		_, _ = Client.ReminderCreate(context.Background(), &api.ReminderCreateRequest{
			Date: utils.TimeToTimestamp(expectedDate),
			Text: expectedText,
		})

		resp, err := Client.ReminderList(context.Background(), &api.ReminderListRequest{})

		require.NoError(t, err)
		require.Equal(t, len(resp.GetReminders()), 1)
		assert.Equal(t, resp.GetReminders()[0].GetId(), expectedID)
		assert.Equal(t, resp.GetReminders()[0].GetDate(), utils.TimeToTimestamp(expectedDate))
		assert.Equal(t, resp.GetReminders()[0].GetText(), expectedText)

	})
}
