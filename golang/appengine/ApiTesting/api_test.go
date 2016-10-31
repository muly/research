package logr

import (
	//"fmt"
	"net/http"
	"net/http/httptest"
	//"strings"
	"testing"

	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
)

var (
	Instance   aetest.Instance
	testserver *httptest.Server
	goalUrl    string
)

func init() {
	testserver = httptest.NewServer(Handlers())

	goalUrl = testserver.URL + "/goal/test1"

}

func TestGoalGet(t *testing.T) {
	opt := &aetest.Options{AppID: "unittest", StronglyConsistentDatastore: true}
	inst, err := aetest.NewInstance(opt)
	if err != nil {
		t.Error(err.Error())
	}
	defer inst.Close()

	req, err := inst.NewRequest("GET", goalUrl, nil)
	if err != nil {
		t.Error(err.Error())
	}

	c := appengine.NewContext(req)
	/*	if err != nil {
		fmt.Println(err.Error())
	}*/

	t.Log(c)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}

	t.Log(res)

	t.Log(goalUrl)

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
