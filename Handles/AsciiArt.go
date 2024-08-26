package Handle

import (
	"net/http"
	"text/template"

	Func "asciiartweb/Function"
)

// /////   work on the results page and send the text to ascii-art
func AsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorHandl(w, http.StatusMethodNotAllowed)
		return
	}
	input := r.FormValue("text")
	if !(IsEmpty(input)) || !(Is_print(input)) || (len(input) > 1000) {
		ErrorHandl(w, http.StatusBadRequest)
		return
	}
	banner := r.FormValue("slection")
	if !(banner == "shadow" || banner == "standard" || banner == "thinkertoy") {
		ErrorHandl(w, http.StatusInternalServerError)
		return
	}

	templat, err := template.ParseFiles("templtes/resulte.html")
	if err != nil {
		ErrorHandl(w, http.StatusNotFound)
		return
	}
	data, errr := Func.Ascii_art_fs(input, banner)
	if errr != nil {
		ErrorHandl(w, http.StatusInternalServerError)
		return
	}
	err = templat.Execute(w, data)
	if err != nil {
		ErrorHandl(w, http.StatusNotFound)
		return
	}
}
