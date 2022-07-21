package utils

import (
	"reflect"
	"time"
)

func UpToDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
}

func Insert[T any](data []T, el T, idx int) []T {
	data = append(data, el)
	swap := reflect.Swapper(data)
	for i := idx; i < len(data); i++ {
		swap(i, len(data)-1)
	}
	return data
}

func Remove[T any](data []T, idx int) []T {
	copy(data[idx:], data[idx+1:])
	return data[:len(data)-1]
}

func Clone[T any](data []T) []T {
	res := make([]T, 0, len(data))
	res = append(res, data...)
	return res
}
