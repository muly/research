package logr

import (
	"errors"
	"fmt"
	"net/http"

	//"ctrl"

	"github.com/gorilla/mux"
)

var ErrorNoMatch = errors.New("No Matching Record")

func init() {

	r := Handlers()

	http.Handle("/", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world! 123")

}

func Handlers() *mux.Router {

	r := mux.NewRouter()
	//r.HandleFunc("/", handler).Methods("GET")

	r.HandleFunc("/goal/{goal}", HandleGoalGet).Methods("GET")
	r.HandleFunc("/goal", HandleGoalPost).Methods("POST")

	return r

}
