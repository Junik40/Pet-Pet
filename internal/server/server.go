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
	router.HandleFunc("/infunit", Gets).Methods("GET")
	router.HandleFunc("/infunit/{id}", GetId).Methods("GET")
	router.HandleFunc("/infunit/{ServiceName}", GetSeviceName).Methods("GET")
	router.HandleFunc("/infunit", Post).Methods("POST")
	http.ListenAndServe(":8080", router)

}