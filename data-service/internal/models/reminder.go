package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// Reminder contains reminder information
type Reminder struct {
	ID   uint64    `db:"id"`
	Text string    `db:"text"`
	Date time.Time `db:"date"`
}

// String converts Reminder to string
func (rem *Reminder) String() string {
	date := rem.Date.Format("Mon, 2 Jan")
	return fmt.Sprintf("[%d]: %s - %s", rem.ID, date, rem.Text)
}

// MarshalBinary converts to []byte
func (rem *Reminder) MarshalBinary() ([]byte, error) {
	return json.Marshal(rem)
}

// UnmarshalBinary reads from []byte
func (rem *Reminder) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, rem)
}
