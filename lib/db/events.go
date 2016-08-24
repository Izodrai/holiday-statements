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