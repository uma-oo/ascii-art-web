package server

import (
	"net/http"

	"asciiArt/functions"
)

var Err_Passed string

type Info struct {
	Banner string
	Text   string
}

var Data = Info{}

type Result struct {
	Art   string
	Error string
}

var res = Result{}

var data = struct {
	Info   Info
	Result Result
}{
	Info:   Data,
	Result: res,
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Err_Passed = "404 Page Not Found"
		renderTemplateError(w, "index.html", Err_Passed, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		Err_Passed = "405 Method Not Allowed!"
		renderTemplateError(w, "index.html", Err_Passed, http.StatusMethodNotAllowed)
		return
	}

	renderTemplate(w, "index.html", data)
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		Err_Passed = "404 Page Not Found"
		renderTemplateError(w, "index.html", Err_Passed, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		Err_Passed = "405 Method Not Allowed!"
		renderTemplateError(w, "index.html", Err_Passed, http.StatusMethodNotAllowed)
		return
	}

	if extractFormData(w, r) {
		art := functions.HandleData(Data.Text, Data.Banner)
		res = Result{Art: art, Error: ""}

		renderTemplate(w, "index.html", data)
		res = Result{}

	}
}
