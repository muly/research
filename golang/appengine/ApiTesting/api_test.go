package logr

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	testserver *httptest.Server
	goalUrl    string
)

func init() {
	testserver = httptest.NewServer(Handlers())

	goalUrl = testserver.URL + "/goal/123"

}

func Test123(t *testing.T) {

	//inst, err := aetest.NewInstance(nil) // ERROR: "appengine: NewContext passed an unknown http.Request"
	//inst, err := aetest.NewInstance(&aetest.Options{}) // ERROR: "appengine: NewContext passed an unknown http.Request"
	//inst, err := aetest.NewInstance(&aetest.Options{AppID: "testApp", StronglyConsistentDatastore: true}) // ERROR: Unable to assign value 'testApp' to attribute 'application'. Value 'testApp' for application does not match expression...
	//inst, err := aetest.NewInstance(&aetest.Options{AppID: "testapp", StronglyConsistentDatastore: true}) // ERROR: "appengine: NewContext passed an unknown http.Request"
	//inst, err := aetest.NewInstance(&aetest.Options{AppID: "testapp", StronglyConsistentDatastore: false}) // ERROR: "appengine: NewContext passed an unknown http.Request"

	/*if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()*/

	body := strings.NewReader("")

	request, err := http.NewRequest("GET", goalUrl, body) //inst.NewRequest("GET", goalUrl, body) //
	if err != nil {
		t.Error(err)
	}
	t.Log(request)

	//c := appengine.NewContext(request)
	//t.Log(c)

	res, err := http.DefaultClient.Do(request) // ERROR: "appengine: NewContext passed an unknown http.Request"

	if err != nil {
		t.Error(err)
	}

	t.Log(res)

	t.Log(goalUrl)

	/*if res.StatusCode != 201 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}*/

}
