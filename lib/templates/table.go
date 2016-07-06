package templates

import (
	"../tools"
	"strconv"
)
func GenerateTable(ds []tools.Depense, res tools.Result) string {
	s := `
<table style="border-style:solid">
`
	
	s = s + `
<tr>
	<td style="border-style:solid;border-width:0px 0px 2px 0px">Date</td>
	<td style="border-style:solid;border-width:0px 0px 2px 0px">Montant</td>
	<td style="border-style:solid;border-width:0px 0px 2px 0px">Payeur</td>
	<td style="border-style:solid;border-width:0px 0px 2px 0px">Emma</td>
	<td style="border-style:solid;border-width:0px 0px 2px 0px">Justine</td>
	<td style="border-style:solid;border-width:0px 0px 2px 0px">Valentin</td>
	<td style="border-style:solid;border-width:0px 0px 2px 0px">Jérôme</td>
	<td style="border-style:solid;border-width:0px 0px 2px 0px">NbrFor</td>
	<td style="border-style:solid;border-width:0px 0px 2px 0px">Desc</td>
</tr>
	`
	for i := len(ds)-1; i >= 0; i-- {
		s = s + printLineDepenses(ds[i])
	}
	
	s = s + `
</table>

<h3>Total :</h3>

<table style="border-style:solid">`
	
	for payeur, debs := range res.Debts {
		
		var tot float64
		
		switch payeur {
			case  "valentin" :
				tot = res.TotV
			case  "emma" :
				tot = res.TotE
			case  "justine" :
				tot = res.TotJu
			case  "jerome" :
				tot = res.TotJe
		}
		
		s = s + printLineTotal(payeur, tot, debs)
	}
	
	
	s = s + `</table>`
	
	return s
}

func printLineTotal(name string, totalUser float64, debtsUser map[string]float64) string {
	var s string = `
<tr>
	<td style="border-style:solid;border-width:0px 0px 2px 0px">`+name+` à dépensé:</td>
	<td style="border-style:solid;border-width:0px 0px 2px 0px" padding="0.5px" collspan="2">`+strconv.FormatFloat(totalUser,'f',2,64)+` €</td>
</tr>
`
	for name, deb := range debtsUser {
		
		if deb == 0 {
			continue
		}
		s = s + `
<tr>
	<td padding="0.5px"> -> Doit</td>
	<td padding="0.5px">`+strconv.FormatFloat(deb,'f',2,64)+`</td>
	<td padding="0.5px">à `+name+`</td>
</tr>
`
	}

	return s
}

func printLineDepenses(d tools.Depense) string {
	return `
<tr>
	<td padding="0.5px">`+d.Date+`</td>
	<td padding="0.5px">`+strconv.FormatFloat(d.Montant,'f',2,64)+`</td>
	<td padding="0.5px">`+d.Payeur+`</td>
	<td padding="0.5px">`+strconv.FormatFloat(d.For["emma"],'f',2,64)+`</td>
	<td padding="0.5px">`+strconv.FormatFloat(d.For["justine"],'f',2,64)+`</td>
	<td padding="0.5px">`+strconv.FormatFloat(d.For["valentin"],'f',2,64)+`</td>
	<td padding="0.5px">`+strconv.FormatFloat(d.For["jerome"],'f',2,64)+`</td>
	<td padding="0.5px">`+strconv.Itoa(d.NFor)+`</td>
	<td padding="0.5px">`+d.Desc+`</td>
</tr>
`
}