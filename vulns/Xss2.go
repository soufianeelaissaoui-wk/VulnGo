package vulns

import "net/http"
import "text/template"

func Xss2(w http.ResponseWriter, r *http.Request) {
	error := r.URL.Query().Get("x")
	tmpl := template.New("error")
	tmpl, _ = tmpl.Parse(`{{define "T"}}{{.}}{{end}}`)
	tmpl.ExecuteTemplate(w, "T", error)
}