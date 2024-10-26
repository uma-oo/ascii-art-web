package server

import (
	"html/template"
	"net/http"
)

var Templates *template.Template

func Init(templates *template.Template) {
	Templates = templates
}

func renderTemplateError(w http.ResponseWriter, template string, data interface{}, status int) {
	err := Templates.ExecuteTemplate(w, template, data)
	if err != nil {
		renderTemplateError(w, template, data, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
}

func renderTemplate(w http.ResponseWriter, template string, data interface{}) {
	err := Templates.ExecuteTemplate(w, template, data)
	if err != nil {
		renderTemplateError(w, template, data, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
}

func extractFormData(w http.ResponseWriter, r *http.Request) bool {
	err := r.ParseForm()
	if err != nil {
		Err_Passed = "400 Bad Request!"
		renderTemplateError(w, "index.html", Err_Passed, http.StatusBadRequest)
		return false
	}

	text := r.FormValue("body")
	banner := r.FormValue("banner")

	if text == "" {
		res.Error = "You need to provide a text!"
	}
	if len(text) > 1000 {
		Err_Passed = "400 Bad Request!"
		renderTemplateError(w, "index.html", Err_Passed, http.StatusBadRequest)
		return false
	}
    
	Data = Info{Banner: banner, Text: text}

	return true
}
