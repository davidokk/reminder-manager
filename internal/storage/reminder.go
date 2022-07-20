package storage

import (
	"fmt"
	"time"
)

var lastId uint64 = 0

type Reminder struct {
	Id   uint64
	What string
	Date time.Time
}

func NewReminder(what string, date time.Time) *Reminder {
	lastId++
	return &Reminder{
		Id:   lastId,
		What: what,
		Date: date,
	}
}

func (rem *Reminder) ToString() string {
	date := rem.Date.Format("Mon, 2 Jan")
	return fmt.Sprintf("[%d]: %s - %s", rem.Id, date, rem.What)
}
