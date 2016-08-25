package db

import (
	"strings"
	"../tools"
	"database/sql"
)

func LoadEventsForThisUser(user *tools.User, evs *[]tools.Event) error {

	var err error
	var rows *sql.Rows
	var evsId []string
	
	rows, err = DbConnect.Query(`select event_id from participants where user_id = ?`, user.Id)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {

		var id string

		err = rows.Scan(&id)
		if err != nil {
			return err
		}
		evsId = append(evsId, id)
	}
	
	rows, err = DbConnect.Query(`select id, reference, created_at, promoter_id from events where id IN (`+strings.Join(evsId, ",")+`) order by id desc`)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {

		var ev tools.Event

		err = rows.Scan(&ev.Id, &ev.Reference, &ev.CreatedAt.TimeStamp, &ev.PromoterId)
		if err != nil {
			return err
		}
		ev.Feed()
		*evs = append(*evs, ev)
	}
	
	return nil
}

func LoadTypeSpending(ts *[]string) error {
	
	var err error
	var rows *sql.Rows
	
	// Load event
	
	rows, err = DbConnect.Query(`
				select 
					reference 
				from 
					spending_type`)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var reference string
		err = rows.Scan(&reference)
		if err != nil {
			return err
		}
		*ts = append(*ts, reference)
	}
	
	return nil
}

func CheckEventForThisUser(user *tools.User, ev *tools.Event) (bool,error) {

	var ok bool
	
	err := DbConnect.QueryRow(`
				select 
					count(0) 
				from 
					participants 
				where 
					user_id = ? 
				and 
					event_id = ?`, user.Id, ev.Id).Scan(&ok)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func LoadThisEvent(ev *tools.Event) error {
	
	var err error
	var rows *sql.Rows
	
	// Load event
	
	rows, err = DbConnect.Query(`
				select 
					id, reference, created_at, promoter_id 
				from 
					events
				where 
					id = ?`, ev.Id)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&ev.Id, &ev.Reference, &ev.CreatedAt.TimeStamp, &ev.PromoterId)
		if err != nil {
			return err
		}
		ev.Feed()
	}
	
	// Load Participants
	
	rows, err = DbConnect.Query(`
				select 
					user_id
				from 
					participants
				where 
					event_id = ?`, ev.Id)
	if err != nil {
		return err
	}
	defer rows.Close()
	
	for rows.Next() {
		var u tools.User
		err = rows.Scan(&u.Id)
		if err != nil {
			return err
		}
		u, _ = tools.UsersId[u.Id]
		ev.Participants = append(ev.Participants, u)
	}
	
	// Load spendings
	
	rows, err = DbConnect.Query(`
				select 
					s.id, s.type_id, st.reference, s.description, s.amount, s.spending_at, s.created_at, s.payer_id 
				from 
					spending as s
				join
					spending_type as st
				on
					st.id = s.type_id
				where 
					s.event_id = ?
				order by 
					s.spending_at desc`, ev.Id)
	if err != nil {
		return err
	}
	defer rows.Close()
	
	for rows.Next() {
		var s tools.Spending
		err = rows.Scan(&s.Id, &s.TypeId, &s.TypeReference, &s.Description, &s.Amount, &s.SpendingAt.TimeStamp, &s.CreatedAt.TimeStamp, &s.PayerId)
		if err != nil {
			return err
		}
		ev.Spending = append(ev.Spending, s)
	}
	
// 	tools.Info(ev.Id)
	
	for i, s := range ev.Spending {
		
		rows, err = DbConnect.Query(`
				select 
					debtor_id, debt 
				from 
					spending_for
				where 
					spending_id = ?`, s.Id)
		if err != nil {
			return err
		}
		defer rows.Close()
		
		for rows.Next() {
			var sf tools.SpendingFor
			err = rows.Scan(&sf.DebtorId, &sf.Debt)
			if err != nil {
				return err
			}
			s.For = append(s.For, sf)
		}
		s.Feed(ev.Participants)
		ev.Spending[i]=s
		
		tools.Info(ev.Spending[i])
	}
	
	return nil
}