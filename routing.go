package bolt

import (
	"fmt"
	"log"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request) (string, Element)
type Layout func(string, *http.Request, ...Element) Element
type HandlerBranch map[string]HandlerMethods
type HandlerMethods struct {
	Get    Handler
	Post   Handler
	Delete Handler
	Put    Handler
	Patch  Handler
}
type CrudType string

const ListUrl CrudType = "list"
const SaveUrl CrudType = "edit"
const EditUrl CrudType = "edit"
const NewUrl CrudType = "new"
const ShowUrl CrudType = "show"
const DeleteUrl CrudType = "delete"

// Routes requests for a given path based on the HTTP method.
// mux is the HTTP ServeMux to register the routes with.
// path is the base path for the routes.
// handlers is any number of handler functions that return bolt elements
func RouteByMethod(mux *http.ServeMux, path string, layout Layout, handlers HandlerMethods) {
	log.Println(`RouteByMethod path: `, path)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		var handler Handler
		switch r.Method {
		case http.MethodGet:
			handler = handlers.Get
		case http.MethodPost:
			handler = handlers.Post
		case http.MethodDelete:
			handler = handlers.Delete
		case http.MethodPut:
			handler = handlers.Put
		case http.MethodPatch:
			handler = handlers.Patch
		}
		if handler == nil {
			http.NotFound(w, r)
			return
		}

		// Handle the request with the appropriate handler
		fmt.Printf("Handling %s request for %s\n", r.Method, path)
		title, content := handler(w, r)
		layout(title, r, content).Send(w)
	})
}
func RouteBranch(mux *http.ServeMux, layout Layout, branch HandlerBranch) {
	for path, handlers := range branch {
		fmt.Println("Mapping ", path, " to ", handlers)
		RouteByMethod(mux, path, layout, handlers)
	}
}
