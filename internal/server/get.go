package server

import(
	"net/http"
	data "Pet-Pet/internal/database"
	"github.com/gorilla/mux"
)

func Gets(w http.ResponseWriter, r *http.Request){
	Output := data.GetDB()
	Out(Output, w, http.StatusOK)
}

func GetId(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	Output := data.GetIdDB(id)
	Out(Output, w, http.StatusOK)
}
func GetSeviceName(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	name := vars["ServiceName"]
	Output := data.GetServiceNameDB(name)
	Out(Output, w, http.StatusOK)
}