package ascii

import (
	"fmt"
	"net/http"
)

func Download(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		HandleError(w, 404)
	}
	v := r.FormValue("Download")
	x := r.FormValue("test")
	fmt.Println(x)
	w.Header().Set("Content-Disposition", "attachment; filename=ascii-art"+x)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(v)))

	_, err := w.Write([]byte(v))
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
	}
}
