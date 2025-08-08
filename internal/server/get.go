package server

import (
	data "Pet-Pet/internal/database"
	"net/http"
)

func Gets(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("userId")
	serviceName := r.URL.Query().Get("serviceName")
	if id == "" && serviceName == "" {
		Output := data.GetAllDB()
		Out(Output, w, http.StatusOK)
		return
	}
	if id != "" && serviceName == "" {
		Output := data.GetUuidDB(id)
		Out(Output, w, http.StatusOK)
		return
	}
	if id == "" && serviceName != "" {
		Output := data.GetServiceNameDB(serviceName)
		Out(Output, w, http.StatusOK)
		return
	}
	if id != "" && serviceName != "" {
		Output := data.GetIdServiceNameDB(id, serviceName)
		Out(Output, w, http.StatusOK)
		return
	}
}
func Money(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("userId")
	if id == "" {
		Out(nil, w, http.StatusBadRequest)
		return
	}
	Output, err := data.MoneyDB(id)
	if err != nil {
		Out(nil, w, http.StatusBadRequest)
		return
	}
	Out(Output, w, http.StatusOK)
}

