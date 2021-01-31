package handler

import (
	"net/http"
)

// Healthz returns the health of the service.
// Handles : /_healthz
// Returns : 200
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
