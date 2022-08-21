package postgres

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/jackc/pgconn"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/internal/models"
)

var expectedText = "test"

var layout = "2006-01-02"
var expectedDate = time.Date(2022, 8, 21, 0, 0, 0, 0, time.UTC)

func TestCreateReminder(t *testing.T) {
	f := setUp()
	defer f.tearDown(t)

	query := regexp.QuoteMeta(`INSERT INTO reminders (date, text) VALUES ($1,$2) RETURNING id`)
	f.pool.ExpectQuery(query).
		WithArgs(expectedDate.Format(layout), expectedText).
		WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow(uint64(1)))

	rem, err := f.repository.CreateReminder(context.Background(), expectedDate, expectedText)

	require.NoError(t, err)
	assert.Equal(t, *rem, models.Reminder{
		ID:   1,
		Text: expectedText,
		Date: expectedDate,
	})
}

func TestListReminders(t *testing.T) {
	f := setUp()
	defer f.tearDown(t)

	query := regexp.QuoteMeta(`SELECT * FROM reminders ORDER BY date`)
	f.pool.ExpectQuery(query).
		WillReturnRows(pgxmock.NewRows([]string{"id", "date", "text"}).
			AddRow(uint64(1), expectedDate, expectedText).
			AddRow(uint64(2), expectedDate, expectedText))

	rem, err := f.repository.ListReminders(context.Background())

	require.NoError(t, err)
	require.Equal(t, len(rem), 2)
	assert.Equal(t, *rem[0], models.Reminder{
		ID:   1,
		Text: expectedText,
		Date: expectedDate,
	})
	assert.Equal(t, *rem[1], models.Reminder{
		ID:   2,
		Text: expectedText,
		Date: expectedDate,
	})
}

func TestGetReminder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		f := setUp()
		defer f.tearDown(t)

		var ID uint64 = 4

		query := regexp.QuoteMeta(`SELECT id, date, text FROM reminders WHERE id = $1`)
		f.pool.ExpectQuery(query).
			WillReturnRows(pgxmock.NewRows([]string{"id", "date", "text"}).AddRow(ID, expectedDate, expectedText))

		rem, err := f.repository.GetReminder(context.Background(), ID)

		require.NoError(t, err)
		assert.Equal(t, *rem, models.Reminder{
			ID:   ID,
			Text: expectedText,
			Date: expectedDate,
		})
	})

	t.Run("error", func(t *testing.T) {
		f := setUp()
		defer f.tearDown(t)

		var ID uint64 = 4

		query := regexp.QuoteMeta(`SELECT id, date, text FROM reminders WHERE id = $1`)
		f.pool.ExpectQuery(query).WillReturnRows(pgxmock.NewRows([]string{"id", "date", "text"}))

		_, err := f.repository.GetReminder(context.Background(), ID)

		require.Error(t, err)
	})
}

func TestUpdateReminder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		f := setUp()
		defer f.tearDown(t)

		var ID uint64 = 4

		getQuery := regexp.QuoteMeta(`SELECT id, date, text FROM reminders WHERE id = $1`)
		f.pool.ExpectQuery(getQuery).WithArgs(ID).
			WillReturnRows(pgxmock.NewRows([]string{"id", "date", "text"}).AddRow(ID, expectedDate, expectedText))

		updateQuery := regexp.QuoteMeta(`UPDATE reminders SET text = $1 WHERE id = $2`)
		f.pool.ExpectExec(updateQuery).WithArgs(expectedText, ID).WillReturnResult(pgconn.CommandTag{})

		err := f.repository.UpdateReminder(context.Background(), ID, expectedText)

		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		f := setUp()
		defer f.tearDown(t)

		var ID uint64 = 4

		getQuery := regexp.QuoteMeta(`SELECT id, date, text FROM reminders WHERE id = $1`)
		f.pool.ExpectQuery(getQuery).WithArgs(ID).
			WillReturnRows(pgxmock.NewRows([]string{"id", "date", "text"}))

		err := f.repository.UpdateReminder(context.Background(), ID, expectedText)

		require.Error(t, err)
	})
}

func TestRemoveReminder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		f := setUp()
		defer f.tearDown(t)

		var ID uint64 = 4

		getQuery := regexp.QuoteMeta(`SELECT id, date, text FROM reminders WHERE id = $1`)
		f.pool.ExpectQuery(getQuery).WithArgs(ID).
			WillReturnRows(pgxmock.NewRows([]string{"id", "date", "text"}).AddRow(ID, expectedDate, expectedText))

		updateQuery := regexp.QuoteMeta(`DELETE FROM reminders WHERE id = $1`)
		f.pool.ExpectExec(updateQuery).WithArgs(ID).WillReturnResult(pgconn.CommandTag{})

		err := f.repository.RemoveReminder(context.Background(), ID)

		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		f := setUp()
		defer f.tearDown(t)

		var ID uint64 = 4

		getQuery := regexp.QuoteMeta(`SELECT id, date, text FROM reminders WHERE id = $1`)
		f.pool.ExpectQuery(getQuery).WithArgs(ID).
			WillReturnRows(pgxmock.NewRows([]string{"id", "date", "text"}))

		err := f.repository.RemoveReminder(context.Background(), ID)

		require.Error(t, err)
	})
}
