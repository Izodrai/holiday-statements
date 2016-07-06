package templates

var Form string =
`
<table style="border-style:solid">
	<form name="inscription" method="post" action="/">
		<tr>
			<td>montant :</td>
			<td><input type="text" name="montant"/></td>
		</tr>
		<tr>
			<td>qui a payé ? :</td>
			<td>
				<input type="radio" name="payeur" value="valentin" id="valPayeur" /> <label for="val">Valentin</label><br />
				<input type="radio" name="payeur" value="emma" id="emmPayeur" /> <label for="emm">Emma</label><br />
				<input type="radio" name="payeur" value="justine" id="JusPayeur" /> <label for="jus">Justine</label><br />
				<input type="radio" name="payeur" value="jerome" id="JerPayeur" /> <label for="jer">Jérôme</label><br />
			</td>
		<tr>
		</tr>
			<td>pour qui ? :</td>
			<td>
				<input type="checkbox" name="allPaye" id="allPaye" /> <label for="allPaye">all</label><br />
				<input type="checkbox" name="valPaye" id="valPaye" /> <label for="valPaye">Valentin</label><br />
				<input type="checkbox" name="emmPaye" id="emmPaye" /> <label for="emmPaye">Emma</label><br />
				<input type="checkbox" name="jusPaye" id="jusPaye" /> <label for="jusPaye">Justine</label><br />
				<input type="checkbox" name="jerPaye" id="jerPaye" /> <label for="jerPaye">Jérôme</label><br />
			</td>
		</tr>
		<tr>
			<td>description :</td>
			<td><textarea type="text" name="description" id="description"></textarea></td>
		</tr>
		<tr>
			<td colspan="2"><input type="submit" name="valider" value="OK"/></td>
		</tr>
	</form>
</table>
<h3>Vos dernières dépenses :</h3>
`

var Bottom string =
`
	</body>
</html>
`

var Top string =
`
<!DOCTYPE html>
<html>
	<head>
		<!-- En-tête de la page -->
		<meta charset="utf-8" />
		<title>add payement</title>
	</head>

	<body>
`

var TopNothing string =
`
<!DOCTYPE html>
<html>
	<head>
		<!-- En-tête de la page -->
		<meta charset="utf-8" />
		<title>add payement</title>
	</head>
		
	<body>
		<p>Nothing to add</p>
`