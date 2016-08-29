package tools

import (
	"time"
	"strconv"
)

type Event struct {
	Id int64
	Reference string
	CreatedAt EventTime
	PromoterId int64
	PromoterName string
	Participants []User
	Spending []Spending
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

type SpendingType struct {
	Id int64
	Reference string
}
type Spending struct {
	Id int64
	TypeId int64
	TypeReference string
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
	Date string
	Amount float64
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
	s.SpendingAt.TimeStruct = time.Unix(0, s.SpendingAt.TimeStamp*int64(time.Second))
	s.SpendingAt.TimeString = s.SpendingAt.TimeStruct.Format("2006-01-02 15:04:05")
	s.CreatedAt.TimeStruct = time.Unix(0, s.CreatedAt.TimeStamp*int64(time.Second))
	s.CreatedAt.TimeString = s.CreatedAt.TimeStruct.Format("2006-01-02 15:04:05")
	payer,_ := UsersId[s.PayerId]
	s.PayerName = payer.Login
	for i, sf := range s.For {
		debtor,_ := UsersId[sf.DebtorId]
		sf.DebtorName = debtor.Login
		s.For[i]=sf
	}
	
	s.Rows.Date = s.SpendingAt.TimeString
	s.Rows.Amount = s.Amount
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