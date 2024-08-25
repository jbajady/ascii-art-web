package Handle

import (
	"net/http"
	"text/template"
)
///// the main page of the site being worked on.
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandl(w, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		ErrorHandl(w, http.StatusMethodNotAllowed)
		return
	}
	template, er := template.ParseFiles("templtes/inputtext.html")
	if er != nil {
		ErrorHandl(w, http.StatusInternalServerError)
		return
	}

	er = template.Execute(w, nil)
	if er != nil {
		ErrorHandl(w, http.StatusInternalServerError)
		return
	}
}
