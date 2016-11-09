package logr

import (
	"encoding/json"
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

func HandleGoalPost(w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)
	//fmt.Println("###################### HandleGoalPost:", r.URL)
	//c := appengine.NewContext(r)
	var c context.Context
	if val, ok := gorillacontext.GetOk(r, "Context"); ok {
		c = val.(context.Context)
	} else {
		c = appengine.NewContext(r)
	}

	goal := Goal{}

	if err := json.NewDecoder(r.Body).Decode(&goal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//  if record already exists with the same goal name, then return
	goalSrc := Goal{}
	goalSrc.Name = goal.Name
	if err := goalSrc.Get(c); err == ErrorNoMatch {
		// do nothing
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		http.Error(w, "record already exists", http.StatusBadRequest)
		return
	}

	goal.CreatedOn = time.Now()

	if err := goal.Put(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(goal); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// Put saves the goal record to database. In this case to Google Appengine Datastore. If already exists, the record will be overwritten.
func (goal *Goal) Put(c context.Context) error {

	// generate the key
	key := datastore.NewKey(c, "Goal", goal.Name, 0, nil)

	// put the record into the database and capture the key
	key, err := datastore.Put(c, key, goal)
	if err != nil {
		return err
	}

	// read from database into the same variable
	if err = datastore.Get(c, key, goal); err != nil {
		return err
	}

	return err
}
