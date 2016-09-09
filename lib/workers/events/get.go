package events

import (
	tmpl "../../templates"
	"../../tools"
	"github.com/abbot/go-http-auth"
	"net/http"
	"strconv"
	"../../db"
	"time"
	"errors"
	"strings"
)

func Get(w http.ResponseWriter, r *auth.AuthenticatedRequest) {

	var err error
	info := struct {
		Title        string
		Nav          []string
		Actualize    int64
		Event tools.Event
		Date string
		SpendingTypes []tools.SpendingType
		Added bool
		Error bool
		ErrorMsg string
		ResultDebts []tools.Debts
		ResultSpending tools.ResultSpending
	}{
		Title: "évènement",
		Nav: tools.GenerateNav(r.Username),
		Actualize: 0,
                Event: tools.Event{},
		Date: time.Now().Format("2006-01-02"),
		SpendingTypes: []tools.SpendingType{},
		Added: false,
		Error: false,
		ErrorMsg: "",
		ResultDebts: []tools.Debts{},
		ResultSpending: tools.ResultSpending{},
	}
	
	params := r.URL.Query()
	
	p, _ := params["get"]
	
	info.Event.Id, err = strconv.ParseInt(p[0],10,64)
	if err != nil {
		tools.Error(err)
		tmpl.Template500(w, r)
		return
	}
	
	user, _ := tools.Users[r.Username]
	
	ok, err := db.CheckEventForThisUser(&user, &info.Event)
	if err != nil {
		tools.Error(err)
		tmpl.Template500(w, r)
		return
	}
	
	if !ok {
		http.Redirect(w, &r.Request, "/events", http.StatusForbidden)
		return
	}
	
	var spendingTypes = make(map[int64]string)
	
	if err = db.LoadSpendingType(&info.SpendingTypes, spendingTypes); err != nil {
		tools.Error(err)
		tmpl.Template500(w, r)
		return
	}
	
	if err = db.LoadThisEvent(&info.Event); err != nil {
		tools.Error(err)
		tmpl.Template500(w, r)
		return
	}
	
	if err = delSpending(r, &info.Event); err != nil {
		tools.Error(err)
		tmpl.Template500(w, r)
		return
	}
	
	if info.Added, err = addSpending(r, &info.Event, spendingTypes); err != nil {
		tools.Error(err)
		tmpl.Template500(w, r)
		return
	}
	
	info.Title = info.Event.Reference
	info.Actualize = info.Event.Id
	
	calculateDebts(&info.Event, &info.ResultDebts, &info.ResultSpending, spendingTypes)
	
	tmpl.TemplateMe(w, r, "lib/templates/events/get.html", info)
}

func delSpending(r *auth.AuthenticatedRequest, ev *tools.Event) error {
	var err error
	
	if delSpd := r.PostFormValue("delSpd"); delSpd == "" {
		return nil
	} else if delSpd != "Supprimer" {
		return errors.New("bad entry for delSpd -> "+delSpd)
	}
	
	for i, spending := range ev.Spending {
		if id := r.PostFormValue(strconv.FormatInt(spending.Id, 10)+"-del"); id == "on" {
			if err = db.DelThisSpending(&spending.Id); err != nil {
				return err
			} else {
				ev.Spending = ev.Spending[:i+copy(ev.Spending[i:], ev.Spending[i+1:])]
			}
		}
	}
	return nil
}
func addSpending(r *auth.AuthenticatedRequest, ev *tools.Event, spendingTypes map[int64]string) (bool, error) {
	
	var err error
	
	if addSpd := r.PostFormValue("addSpd"); addSpd == "" {
		return false, nil
	} else if addSpd != "Ajouter" {
		return false, errors.New("bad entry for addSpd -> "+addSpd)
	}
	
	var spd tools.Spending
	
	spd.EventId = ev.Id
	
	if amount := r.PostFormValue("amount"); amount == "" {
		return false, nil /*errors.New("bad entry for amount*/
	} else {
		if spd.Amount, err = strconv.ParseFloat(strings.Replace(amount,",",".",-1),64); err != nil {
			return false, err
		}
	}
	
	if desc := r.PostFormValue("desc"); desc == "" {
		return false, nil /*errors.New("bad entry for amount*/
	} else {
		spd.Description = desc
	}
	
	if date := r.PostFormValue("date"); date == "" {
		return false, nil /*errors.New("bad entry for amount*/
	} else {
		if spd.SpendingAt.TimeStruct, err = time.Parse("2006-01-02", date); err != nil {
			return false, err
		}
		spd.SpendingAt.FeedEventTimeFromStruct()
	}
	
	if spdType := r.PostFormValue("spdType"); spdType == "" {
		return false, nil /*errors.New("bad entry for amount*/
	} else {
		var i int64
		if i, err = strconv.ParseInt(spdType, 10, 64); err != nil {
			return false, err
		}
		
		if s, ok := spendingTypes[i]; !ok {
			return false, errors.New("Spending type not existing ! -> "+ spdType)
		} else {
			spd.TypeId = i
			spd.TypeReference = s
		}
	}
	
	if payer := r.PostFormValue("payer"); payer == "" {
		return false, nil /*errors.New("bad entry for amount*/
	} else {
		var i int64
		if i, err = strconv.ParseInt(payer, 10, 64); err != nil {
			return false, err
		}
		
		if u, ok := tools.UsersId[i]; !ok {
			return false, errors.New("Spending type not existing ! -> "+ payer)
		} else {
			spd.PayerId = i
			spd.PayerName = u.Login
		}
	}
	
	var forDebtor []tools.User 
	 
	if allPaid := r.PostFormValue("allPaid"); allPaid == "on" {
		for _, user := range ev.Participants {
			forDebtor = append(forDebtor, user)
		}
	} else {
		for _, user := range ev.Participants {
			if paid := r.PostFormValue(strconv.FormatInt(user.Id, 10)+"-Paid"); paid == "on" {
				forDebtor = append(forDebtor, user)
			}
		}
	}
	
	if len(forDebtor) == 0 || len(forDebtor) > len(ev.Participants) {
		return false, errors.New("Bad numbers of debtors")
	}
	
	for _, user := range forDebtor {
		spd.For = append(spd.For, tools.SpendingFor{user.Id,user.Login,spd.Amount/float64(len(forDebtor))})
	}
	
	if err = db.AddThisSpending(ev, &spd); err != nil {
		return false, err
	}
	
	return true, nil
}

