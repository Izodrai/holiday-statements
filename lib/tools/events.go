package tools

import (
	"time"
	"strconv"
)

type Debts struct {
	DebtorId int64
	DebtorName string
	TotalSpending float64
	STotalSpending string
	Debts map[int64]float64
	SDebts []DebsToPrint
}

type DebsToPrint struct {
	Name string 
	Amount string
}

type Event struct {
	Id int64
	Reference string
	CreatedAt EventTime
	PromoterId int64
	PromoterName string
	Participants []User
	Spending []Spending
}

func (e *Event) Feed() {
	e.CreatedAt.TimeStruct = time.Unix(e.CreatedAt.TimeStamp,0)
	e.CreatedAt.TimeString = e.CreatedAt.TimeStruct.Format("2006-01-02")
	promoter,_ := UsersId[e.PromoterId]
	e.PromoterName = promoter.Login
}

type EventTime struct {
	TimeStruct time.Time
	TimeString string
	TimeStamp int64
}

func (et *EventTime) FeedEventTimeFromStruct() {
	et.TimeString = et.TimeStruct.Format("2006-01-02")
	et.TimeStamp = et.TimeStruct.Unix()
}

type SpendingType struct {
	Id int64
	Reference string
}
type Spending struct {
	Id int64
	TypeId int64
	TypeReference string
	EventId int64
	Description string
	Amount float64
	SpendingAt EventTime
	CreatedAt EventTime
	PayerId int64
	PayerName string
	For []SpendingFor
	Rows RowToDisplay
}

type RowToDisplay struct {
	Id int64
	Date string
	Amount string
	PayerName string
	Debts []string
	TypeSpending string
	Description string
}

type SpendingFor struct {
	DebtorId int64
	DebtorName string
	Debt float64
}

func (s *Spending) Feed(participants []User) {
	s.SpendingAt.TimeStruct = time.Unix(s.SpendingAt.TimeStamp,0)
	s.SpendingAt.TimeString = s.SpendingAt.TimeStruct.Format("2006-01-02")
	s.CreatedAt.TimeStruct = time.Unix(s.CreatedAt.TimeStamp,0)
	s.CreatedAt.TimeString = s.CreatedAt.TimeStruct.Format("2006-01-02")
	payer,_ := UsersId[s.PayerId]
	s.PayerName = payer.Login
	for i, sf := range s.For {
		debtor,_ := UsersId[sf.DebtorId]
		sf.DebtorName = debtor.Login
		s.For[i]=sf
	}
	
	s.Rows.Id = s.Id
	s.Rows.Date = s.SpendingAt.TimeString
	s.Rows.Amount = strconv.FormatFloat(s.Amount,'f',2,64) 
	s.Rows.PayerName = s.PayerName
	
	for i, user := range participants {
		s.Rows.Debts = append(s.Rows.Debts, "")
		for _, sf := range s.For {
			if sf.DebtorName == user.Login {
				s.Rows.Debts[i]=strconv.FormatFloat(sf.Debt,'f',2,64)
				break
			}
		}
	}
	
	s.Rows.TypeSpending = s.TypeReference
	s.Rows.Description = s.Description
}