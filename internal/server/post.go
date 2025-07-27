package server

import (
	"Pet-Pet/internal/database"
	"net/http"
	"encoding/json"
)

func Decode(r *http.Request) database.InfUnit{
	var inf database.InfUnit
	err := json.NewDecoder(r.Body).Decode(&inf)
	if err != nil{
		panic(err)
	}
	return inf
}


func Post(w http.ResponseWriter, r *http.Request){
	inf := Decode(r)
	database.PostDB(inf)
	Out(inf, w, http.StatusOK)

}