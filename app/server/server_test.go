package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var apiTestsCases = []struct {
	path               string
	expectedStatusCode int
}{
	{"/", http.StatusNotFound},
	{"/foo", http.StatusNotFound},
}

func TestAPI(t *testing.T) {
	for _, n := range apiTestsCases {
		req, err := http.NewRequest("GET", n.path, nil)
		if err != nil {
			t.Fatalf(err.Error())
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(notFound)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != n.expectedStatusCode {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, n.expectedStatusCode)
		}
	}
	t.Log(len(apiTestsCases), "test cases")
}
