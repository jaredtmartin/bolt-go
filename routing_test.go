package bolt

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getHandler(w http.ResponseWriter, r *http.Request) Element {
	return String("GET")
}
func postHandler(w http.ResponseWriter, r *http.Request) Element {
	return String("POST")
}
func testlayoutfunc(title string, r *http.Request, e ...Element) Element {
	return String(fmt.Sprintf("%s %s", title, e[0].Render()))
}
func TestRouteByMethodWithLayout(t *testing.T) {
	mux := http.NewServeMux()
	RouteByMethod(mux, "/test", "Test", testlayoutfunc, getHandler, postHandler)
	tests := []struct {
		method   string
		wantBody string
	}{
		{http.MethodGet, "Test GET"},
		{http.MethodPost, "Test POST"},
		{http.MethodDelete, "Test POST"}, // Should fallback to last handler
		{http.MethodPut, "Test POST"},    // Should fallback to last handler
		{http.MethodPatch, "Test POST"},  // Should fallback to last handler
		{"INVALID", ""},                  // Should not call layout or handler
	}
	for _, tt := range tests {
		req := httptest.NewRequest(tt.method, "/test", nil)
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)
		body := rr.Body.String()
		if body != tt.wantBody {
			t.Errorf("method %s: expected body to contain %q, got %q", tt.method, tt.wantBody, body)
		}
	}
}

func TestRouteByMethod_NoHandlers(t *testing.T) {
	mux := http.NewServeMux()
	layout := func(title string, r *http.Request, e ...Element) Element {
		t.Error("layout should not be called")
		return e[0]
	}
	RouteByMethod(mux, "/nohandler", "NoHandler", layout)
	req := httptest.NewRequest(http.MethodGet, "/nohandler", nil)
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)
	expected := "404 page not found\n"
	if rr.Body.String() != expected {
		t.Errorf("expected %s, got %q", expected, rr.Body.String())
	}
}

func TestRouteByMethod_HandlerIndexBounds(t *testing.T) {
	mux := http.NewServeMux()
	layout := func(title string, r *http.Request, e ...Element) Element {
		return e[0]
	}
	onlyHandler := func(w http.ResponseWriter, r *http.Request) Element {
		return String("ONLY")
	}
	RouteByMethod(mux, "/one", "One", layout, onlyHandler)

	// All valid methods should use the only handler
	methods := []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodPatch}
	for _, method := range methods {
		req := httptest.NewRequest(method, "/one", nil)
		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)
		if rr.Body.String() != "ONLY" {
			t.Errorf("method %s: expected 'ONLY', got %q", method, rr.Body.String())
		}
	}
}
func TestRouteBranch(t *testing.T) {
	mux := http.NewServeMux()
	newHandler := func(w http.ResponseWriter, r *http.Request) Element {
		return String("New")
	}
	createHandler := func(w http.ResponseWriter, r *http.Request) Element {
		return String("Create")
	}
	editHandler := func(w http.ResponseWriter, r *http.Request) Element {
		return String("Edit")
	}
	showHandler := func(w http.ResponseWriter, r *http.Request) Element {
		return String("Show")
	}
	listHandler := func(w http.ResponseWriter, r *http.Request) Element {
		return String("List")
	}
	updateHandler := func(w http.ResponseWriter, r *http.Request) Element {
		return String("Update")
	}
	deleteHandler := func(w http.ResponseWriter, r *http.Request) Element {
		return String("Delete")
	}
	RouteBranch("/dogs", mux, "Test", testlayoutfunc, HandlerBranch{
		"":          {listHandler},
		"/:id":      {showHandler},
		"/new":      {newHandler, createHandler},
		"/:id/edit": {editHandler, updateHandler, deleteHandler},
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

	// rr := httptest.NewRecorder()
	// mux.ServeHTTP(rr, req)
	// expected := "LIST"
	// if rr.Body.String() != expected {
	// 	t.Errorf("expected %s, got %q", expected, rr.Body.String())
	// }
}
