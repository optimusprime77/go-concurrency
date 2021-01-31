package handler

import (
	"context"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type pigeonClient interface {
	FetchData(context.Context, string) ([]byte, error)
}

// Pigeon is a handler that fetches data from the endpoint specified.
// Handles : /api/v1/pigeon
// Returns : 200,504
func Pigeon(e pigeonClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		data, err := e.FetchData(ctx, r.URL.RawQuery)

		if err != nil {
			log.Error(err.Error())
			w.WriteHeader(http.StatusGatewayTimeout)
			w.Write([]byte(http.StatusText(http.StatusGatewayTimeout)))
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
