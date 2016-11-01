package logr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	gorillacontext "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type (
	Goal struct {
		Name       string
		Notes      string `json:"Notes,omitempty"`
		CreatedOn  time.Time
		ModifiedOn time.Time `json:"ModifiedOn,omitempty"`
	}
	Goals []Goal
)

func HandleGoalGet(w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	var c context.Context

	if val, ok := gorillacontext.GetOk(r, "Context"); ok {
		c = val.(context.Context)
	} else {
		c = appengine.NewContext(r)
	}
	fmt.Println("######################", r.URL)

	params := mux.Vars(r)

	goalName, exists := params["goal"]
	if !exists {
		http.Error(w, "Goal parameter is missing in URI", http.StatusBadRequest)
		return
	}

	goal := Goal{}
	goal.Name = goalName

	// if given goal is not found, return appropriate error
	if err := goal.Get(c); err == ErrorNoMatch {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(goal); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// Get retrieves the record based on the provided key.
//
func (goal *Goal) Get(c context.Context) (err error) {
	key := datastore.NewKey(c, "Goal", goal.Name, 0, nil)

	err = datastore.Get(c, key, goal)
	if err != nil && err.Error() == "datastore: no such entity" {
		err = ErrorNoMatch
	}

	return
}
