package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name           string
	url            string
	method         string
	params         []postData
	expectedStatus int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"gq", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"ms", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"sa", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"rs", "/reservation-summary", "GET", []postData{}, http.StatusOK},
	{"mr", "/make-reservation", "GET", []postData{}, http.StatusOK},

	{"psa", "/search-availability", "POST", []postData{
		{key: "start", value: "10-20-2023"},
		{key: "end", value: "10-21-2023"},
	}, http.StatusOK},

	{"psaj", "/search-availability-json", "POST", []postData{
		{key: "start", value: "10-20-2023"},
		{key: "end", value: "10-21-2023"},
	}, http.StatusOK},

	{"pmr", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "Manoj"},
		{key: "last_name", value: "Gowda"},
		{key: "email", value: "mg@vm.com"},
		{key: "phone", value: "1234567890"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatus {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatus, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, p := range e.params {
				values.Add(p.key, p.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatus {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatus, resp.StatusCode)
			}
		}
	}
}
