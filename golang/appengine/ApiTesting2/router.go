package apitest

import (

	//"errors"
	"encoding/json"
	//"fmt"
	"net/http"
	"time"

	//"ctrl"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"

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

type Goal struct {
	Name       string
	Notes      string `json:"Notes,omitempty"`
	CreatedOn  time.Time
	ModifiedOn time.Time `json:"ModifiedOn,omitempty"`
}

func HandleApiTest(w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	/*
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
	*/
	//fmt.Println(res)

	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	res, err := client.Get(goalUrl)
	if err != nil {
		http.Error(w, "Error with client.Get(goalUrl):"+err.Error(), http.StatusInternalServerError)
		return
	}

	goal := Goal{}

	if err := json.NewDecoder(res.Body).Decode(&goal); err != nil {
		http.Error(w, "Error with json.NewDecoder(r.Body).Decode(&goal):"+err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(res.Status); err != nil {
		http.Error(w, "Error with json.NewEncoder(w).Encode(res.Status):"+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(goal); err != nil {
		http.Error(w, "Error with json.NewEncoder(w).Encode(goal): "+err.Error(), http.StatusInternalServerError)
		return
	}
}
