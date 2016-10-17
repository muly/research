package guestbook

import (
	//"html/template"
	"encoding/json"
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
	"appengine/user"
)

// [START greeting_struct]
type Greeting struct {
	Author  string
	Content string
	Date    time.Time
}

// [END greeting_struct]

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/sign", sign)
}

// guestbookKey returns the key used for all guestbook entries.
func guestbookKey(c appengine.Context) *datastore.Key {
	// The string "default_guestbook" here could be varied to have multiple guestbooks.
	return datastore.NewKey(c, "Guestbook", "default_guestbook", 0, nil)
}

// [START func_root]
func root(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusNotFound)
		return
	}
	c := appengine.NewContext(r)
	// Ancestor queries, as shown here, are strongly consistent with the High
	// Replication Datastore. Queries that span entity groups are eventually
	// consistent. If we omitted the .Ancestor from this query there would be
	// a slight chance that Greeting that had just been written would not
	// show up in a query.
	// [START query]
	q := datastore.NewQuery("Greeting").Ancestor(guestbookKey(c)).Order("-Date").Limit(10)
	// [END query]
	// [START getall]
	greetings := make([]Greeting, 0, 10)
	if _, err := q.GetAll(c, &greetings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(greetings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// [END func_root]

// [START func_sign]
func sign(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusNotFound)
		return
	}
	// [START new_context]
	c := appengine.NewContext(r)

	g := Greeting{}
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	g.Date = time.Now()

	// [START if_user]
	if u := user.Current(c); u != nil {
		g.Author = u.String()
	}
	// We set the same parent key on every Greeting entity to ensure each Greeting
	// is in the same entity group. Queries across the single entity group
	// will be consistent. However, the write rate to a single entity group
	// should be limited to ~1/second.
	key := datastore.NewIncompleteKey(c, "Greeting", guestbookKey(c))
	_, err := datastore.Put(c, key, &g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// [END func_sign]
