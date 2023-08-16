package shortener

import (
	"encoding/json"
	"net/http"
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
