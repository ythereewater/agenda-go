package model

import(
	"time"
	"fmt"
)

type Date struct {
	Time time.Time
}


func (date *Date) Init(st string) {
	t, err := time.Parse(time.RFC3339, st + ":00Z")
	if err == nil {
		date.Time = t
	} else {
		fmt.Println(err);
	}
}

func (date Date) GetYear() int {
	return date.Time.Year()
}


func (date Date) GetMonth() string {
	return date.Time.Month().String()
}


func (date Date) GetDay() int {
	return date.Time.Day()
}

func (date Date) GetHour() int {
	return date.Time.Hour()
}

func (date Date) GetMinute() int {
	return date.Time.Minute()
}


func (date Date)IsEqual(other Date) bool {
	return date.Time.Equal(other.Time)
}

func (date Date)IsAfter(other Date) bool {
	return date.Time.After(other.Time)
}


func (date Date)DateToString() string {
	st := date.Time.String()[0 : 16]
	return st
}


func StringToDate(st string) (Date) {
	var date Date
	t, err := time.Parse(time.RFC3339, st + ":00Z")
	if err == nil {
		date.Time = t
	} else {
		date.Init("2000-01-01T00:00")
	}
	return date
}
