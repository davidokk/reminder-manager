package models

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/davidokk/reminder-manager/utils"
)

const maxTextLength = 255

// Reminder contains reminder information
type Reminder struct {
	ID   uint64
	Text string
	Date time.Time
}

// String converts Reminder to string
func (rem *Reminder) String() string {
	date := rem.Date.Format("Mon, 2 Jan")
	return fmt.Sprintf("[%d]: %s - %s", rem.ID, date, rem.Text)
}

// Valid checks if all fields are correct
func (rem *Reminder) Valid() error {
	if rem.Date.Before(utils.UpToDay(time.Now())) {
		return errors.New("past date")
	}
	if len(rem.Text) > maxTextLength {
		return errors.New(fmt.Sprintf("text lenght should not exceed %d", maxTextLength))
	}
	return nil
}
