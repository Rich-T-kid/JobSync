package Handlers

import "net/http"


func SettingsHome(w http.ResponseWriter,r *http.Request){
	info := []byte("hello world")
	w.Write(info)
}
func SettingsHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case "GET":
		renderTemplate(w,"Settings.html",nil) 
	case "PUT":
		return
	default:
		http.Error(w,"Method not allowed" , http.StatusMethodNotAllowed)
}}


