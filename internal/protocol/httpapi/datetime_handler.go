package httpapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/0himera/go-microne/internal/domains"
)

type datetimeResponse struct{
	Datetime string `json:"datetime"`
}

func CreateDatetimeHandler(dtService *domains.DatetimeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request)  {
		unixSeconds := dtService.CurrentUnixSeconds()

		t := time.Unix(unixSeconds, 0).UTC()
		iso8601 := t.Format(time.RFC3339)
		
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(w).Encode(datetimeResponse{Datetime: iso8601})
	}

}