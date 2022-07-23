package storage

import (
	"fmt"
	"time"
)

var lastID uint64

// Reminder contains reminder information
type Reminder struct {
	ID   uint64
	What string
	Date time.Time
}

// NewReminder creates a new Reminder with unique ID
func NewReminder(what string, date time.Time) *Reminder {
	lastID++
	return &Reminder{
		ID:   lastID,
		What: what,
		Date: date,
	}
}

// ToString converts Reminder to string
func (rem *Reminder) ToString() string {
	date := rem.Date.Format("Mon, 2 Jan")
	return fmt.Sprintf("[%d]: %s - %s", rem.ID, date, rem.What)
}
