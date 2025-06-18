package main

import (
	"encoding/json"
	"net/http"

	"github.com/tgrif-dev/inversionwebsite/utils"
)

type RequestBody struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Html    string `json:"html"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var body RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = utils.SendEmail(body.To, body.Subject, body.Html)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Email sent successfully"))
}
