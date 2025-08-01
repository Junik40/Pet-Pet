package server

import (
	"Pet-Pet/internal/database"
	"net/http"
)


func Put(w http.ResponseWriter, r *http.Request){
	inf := Decode(r)
	database.PutDB(inf)
	Out(inf, w, http.StatusOK)
}