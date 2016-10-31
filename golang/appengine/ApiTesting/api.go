package main

import (
	"fmt"
	//"net/http"
	//"strings"
	//"testing"

	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
)

var (
	Instance aetest.Instance
)

func main() {
	opt := &aetest.Options{AppID: "unittest", StronglyConsistentDatastore: true}
	inst, err := aetest.NewInstance(opt)
	if err != nil {
		fmt.Println(err.Error())
	}

	req, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	appengine.NewContext(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(c)

	/*
		//
		//
		//
		//
		//
		//
		body := strings.NewReader("")

		request, err := http.NewRequest("GET", "", body) //inst.NewRequest("GET", goalUrl, body) //
		if err != nil {
			t.Error(err)
		}
		t.Log(request)

		c := appengine.NewContext(request) // ERROR: appengine: NewContext passed an unknown http.Request
		t.Log(c)*/

}
