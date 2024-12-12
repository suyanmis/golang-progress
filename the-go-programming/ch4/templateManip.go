package main

import (
	htemplate "html/template"
	"log"
	ttemplate "text/template"
	"time"
)

func main() {
	// Text manipulation, similar with helm
	response := ""
	const textTempl = `{{.TotalCount}} issues: 
	{{range .Items}}----------------------------------------
	Number: {{.Number}}
	User: {{.User.Login}}
	Title: {{.Title | printf "%64s"}}
	Age: {{.CreatedAt | daysAgo}}days
	{{end}}`
	report, err := ttemplate.New("report").Funcs(ttemplate.FuncMap{"daysAgo": daysAgo}).Parse(textTempl)
	if err != nil {
		log.Fatal(err)
	}
	report.Execute(response)
	// HTML Manipulation
	var issueList = htemplate.Must(htemplate.New(("issuelist")).Parse(`
	<h1>{{.TotalCount}} issues</h1>
	<table>
	<tr style='text-align: left'>
	  	<th>#</th>
		<th>State</th>
		<th>User</th>
		<th>Title</th>
	</tr>	
	{{range .Items}}
	<tr>
		<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
		<td>{{.State}}</td>
		<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
		<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	</tr>
	{{end}}
	</table>
`))
}
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
