package apitest

import (

	//"errors"
	"encoding/json"
	//"fmt"
	"net/http"

	//"ctrl"

	"github.com/gorilla/mux"
)

func init() {

	r := Handlers()

	http.Handle("/", r)
}

func Handlers() *mux.Router {

	r := mux.NewRouter()
	//r.HandleFunc("/", handler).Methods("GET")

	r.HandleFunc("/apitest", HandleApiTest).Methods("GET")

	return r

}

var (
	goalUrl string
)

func init() {

	goalUrl = "http://localhost:8080/goal/test1"

}

func HandleApiTest(w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)

	req, err := http.NewRequest("GET", goalUrl, nil)
	if err != nil {
		http.Error(w, "Error with http.NewRequest():"+err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Error with http.DefaultClient.Do():"+err.Error(), http.StatusInternalServerError)
		return
	}

	//fmt.Println(res)

	if err := json.NewEncoder(w).Encode(res.Status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
