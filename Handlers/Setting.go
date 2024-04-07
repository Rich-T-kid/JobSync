package Handlers

import "net/http"

// define  a render function in each of the subrouters. Mabey theres  a cleaner way to render html templates i just dont care tight now

func SettingsHome(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "Setting/settings.html", nil)
}

func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		renderTemplate(w, "Settings.html", nil)
	case "PUT":
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
