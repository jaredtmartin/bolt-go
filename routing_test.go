package bolt

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testlayoutfunc(title string, r *http.Request, e ...Element) Element {
	return String(fmt.Sprintf("%s %s", title, e[0].Render()))
}

func TestRouteByMethod(t *testing.T) {
	mux := http.NewServeMux()
	getHandler := func(w http.ResponseWriter, r *http.Request) (string, Element) {
		return "Test", String("Get")
	}
	postHandler := func(w http.ResponseWriter, r *http.Request) (string, Element) {
		return "Test", String("Post")
	}
	deleteHandler := func(w http.ResponseWriter, r *http.Request) (string, Element) {
		return "Test", String("Delete")
	}
	RouteByMethod(mux, "/test", testlayoutfunc, HandlerMethods{
		Get:    getHandler,
		Post:   postHandler,
		Delete: deleteHandler,
	})

	tests := []struct {
		path     string
		method   string
		wantBody string
	}{
		{"/test", http.MethodGet, "Test Get"},
		{"/test", http.MethodPost, "Test Post"},
		{"/test", http.MethodDelete, "Test Delete"},
		{"/test", http.MethodPatch, "404 page not found\n"},
		{"/test", http.MethodPut, "404 page not found\n"},
	}
	for _, tt := range tests {
		req := httptest.NewRequest(tt.method, tt.path, nil)
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)
		body := rr.Body.String()
		if body != tt.wantBody {
			t.Errorf("method %s: expected body to contain %q, got %q", tt.method, tt.wantBody, body)
		}
	}
}

func TestRouteBranch(t *testing.T) {
	mux := http.NewServeMux()
	newHandler := func(w http.ResponseWriter, r *http.Request) (string, Element) {
		return "Test", String("New")
	}
	createHandler := func(w http.ResponseWriter, r *http.Request) (string, Element) {
		return "Test", String("Create")
	}
	editHandler := func(w http.ResponseWriter, r *http.Request) (string, Element) {
		return "Test", String("Edit")
	}
	showHandler := func(w http.ResponseWriter, r *http.Request) (string, Element) {
		return "Test", String("Show")
	}
	listHandler := func(w http.ResponseWriter, r *http.Request) (string, Element) {
		return "Test", String("List")
	}
	updateHandler := func(w http.ResponseWriter, r *http.Request) (string, Element) {
		return "Test", String("Update")
	}
	deleteHandler := func(w http.ResponseWriter, r *http.Request) (string, Element) {
		return "Test", String("Delete")
	}
	RouteBranch(mux, testlayoutfunc, HandlerBranch{
		"/dogs":          {Get: listHandler},
		"/dogs/:id":      {Get: showHandler},
		"/dogs/new":      {Get: newHandler, Post: createHandler},
		"/dogs/:id/edit": {Get: editHandler, Post: updateHandler, Delete: deleteHandler},
	})

	tests := []struct {
		path     string
		method   string
		wantBody string
	}{
		{"/dogs", http.MethodGet, "Test List"},
		{"/dogs/new", http.MethodGet, "Test New"},
		{"/dogs/new", http.MethodPost, "Test Create"},
		{"/dogs/:id/edit", http.MethodGet, "Test Edit"},
		{"/dogs/:id/edit", http.MethodPost, "Test Update"},
		{"/dogs/:id/edit", http.MethodDelete, "Test Delete"},
		{"/dogs/:id", http.MethodGet, "Test Show"},
	}
	for _, tt := range tests {
		req := httptest.NewRequest(tt.method, tt.path, nil)
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)
		body := rr.Body.String()
		if body != tt.wantBody {
			t.Errorf("method %s: expected body to contain %q, got %q", tt.method, tt.wantBody, body)
		}
	}
}
