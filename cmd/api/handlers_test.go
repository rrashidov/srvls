package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testServer struct {
	*httptest.Server
}

func TestListFunctions(t *testing.T) {
	requestAndCheck(t, "get", "/api/v1/testtenantid/functions", "List functions")
}

func TestGetFunction(t *testing.T) {
	requestAndCheck(t, "get", "/api/v1/testtenantid/functions/testfunctionid", "Get function")
}

func TestCreateNewFunction(t *testing.T) {
	requestAndCheck(t, "post", "/api/v1/testtenantid/functions", "Create function")
}

func TestUpdateFunction(t *testing.T) {
	requestAndCheck(t, "put", "/api/v1/testtenantid/functions/testfunctionid", "Update function")
}

func TestDeleteFunction(t *testing.T) {
	requestAndCheck(t, "delete", "/api/v1/testtenantid/functions/testfunctionid", "Delete function")
}

func TestListFunctionExecutions(t *testing.T) {
	requestAndCheck(t, "get", "/api/v1/testtenantid/functions/testfunctionid/executions", "List function executions")
}

func TestGetFunctionExecution(t *testing.T) {
	requestAndCheck(t, "get", "/api/v1/testtenantid/functions/testfunctionid/executions/testexecutionid", "Get function execution")
}

func TestTriggerFunctionExecution(t *testing.T) {
	requestAndCheck(t, "post", "/api/v1/testtenantid/functions/testfunctionid/executions", "Trigger function execution")
}

func requestAndCheck(t *testing.T, m string, url string, expBody string) (int, []byte) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())

	defer ts.Close()

	var sc int
	var body []byte

	switch strings.ToUpper(m) {
	case http.MethodGet:
		sc, body = ts.get(t, url)
	case http.MethodPost:
		sc, body = ts.post(t, url)
	case http.MethodPut:
		sc, body = ts.put(t, url)
	case http.MethodDelete:
		sc, body = ts.delete(t, url)
	default:
		t.Errorf("Unexpected method recieved: %q", m)
	}

	if string(body) != expBody {
		t.Errorf("%q to %q failed; wrong response body; wanted: %q; got: %q", m, url, expBody, body)
	}

	return sc, body
}

func newTestApplication(t *testing.T) *application {
	return &application{}
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewServer(h)

	return &testServer{ts}
}

func (ts *testServer) get(t *testing.T, urlPath string) (int, []byte) {
	r, err := http.NewRequest("GET", ts.URL+urlPath, nil)

	if err != nil {
		t.Fatal(err)
	}

	return ts.sendRequest(t, r)
}

func (ts *testServer) post(t *testing.T, urlPath string) (int, []byte) {
	r, err := http.NewRequest("POST", ts.URL+urlPath, nil)

	if err != nil {
		t.Fatal(err)
	}

	return ts.sendRequest(t, r)
}

func (ts *testServer) put(t *testing.T, urlPath string) (int, []byte) {
	r, err := http.NewRequest("PUT", ts.URL+urlPath, nil)

	if err != nil {
		t.Fatal(err)
	}

	return ts.sendRequest(t, r)
}

func (ts *testServer) delete(t *testing.T, urlPath string) (int, []byte) {
	r, err := http.NewRequest("DELETE", ts.URL+urlPath, nil)

	if err != nil {
		t.Fatal(err)
	}

	return ts.sendRequest(t, r)
}

func (ts *testServer) sendRequest(t *testing.T, r *http.Request) (int, []byte) {
	rs, err := ts.Client().Do(r)

	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)

	if err != nil {
		t.Fatal(err)
	}

	return rs.StatusCode, body
}
