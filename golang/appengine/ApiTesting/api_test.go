package logr

import (
	"net/http"
	"strings"
	"testing"

	"google.golang.org/appengine"
)

func Test123(t *testing.T) {

	body := strings.NewReader("")

	request, err := http.NewRequest("GET", "", body) //inst.NewRequest("GET", goalUrl, body) //
	if err != nil {
		t.Error(err)
	}
	t.Log(request)

	c := appengine.NewContext(request) // ERROR: appengine: NewContext passed an unknown http.Request
	t.Log(c)

}
