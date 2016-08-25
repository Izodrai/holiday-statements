package tools

import (
	"time"
)

type Event struct {
	Id int64
	Reference string
	CreatedAt EventTime
	PromoterId int64
	PromoterName string
}

type EventTime struct {
	TimeStruct time.Time
	TimeString string
	TimeStamp int64
}

func (e *Event) Feed() {
	e.CreatedAt.TimeStruct = time.Unix(0, e.CreatedAt.TimeStamp*int64(time.Second))
	e.CreatedAt.TimeString = e.CreatedAt.TimeStruct.Format("2006-01-02 15:04:05")
	promoter,_ := UsersId[e.PromoterId]
	e.PromoterName = promoter.Login
}