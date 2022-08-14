package models

import (
	"fmt"
	"time"
)

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
