package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type S struct {
	Name string
	Time time.Time
}

func resourseHandler(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {

		s := S{}

		if err := json.NewDecoder(req.Body).Decode(&s); err != nil {
			rw.Write([]byte(err.Error()))
			return
		}

		fmt.Fprintln(rw, s)

	}
}

func main() {

	http.HandleFunc("/", resourseHandler)

	http.ListenAndServe(":8080", nil)

}
