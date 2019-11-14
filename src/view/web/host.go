package web

import (
	"html/template"
	"io"
	"log"

	"nagios-conf-manager/src/model"
)

func PrintHostList(hosts []*model.Host, writer io.Writer) {
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>

	<!-- load bootstrap via CDN -->
	<link rel="icon" href="https://xarevision.pt/xarevision/wp-content/uploads/2018/06/cropped-logo_x_flat-01-32x32.png" sizes="32x32">
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.0/css/bootstrap.min.css">
	<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.4.1/jquery.min.js"></script>
	<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.0/js/bootstrap.min.js"></script>
	<script src="https://code.jquery.com/jquery-3.4.1.min.js"></script>

	<style>
		.container {
			width: 95%;
			margin: auto;
			padding: 10px;
		}
	</style>

	</head>
	<body>
<div class="container">
					<table class="table">
						<thead>
							<tr class="row100 head">
								<th class="column100 column1" data-column="column1">HostName</th>
								<th class="column100 column2" data-column="column2">Name</th>
								<th class="column100 column3" data-column="column3">Alias</th>
								<th class="column100 column4" data-column="column4">Address</th>
								<th class="column100 column5" data-column="column5">Template</th>
								<th class="column100 column6" data-column="column6">Use</th>
								<th class="column100 column7" data-column="column7">NotificationInterval</th>
								<th class="column100 column8" data-column="column8">MaxCheckAttempts</th>
							</tr>
						</thead>
						<tbody>
							{{range .Items}}
							<tr class="row100">
								<td class="column100 column1" data-column="column1">{{.HostName}}</td>
								<td class="column100 column2" data-column="column2">{{.Name}}</td>
								<td class="column100 column3" data-column="column3">{{.Alias}}</td>
								<td class="column100 column4" data-column="column4">{{.Address}}</td>
								<td class="column100 column5" data-column="column5">{{.Register}}</td>
								<td class="column100 column6" data-column="column6">{{.Use}}</td>
								<td class="column100 column7" data-column="column7">{{.NotificationInterval}}</td>
								<td class="column100 column8" data-column="column8">{{.MaxCheckAttempts}}</td>
							</tr>
							{{else}}<tr><strong>no hosts fond</strong></tr>{{end}}
						</tbody>
					</table>
				</div>
	</body>
</html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	data := struct {
		Title string
		Items []*model.Host
	}{
		Title: "Web page",
		Items: hosts,
	}

	err = t.Execute(writer, data)

}