func calculateDebts(ev *tools.Event, ResultDebts *[]tools.Debts, ResultSpending *tools.ResultSpending, spendingTypes map[int64]string) {
	
	var debts = make(map[int64]tools.Debts)
	ResultSpending.TotalSpendingByType = make(map[int64]float64)
	
	for _, participant := range ev.Participants {
		var d = make(map[int64]float64)
		debts[participant.Id]= tools.Debts{participant.Id, tools.UsersId[participant.Id].Login,0,"",d, []tools.ToPrint{}}
	}
	
	for _, spending := range ev.Spending {
		d, _ := debts[spending.PayerId] 
		d.TotalSpending += spending.Amount
		debts[spending.PayerId] = d
		
		ResultSpending.TotalSpending += spending.Amount
		
		if tot, ok := ResultSpending.TotalSpendingByType[spending.TypeId]; ok {
			tot += spending.Amount
			ResultSpending.TotalSpendingByType[spending.TypeId] = tot
		} else {
			ResultSpending.TotalSpendingByType[spending.TypeId] = spending.Amount
		}
		
		for _, spFor := range spending.For {
			if spFor.DebtorId != spending.PayerId {
				dbs,_ := debts[spFor.DebtorId]
				dbs.Debts[spending.PayerId] += spFor.Debt
				debts[spFor.DebtorId] = dbs
			}
		}
	}
	
	for debtorId, debtorDebt := range debts {
		for creditorId, debtorToCreditor := range debtorDebt.Debts {
			if creditorDebt, ok := debts[creditorId].Debts[debtorId]; ok {
				switch {
				case debtorToCreditor == creditorDebt :
					delete(debts[creditorId].Debts, debtorId)
					delete(debts[debtorId].Debts, creditorId)
				case debtorToCreditor > creditorDebt :
					delete(debts[creditorId].Debts, debtorId)
					debts[debtorId].Debts[creditorId]=debtorToCreditor-creditorDebt
				case debtorToCreditor < creditorDebt  :
					debts[creditorId].Debts[debtorId]=creditorDebt-debtorToCreditor
					delete(debts[debtorId].Debts, creditorId)
				}
			}
		}
	}
	
	for _, participant := range ev.Participants {
		d, _ := debts[participant.Id]
		
		d.STotalSpending = strconv.FormatFloat(d.TotalSpending,'f',2,64)
		
		for id, amount := range d.Debts {
			var sd = tools.ToPrint{tools.UsersId[id].Login, strconv.FormatFloat(amount,'f',2,64)}
			d.SDebts = append(d.SDebts, sd)
		}
		
		*ResultDebts = append(*ResultDebts, d)
	}
	ResultSpending.Feed(spendingTypes)
}