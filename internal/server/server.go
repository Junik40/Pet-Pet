package server
import(
	"github.com/gorilla/mux"
	"encoding/json"	
	"net/http"
)

func Out(inf any, w http.ResponseWriter, code int){
	jsonData,err := json.Marshal(inf)
	if err != nil{
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonData)

}

func Init(){
	
	router := mux.NewRouter()
	router.HandleFunc("/infunits", Gets).Methods("GET")
	router.HandleFunc("/infunit", Post).Methods("POST")
	router.HandleFunc("/infunit", Delete).Methods("DELETE")
	router.HandleFunc("/infunit/user", DeleteUuid).Methods("DELETE")
	router.HandleFunc("/infunit", Put).Methods("PUT")
	http.ListenAndServe(":8080", router)

}