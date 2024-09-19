package ascii

import (
	"fmt"
	"net/http"
)

func Download(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		HandleError(w, http.StatusBadRequest)
	}
	value := r.FormValue("Download")
	ext := r.FormValue("extention")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii-art"+ext)
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(value)))

	_, err := w.Write([]byte(value))
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
	}
}
