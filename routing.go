package bolt

import (
	"fmt"
	"log"
	"net/http"
	"unicode"
)

type Handler func(http.ResponseWriter, *http.Request) (Element, error)
type Layout func(http.ResponseWriter, *http.Request, ...Element) Element
type HandlerBranch map[string]HandlerMethods
type HandlerMethods struct {
	Get    Handler
	Post   Handler
	Delete Handler
	Put    Handler
	Patch  Handler
}
type UrlType string

const ListUrl UrlType = "list"
const SaveUrl UrlType = "save"
const EditUrl UrlType = "edit"
const NewUrl UrlType = "new"
const ShowUrl UrlType = "show"
const DeleteUrl UrlType = "delete"

func capitalizeFirstLetter(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// DEPRECATED: Use github.com/jaredtmartin/route instead
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
		content, err := handler(w, r)
		if err != nil {
			// w.WriteHeader(http.StatusBadRequest)
			// w.Write([]byte(capitalizeFirstLetter(err.Error())))
			http.Error(w, capitalizeFirstLetter(err.Error()), http.StatusBadRequest)
			return
		}
		layout(w, r, content).Send(w)
	})
}
func GetPageTitle(elements ...Element) string {
	if len(elements) == 0 || elements[0] == nil {
		return ""
	}
	return elements[0].GetAttr("page-title")
}

// DEPRECATED: Use github.com/jaredtmartin/route instead
func RouteBranch(mux *http.ServeMux, layout Layout, branch HandlerBranch) {
	for path, handlers := range branch {
		fmt.Println("Mapping ", path, " to ", handlers)
		RouteByMethod(mux, path, layout, handlers)
	}
}

// DEPRECATED: Use github.com/jaredtmartin/route instead
func Route(mux *http.ServeMux, path string, layout Layout, handler Handler) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		content, err := handler(w, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(capitalizeFirstLetter(err.Error())))
			return
		}
		if content == nil {
			content = None()
		}
		layout(w, r, content).Send(w)
	})
}
