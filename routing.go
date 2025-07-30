package bolt

import (
	"fmt"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request) (string, Element)
type Layout func(string, *http.Request, ...Element) Element
type HandlerBranch map[string][]Handler
type UrlType string

const ListUrl UrlType = "list"
const SaveUrl UrlType = "edit"
const EditUrl UrlType = "edit"
const NewUrl UrlType = "new"
const ShowUrl UrlType = ""
const DeleteUrl UrlType = ""

// type path map[string][]Handler

func methodIndexRequested(method string) (int, bool) {
	switch method {
	case http.MethodGet:
		return 0, true
	case http.MethodPost:
		return 1, true
	case http.MethodDelete:
		return 2, true
	case http.MethodPut:
		return 3, true
	case http.MethodPatch:
		return 4, true
	default:
		return -1, false // Invalid method
	}
}

// Routes requests for a given path based on the HTTP method.
// mux is the HTTP ServeMux to register the routes with.
// path is the base path for the routes.
// handlers is any number of handler functions that return bolt elements
func RouteByMethod(mux *http.ServeMux, path string, layout Layout, handlers ...Handler) {
	if len(handlers) == 0 {
		// No handlers provided, return without registering
		return
	}
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		handlerIndexToUse, validMethod := methodIndexRequested(r.Method)
		if !validMethod {
			// Invalid method, return without registering
			return
		}
		if handlerIndexToUse > len(handlers)-1 {
			// if there is no specific handler for the method, use the last handler
			handlerIndexToUse = len(handlers) - 1
		}
		// Handle the request with the appropriate handler
		fmt.Printf("Handling %s request for %s with handler index %d\n", r.Method, path, handlerIndexToUse)
		title, content := handlers[handlerIndexToUse](w, r)
		layout(title, r, content).Send(w)
	})
}
func Url(id string, base string, action ...UrlType) string {
	actn := UrlType("")
	if len(action) > 0 {
		if action[0] != "" {
			actn = action[0]
		}
	}
	// fmt.Println("Url called with id:", id, "base:", base, "action:", actn, "url: ", fmt.Sprintf("%s/%s%s", base, id, actn))
	switch actn {
	case ListUrl:
		return "/" + base
	case NewUrl:
		return "/" + base + "/new"
	case EditUrl:
		return "/" + base + "/" + id + "/edit"
	case ShowUrl:
		return "/" + base + "/" + id
	}
	return fmt.Sprintf("/%s/%s/%s", base, id, actn)
}
func RouteUrl(base string, action ...UrlType) string {
	return Url("{id}", base, action...)
}
func RouteBranch(basePath string, mux *http.ServeMux, layout Layout, paths map[string][]Handler) {
	for path, handlers := range paths {
		fmt.Println("Mapping ", basePath+path, " to ", handlers)
		RouteByMethod(mux, basePath+path, layout, handlers...)
	}
}
