package logr

import (
	//"fmt"
	//"net/http"
	"net/http/httptest"
	//"strings"
	"testing"

	//"google.golang.org/appengine"
	//"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/urlfetch"
	//gorillacontext "github.com/gorilla/context"
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
	/*opt := &aetest.Options{AppID: "unittest", StronglyConsistentDatastore: true}
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
		if err != nil {
		fmt.Println(err.Error())
	}


	t.Log(c)
	*/

	//ctx := appengine.NewContext(r)
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	client := urlfetch.Client(ctx)

	res, err := client.Get(goalUrl)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res.Status)

	/*ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	t.Log(ctx)

	h := Handlers()
	record := httptest.NewRecorder()

	req, err := http.NewRequest("GET", goalUrl, nil)
	if err != nil {
		t.Error(err.Error())
	}

	gorillacontext.Set(req, "Context", ctx)

	h.ServeHTTP(record, req)
	t.Log(record.Code)*/

	/*res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}

	t.Log(res)*/

	//t.Log(goalUrl)

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
