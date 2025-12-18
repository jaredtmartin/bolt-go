package bolt

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testlayoutfunc(r *http.Request, e ...Element) Element {
	return String(fmt.Sprintf("%s %s", GetPageTitle(e...), e[0].Render()))
}

func TestRouteByMethod(t *testing.T) {
	mux := http.NewServeMux()
	getHandler := func(w http.ResponseWriter, r *http.Request) (Element, error) {
		return Div("Get").PageTitle("Test Get"), nil
	}
	postHandler := func(w http.ResponseWriter, r *http.Request) (Element, error) {
		return Div("Post").PageTitle("Test Post"), nil
	}
	deleteHandler := func(w http.ResponseWriter, r *http.Request) (Element, error) {
		return Div("Delete").PageTitle("Test Delete"), nil
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
		{"/test", http.MethodGet, ` &lt;div page-title=&quot;Test Get&quot; class=&quot;Get&quot;&gt;&lt;/div&gt;`},
		{"/test", http.MethodPost, ` &lt;div page-title=&quot;Test Post&quot; class=&quot;Post&quot;&gt;&lt;/div&gt;`},
		{"/test", http.MethodDelete, ` &lt;div page-title=&quot;Test Delete&quot; class=&quot;Delete&quot;&gt;&lt;/div&gt;`},
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
	newHandler := func(w http.ResponseWriter, r *http.Request) (Element, error) {
		return String("New").PageTitle("Test"), nil
	}
	createHandler := func(w http.ResponseWriter, r *http.Request) (Element, error) {
		return String("Create").PageTitle("Test"), nil
	}
	editHandler := func(w http.ResponseWriter, r *http.Request) (Element, error) {
		return String("Edit").PageTitle("Test"), nil
	}
	showHandler := func(w http.ResponseWriter, r *http.Request) (Element, error) {
		return String("Show").PageTitle("Test"), nil
	}
	listHandler := func(w http.ResponseWriter, r *http.Request) (Element, error) {
		return String("List").PageTitle("Test"), nil
	}
	updateHandler := func(w http.ResponseWriter, r *http.Request) (Element, error) {
		return String("Update").PageTitle("Test"), nil
	}
	deleteHandler := func(w http.ResponseWriter, r *http.Request) (Element, error) {
		return String("Delete").PageTitle("Test"), nil
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
