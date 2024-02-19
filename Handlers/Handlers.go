package Handlers 
import ( //"os" 
	"net/http"
	//"path/filepath"
	"html/template"


)

func HomeHandler(w http.ResponseWriter,r *http.Request){
	renderTemplate(w,"home.html",nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) { // this is for static html page rendering. 
	tmplPath := "../static" + tmpl
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
