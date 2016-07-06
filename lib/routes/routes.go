package routes

import (
	"os"
	"io"
//     "log"
	"net/http"
	"time"
	"../templates"
	"../tools"
	"strconv"
	"strings"
	"encoding/csv"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	
	var ds []tools.Depense
	var res1, res2 tools.Result
	
	f, err := os.OpenFile("save.csv", os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModeAppend|0755)
	if err != nil {
		tools.FatalError(err)
	}
	
	defer f.Close()
	
	if err = Load(f, &ds); err != nil {
		tools.FatalError(err)
	}
	
	Calc(ds, &res1)
	
	r.ParseForm()
	
	var montant, description, payeur, allPaye, valPaye, emmPaye, jusPaye, jerPaye []string
	
	montant, _ = r.PostForm["montant"]
	description, _ = r.PostForm["description"]
	payeur, _ = r.PostForm["payeur"]
	allPaye, _ = r.PostForm["allPaye"]
	valPaye, _ = r.PostForm["valPaye"]
	emmPaye, _ = r.PostForm["emmPaye"]
	jusPaye, _ = r.PostForm["jusPaye"]
	jerPaye, _ = r.PostForm["jerPaye"]
	
	if len(payeur) == 0 || len(montant) == 0 || len(description) == 0 {
		io.WriteString(w, templates.TopNothing+templates.Form+templates.GenerateTable(ds, res1)+templates.Bottom)
		return
	}
	
	if len(allPaye) == 0 && len(valPaye) == 0 && len(emmPaye) == 0 && len(jusPaye) == 0 && len(jerPaye) == 0 {
		io.WriteString(w, templates.TopNothing+templates.Form+templates.GenerateTable(ds, res1)+templates.Bottom)
		return
	}
	
	var d tools.Depense 
	
	d.For = make(map[string]float64)
	
	d.Payeur = payeur[0]
	
	if d.Montant, err = strconv.ParseFloat(strings.Replace(montant[0],",",".",-1),64); err != nil {
		io.WriteString(w, templates.TopNothing+templates.Form+templates.GenerateTable(ds, res1)+templates.Bottom)
		return
	}
	
	if len(allPaye) != 0 {
		d.For["valentin"]=0
		d.For["emma"]=0
		d.For["justine"]=0
		d.For["jerome"]=0
	}
	
	if len(valPaye) != 0 {
		d.For["valentin"]=0
	}
	
	if len(emmPaye) != 0 {
		d.For["emma"]=0
	}
	
	if len(jusPaye) != 0 {
		d.For["justine"]=0
	}
	
	if len(jerPaye) != 0 {
		d.For["jerome"]=0
	}
	
	d.Desc = description[0]
	
	d.Date = time.Now().Format("2006-01-02")
	
	var row string = d.Date+";"+strconv.FormatFloat(d.Montant,'f',2,64)+";"+d.Payeur+";"
	
	div := d.Montant / float64(len(d.For))
	
	d.NFor = len(d.For)
	
	if _, ok := d.For["emma"]; ok {
		row = row + "X;"+strconv.FormatFloat(div,'f',2,64)+";"
		d.For["emma"]=div
	} else {
		row = row + ";0;"
	}
	if _, ok := d.For["justine"]; ok {
		row = row + "X;"+strconv.FormatFloat(div,'f',2,64)+";"
		d.For["justine"]=div
	} else {
		row = row + ";0;"
	}
	if _, ok := d.For["valentin"]; ok {
		row = row + "X;"+strconv.FormatFloat(div,'f',2,64)+";"
		d.For["valentin"]=div
	} else {
		row = row + ";0;"
	}
	if _, ok := d.For["jerome"]; ok {
		row = row + "X;"+strconv.FormatFloat(div,'f',2,64)+";"
		d.For["jerome"]=div
	} else {
		row = row + ";0;"
	}
	
	row = row + strconv.Itoa(d.NFor)+";"
	row = row + d.Desc +"\n"
	
	if _, err := f.Write([]byte(row)); err != nil {
		panic(err)
	}
	
	ds=append(ds,d)
	
	Calc(ds, &res2)
	
	io.WriteString(w, templates.Top+templates.Form+templates.GenerateTable(ds, res2)+templates.Bottom)
	
}

func addDebts(d tools.Depense, res *tools.Result) {
	for name, ct := range d.For {
		if d.Payeur == name {
			if _, ok := res.Debts[d.Payeur]; !ok {
				debts := make(map[string]float64)
				res.Debts[name] = debts
			}
			continue
		}
		
		if _, ok := res.Debts[name]; !ok {
			debts := make(map[string]float64)
			debts[d.Payeur]=ct
			res.Debts[name] = debts
		} else {
			if val, ok := res.Debts[name][d.Payeur]; !ok {
				res.Debts[name][d.Payeur]=ct
			} else {
				val = val+ct
				res.Debts[name][d.Payeur]=val
			}
		}
	}
}
func Calc(ds []tools.Depense, res *tools.Result) {
	
	res.Debts = make(map[string]map[string]float64)

	res.NPer = 4
	
	for _, d := range ds {
		res.Tot = res.Tot + d.Montant
		
		addDebts(d, res)
		
		if d.Payeur == "valentin" {
			res.TotV = res.TotV + d.Montant
		}
		if d.Payeur == "emma" {
			res.TotE = res.TotE + d.Montant
		}
		if d.Payeur == "justine" {
			res.TotJu = res.TotJu + d.Montant
		}
		if d.Payeur == "jerome" {
			res.TotJe = res.TotJe + d.Montant
		}
	}
	
	res.Average = res.Tot/float64(res.NPer)
	
	for payeurName, debts := range res.Debts {
		for debtName, deb := range debts {
			if d, ok := res.Debts[debtName][payeurName]; ok {
				switch {
				case deb == d :
					res.Debts[debtName][payeurName]=0
					res.Debts[payeurName][debtName]=0
				case deb > d :
					res.Debts[debtName][payeurName]=0
					res.Debts[payeurName][debtName]=deb-d
				case deb < d  :
					res.Debts[debtName][payeurName]=d-deb
					res.Debts[payeurName][debtName]=0
				}
			}
		}
	}
}

			
func Load(f *os.File, ds *[]tools.Depense) error {
	
	reader := csv.NewReader(f)
	reader.Comma = ';'
	reader.FieldsPerRecord = -1
	lines, err := reader.ReadAll()
	if err != nil {
		return err
	}
	for i, line := range lines {
		
		if i == 0 {
			continue
		}
		var d tools.Depense 
		d.For = make(map[string]float64)
	
		d.Date = line[0]
		d.Montant, _ = strconv.ParseFloat(line[1],64)
		d.Payeur = line[2]
		
		if line[3] == "X" {
			div, _ := strconv.ParseFloat(line[4],64)
			d.For["emma"]=div
		}
		if line[5] == "X" {
			div, _ := strconv.ParseFloat(line[6],64)
			d.For["justine"]=div
		}
		if line[7] == "X" {
			div, _ := strconv.ParseFloat(line[8],64)
			d.For["valentin"]=div
		}
		
		if line[9] == "X" {
			div, _ := strconv.ParseFloat(line[10],64)
			d.For["jerome"]=div
		}
		
		d.NFor, _ = strconv.Atoi(line[11]) 
		
		d.Desc = line[12]
		
		*ds=append(*ds,d)
	}
	
	return nil
}