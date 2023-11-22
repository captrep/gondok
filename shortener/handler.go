package shortener

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type J struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}

func writeMessage(w http.ResponseWriter, code int, data interface{}, response string) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	var j J
	if response == "error" {
		j = J{
			Code:   code,
			Errors: data,
		}
	} else if response == "success" {
		j = J{
			Code: code,
			Data: data,
		}
	}
	json.NewEncoder(w).Encode(j)
}

func writeError(w http.ResponseWriter, code int, err error) {
	writeMessage(w, code, err.Error(), "error")
}

func create(w http.ResponseWriter, req *http.Request) {

	longURL := req.FormValue("long")
	shortURL := req.FormValue("short")
	resp, err := CreateLink(req.Context(), shortURL, longURL)
	if err != nil {
		var errStr string
		if strings.Contains(err.Error(), "failed") {
			errStr = "internal server error"
		} else {
			errStr = err.Error()
		}
		tmpl, err := template.New("t").Parse(fmt.Sprintf("<p class=\"mt-1 mb-5 text-md text-red-500\">%s</p>", errStr))
		if err != nil {
			log.Print(err)
		}
		tmpl.Execute(w, nil)
		return
	}
	tmpl, err := template.New("t").Parse(fmt.Sprintf(" <p class=\"mt-1 mb-5 text-md font-mono text-indigo-500\">success, here is ur shorten link http://127.0.0.1:8000/%s", resp.ShortURL))
	if err != nil {
		log.Print(err)
	}
	tmpl.Execute(w, nil)
}

func createLink(w http.ResponseWriter, req *http.Request) {
	var request struct {
		ShortURL string `json:"short_url"`
		LongURL  string `json:"long_url"`
	}
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		return
	}
	resp, err := CreateLink(req.Context(), request.ShortURL, request.LongURL)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	writeMessage(w, http.StatusCreated, resp, "success")
}

func redirectLink(w http.ResponseWriter, req *http.Request) {
	link, err := GetLink(req.Context(), chi.URLParam(req, "url"))
	if err != nil {
		http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
		return
	}
	http.Redirect(w, req, link.LongURL, http.StatusFound)
}
