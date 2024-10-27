package server

import (
	"html/template"
	"net/http"
	"regexp"
)

func renderTemplateError(w http.ResponseWriter, filename string, data interface{}, status int) {
	w.WriteHeader(status)
	renderTemplate(w, filename, data)
}

func renderTemplate(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles("templates/" + filename)
	if filename == "error_page.html" && err != nil {
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	if err != nil {
		renderTemplateError(w, "error_page.html", "500 Internal server Error", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func extractFormData(w http.ResponseWriter, r *http.Request) bool {
	err := r.ParseForm()
	if err != nil {
		Err_Passed = "400 Bad Request!"
		renderTemplateError(w, "error_page.html", Err_Passed, http.StatusBadRequest)
		return false
	}

	text := r.FormValue("body")
	if textReg := regexp.MustCompile(`^\r\n+`); textReg.MatchString(text) {
		Res.Text = "\r\n" + text
	} else {
		Res.Text = text
	}
	banner := r.FormValue("banner")
	if !isBanner(banner) {
		renderTemplateError(w, "error_page.html", "400 Bad Request!", http.StatusBadRequest)
		return false
	}

	if text == "" {
		Res.ErrorText = "You need to provide a text!"
	}
	if len(text) > 1000 {
		renderTemplateError(w, "error_page.html", "400 Bad Request!", http.StatusBadRequest)
		return false
	}

	Res.Banner = banner

	return true
}

func isBanner(banner string) bool {
	return banner == "standard" || banner == "shadow" || banner == "thinkertoy"
}
