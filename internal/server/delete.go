package server

import(
	"net/http"
	"Pet-Pet/internal/database"
	"strconv"
)

func Delete( w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("userId")
	if id == "" {
		Out(nil, w, http.StatusBadRequest)
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		Out(nil, w, http.StatusBadRequest)
		return
	}
	database.DeleteDB(idInt)
	Out("All Good", w, http.StatusOK)
}

func DeleteUuid( w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("userId")
	if id == "" {
		Out(nil, w, http.StatusBadRequest)
		return
	}
	database.DeleteUuidDB(id)
	Out("All Good", w, http.StatusOK)
}