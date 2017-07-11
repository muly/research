package main

import (
	"github.com/google/jsonapi"
	"github.com/gorilla/mux"
	"net/http"
)

type Contact struct {
	ID    int    `jsonapi:"primary,contact"`
	Name  string `jsonapi:"attr,name"`
	Email string `jsonapi:"attr,email"`
	//HomeAddress *Address `jsonapi:"relation,homeaddress"`
	PhoneNumbers []*PhoneNumber `jsonapi:"relation,phonenumbers"`
}

type Address struct {
	ID    int    `jsonapi:"primary,address"`
	Line1 string `jsonapi:"attr,line1"`
	City  string `jsonapi:"attr,city"`
}

type PhoneNumber struct {
	ID    int    `jsonapi:"primary,phonenumbers"`
	Phone string `jsonapi:"attr,phone"`
	Type string `jsonapi:"attr,type"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/jsonapi", apiHandler)
	http.ListenAndServe(":8080", r)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {

	cs := []*Contact{}
	//c := Contact{}
	//a:= Address{Line1:"100 Main St", City:"SFO"}
	p:=[]*PhoneNumber{&PhoneNumber{Phone:"1234567890", Type:"mobile"}, &PhoneNumber{Phone:"1111111111", Type:"work"}}

	//cs = append(cs, &Contact{ID: 1, Name: "contact A", Email: "email123A@abc.com", HomeAddress:&a})
	cs = append(cs, &Contact{ID: 2, Name: "contact B", Email: "email123B@abc.com", PhoneNumbers:p})
	cs = append(cs, &Contact{ID: 3, Name: "contact C", Email: "email123C@abc.com"})

	if err := jsonapi.MarshalPayload(w, cs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
