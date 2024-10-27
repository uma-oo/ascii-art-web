package server

import (
	"net/http"

	"asciiArt/functions"
)

type Info struct {
	Banner    string
	Text      string
	Ascii     string
	ErrorText string
}

var (
	Res        Info
	Err_Passed string
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Err_Passed = "404 Page Not Found"
		renderTemplateError(w, "error_page.html", Err_Passed, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		Err_Passed = "405 Method Not Allowed!"
		renderTemplateError(w, "error_page.html", Err_Passed, http.StatusMethodNotAllowed)
		return
	}

	renderTemplate(w, "index.html", nil)
	Res = Info{}
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Err_Passed = "405 Method Not Allowed!"
		renderTemplateError(w, "error_page.html", Err_Passed, http.StatusMethodNotAllowed)
		return
	}

	if sucess := extractFormData(w, r); sucess {
		art , erreur  := functions.HandleData(Res.Text, Res.Banner)
		if erreur!="" {
			Res.ErrorText=erreur
		}
		Res.Ascii = art
		renderTemplate(w, "index.html", Res)
		Res=Info{}

	}
}
