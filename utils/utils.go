package utils

import (
	"reflect"
	"time"
)

// UpToDay rounds the date down to the day
func UpToDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
}

// Insert inserts el into a slice, new element will have index idx
func Insert[T any](data []T, el T, idx int) []T {
	data = append(data, el)
	swap := reflect.Swapper(data)
	for i := idx; i < len(data); i++ {
		swap(i, len(data)-1)
	}
	return data
}

// Remove removes given index from slice
func Remove[T any](data []T, idx int) []T {
	copy(data[idx:], data[idx+1:])
	return data[:len(data)-1]
}

// Clone returns copy of slice
func Clone[T any](data []T) []T {
	res := make([]T, 0, len(data))
	res = append(res, data...)
	return res
}
